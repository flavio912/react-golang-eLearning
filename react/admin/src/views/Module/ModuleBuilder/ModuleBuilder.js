import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Typography,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';
import ReoderableListItem from 'src/components/ReorderableList/ReorderableListItem';
import ReoderableDropdown from 'src/components/ReorderableList/ReorderableDropdown';
import SuggestedTable from 'src/components/SuggestedTable';
import SyllabusSearch from 'src/components/SyllabusSearch';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  heading: {
    margin: theme.spacing(2)
  }
}));

const GET_LESSONS = gql`
  query SearchLessons($tags: [UUID]) {
    lessons(filter: { tags: $tags }) {
      edges {
        uuid
        name
        text
        type
        tags {
          uuid
        }
      }
      pageInfo {
        total
        limit
        offset
        given
      }
    }
  }
`;

const SEARCH = gql`
  query SearchSyllabus($name: String!, $excludeLesson: Boolean, $excludeTest: Boolean, $page: Page!) {
    searchSyllabus(filter: { name: $name, excludeLesson:$excludeLesson, excludeTest:$excludeTest, excludeModule: true }, page: $page) {
      edges {
        uuid
        name
        type
        complete
      }
      pageInfo {
        total
        limit
        offset
        given
      }
    }
  }
`;

function useSuggestedQuery(text, page) {
  const { error, data } = useQuery(GET_LESSONS, {
    variables: {
      name: text,
      page: page
    }
  });

  if (!data || !data.lessons || !data.lessons.edges || !data.lessons.pageInfo) {
    error && console.error('Could not get data', data, error);
    return {
      resultItems: [],
      pageInfo: {
        total: 1,
        offset: 0,
        limit: 4,
        given: 0
      }
    };
  }

  const resultItems = data.lessons.edges.map(lesson => ({
    uuid: lesson?.uuid ?? '',
    name: lesson?.name ?? '',
    text: lesson?.text ?? '',
    type: lesson?.type ?? '',
    tags: lesson?.tags ?? ''
  }));

  const pageInfo = {
    total: data.lessons.pageInfo?.total,
    offset: data.lessons.pageInfo.offset,
    limit: data.lessons.pageInfo.limit,
    given: data.lessons.pageInfo.given
  };

  return { resultItems, pageInfo };
}

function useSearchQuery(text, filter, page) {
  const { error, data } = useQuery(SEARCH, {
    variables: {
      name: text,
      excludeLesson: filter.excludeLesson,
      excludeTest: filter.excludeTest,
      page: page
    }
  });

  if (!data || !data.searchSyllabus || !data.searchSyllabus.edges || !data.searchSyllabus.pageInfo) {
    error && console.error('Could not get data', data, error);
    return {
      resultItems: [],
      pageInfo: {
        total: 1,
        offset: 0,
        limit: 4,
        given: 0
      }
    };
  }

  const resultItems = data.searchSyllabus.edges.map(syllabus => ({
    uuid: syllabus?.uuid ?? '',
    name: syllabus?.name ?? '',
    text: syllabus?.text ?? '',
    type: syllabus?.type ?? '',
    tags: syllabus?.tags ?? ''
  }));

  const pageInfo = {
    total: data.searchSyllabus.pageInfo?.total,
    offset: data.searchSyllabus.pageInfo.offset,
    limit: data.searchSyllabus.pageInfo.limit,
    given: data.searchSyllabus.pageInfo.given
  };

  return { resultItems, pageInfo };
}

function ModuleBuilder({ state, setState }) {
  const classes = useStyles();

  const [searchFilters, setSearchFilters] = React.useState({
    excludeLesson: false, excludeTest: false, filters: [
      {name: 'Exclude Tests', type: 'Filter', isFilter: 'excludeTest'},
      {name: 'Exclude Lessons', type: 'Filter', isFilter: 'excludeLesson'},
    ]
  });
  const [searchText, setSearchText] = React.useState('');
  const page = { total: 4, offset: 0, given: 4 };
  const searchResults = useSearchQuery(searchText, searchFilters, page);
  const suggestedLessons = useSuggestedQuery(state.tags);

  const onDelete = uuid => {
    setState({
      syllabus: state.syllabus.filter(item => item.uuid !== uuid)
    });
  };
  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item className={classes.heading}>
            <Typography
              variant="h3"
              color="textPrimary"
              className={classes.header}
            >
              Build this Module
            </Typography>
            <Typography variant="body1" color="textPrimary">
              All modules are comprised of <strong>Lessons</strong> and{' '}
              <strong>Tests</strong>
              <br />
              Add your first lesson to get started
            </Typography>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader
                title="Search Lessons and Tests"
                className={classes.noPadding}
              />
              <CardContent>
                <SyllabusSearch
                  placeholder="Search Lessons or Tests"
                  searchFilters={searchFilters}
                  setSearchFilters={setSearchFilters}
                  searchResults={searchResults.resultItems}
                  setSearchText={setSearchText}
                  onChange={({ uuid, name, type }) => {
                    setState({
                      syllabus: [...state.syllabus, { uuid, name, type }]
                    });
                  }}
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <SuggestedTable
              title="Suggested Lessons based on Tags"
              suggestions={suggestedLessons.resultItems.slice(0, 3)}
              onAdd={({ uuid, name }) =>
                setState({ syllabus: [...state.syllabus, { uuid, name, type: 'lesson' }] })
              }
            />
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Module Structure" />
              <Divider />
              <CardContent>
                <ReoderableDropdown
                  title="Module 1"
                  items={state.syllabus.map(({ name, uuid }) => ({
                    id: uuid,
                    component: (
                      <ReoderableListItem
                        uuid={uuid}
                        text={name}
                        onDelete={onDelete}
                      />
                    )
                  }))}
                  setItems={setState}
                />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default ModuleBuilder;
