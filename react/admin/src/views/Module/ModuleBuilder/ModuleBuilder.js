import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Typography
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
  },
  marginBottom: {
    marginBottom: '100px',
  },
  padding: {
    padding: theme.spacing(2),
  }
}));

const GET_LESSONS = gql`
  query SearchLessons($tags: [UUID]) {
    lessons(filter: { tags: $tags }) {
      edges {
        uuid
        name
        type
        tags {
          uuid
        }
      }
    }
  }
`;

const SEARCH = gql`
  query SearchSyllabus(
    $name: String!
    $excludeLesson: Boolean
    $excludeTest: Boolean
  ) {
    searchSyllabus(filter: {
      name: $name
      excludeLesson: $excludeLesson
      excludeTest: $excludeTest
      excludeModule: true
    }) {
      edges {
        uuid
        name
        type
      }
    }
  }
`;

const cleanResult = (data, error) => {
  if (!data || !data.edges) {
    error && console.error('Could not get data', data, error);
    return [];
  }

  const resultItems = data.edges.map(item => ({
    uuid: item?.uuid ?? '',
    name: item?.name ?? '',
    text: item?.text ?? '',
    type: item?.type ?? '',
    tags: item?.tags ?? ''
  }));

  return resultItems;
}

function useSuggestedQuery(text) {
  const { error, data } = useQuery(GET_LESSONS, {
    variables: {
      name: text,
    }
  });

  return cleanResult(data && data.lessons, error);
}

function useSearchQuery(text, filter) {
  const { error, data } = useQuery(SEARCH, {
    variables: {
      name: text,
      excludeLesson: filter.excludeLesson,
      excludeTest: filter.excludeTest,
    }
  });

  return cleanResult(data && data.searchSyllabus, error);
}

function ModuleBuilder({ state, setState }) {
  const classes = useStyles();

  const [searchFilters, setSearchFilters] = React.useState({
    excludeLesson: false,
    excludeTest: false,
    filters: [
      { name: 'Exclude Tests', type: 'Filter', isFilter: 'excludeTest' },
      { name: 'Exclude Lessons', type: 'Filter', isFilter: 'excludeLesson' },
    ]
  });
  const [searchText, setSearchText] = React.useState('');
  const searchResults = useSearchQuery(searchText, searchFilters);
  const suggestedLessons = useSuggestedQuery(state.tags);

  const onDelete = uuid => {
    setState({
      syllabus: state.syllabus.filter(item => item.uuid !== uuid)
    });
  };

  const newItem = (item) => ({
    uuid: item.uuid,
    item,
    items: [],
    component: (
        <ReoderableListItem
          uuid={item.uuid}
          text={item.name}
          onDelete={onDelete}
        />
    )
  })

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
                  searchResults={searchResults}
                  setSearchText={setSearchText}
                  onChange={({ uuid, name, type }) => {
                    // No duplicates
                    if (!state.syllabus.find(x => x.uuid === uuid)) {
                      setState({
                        syllabus: [...state.syllabus, { uuid, name, type }]
                      });
                    }
                  }}
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <SuggestedTable
              title="Suggested Lessons based on Tags"
              suggestions={suggestedLessons.slice(0, 3)}
              onAdd={({ uuid, name, type }) => {
                // No duplicates
                if (!state.syllabus.find(x => x.uuid === uuid)) {
                  setState({
                    syllabus: [...state.syllabus, { uuid, name, type }]
                  });
                }
              }}
            />
          </Grid>
          <Grid item className={classes.marginBottom}>
            <Card>
              <CardHeader title="Module Structure" />
              <Divider />
              <ReoderableDropdown
                className={classes.padding}
                newItem={newItem}
                title={state.name ?? ''}
                items={state.syllabus.map((item) => (
                  newItem(item)
                ))}
                setItems={items => {
                  setState({ syllabus: items.map(({item}) => item)})
                }}
              />
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default ModuleBuilder;
