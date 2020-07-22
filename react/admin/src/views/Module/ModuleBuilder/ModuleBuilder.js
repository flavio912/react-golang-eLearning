import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Typography,
  TextField,
  InputAdornment
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import { useQuery } from '@apollo/react-hooks';
import Autocomplete from '@material-ui/lab/Autocomplete';
import SearchIcon from '@material-ui/icons/Search';
import ReoderableListItem from 'src/components/ReorderableList/ReorderableListItem';
import ReoderableDropdown from 'src/components/ReorderableList/ReorderableDropdown';
import SuggestedTable from 'src/components/SuggestedTable';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  heading: {
    margin: theme.spacing(2)
  }
}));

const lessons = [
  {
    uuid: '1231231231',
    name: 'Firesafety Module 1 - Lesson 1',
    numCoursesUsedIn: 3,
    type: 'Lesson',
    tags: [
      {
        name: 'FIRE SAFETY',
        color: 'default'
      },
      {
        name: 'HEALTH & SAFETY',
        color: 'default'
      }
    ]
  },
  {
    uuid: '1231231232',
    name: 'Firesafety Module 1 - Lesson 2',
    numCoursesUsedIn: 3,
    type: 'Lesson',
    tags: [
      {
        name: 'FIRE SAFETY',
        color: 'default'
      }
    ]
  },
  {
    uuid: '1231231233',
    name: 'Firesafety Module 1 - Lesson 3',
    numCoursesUsedIn: 3,
    type: 'Lesson',
    tags: [
      {
        name: 'FIRE SAFETY',
        color: 'default'
      }
    ]
  }
];

const GET_LESSONS = gql`
  query SearchLessons($name: String!, $page: Page!) {
    lessons(filter: { name: $name }, page: $page) {
      edges {
        uuid
        name
        text
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

function useSearchQuery(text, page) {
  const { error, data } = useQuery(GET_LESSONS, {
    variables: {
      name: text,
      page: page
    }
  });

  if (!data || !data.lessons || !data.lessons.edges || !data.lessons.pageInfo) {
    console.error('Could not get data', data, error);
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

  console.log(data);

  const resultItems = data.lessons.edges.map(lesson => ({
    uuid: lesson?.uuid ?? '',
    name: lesson?.name ?? '',
    text: lesson?.text ?? '',
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

function ModuleBuilder({ state, setState }) {
  const classes = useStyles();

  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState({ total: 4, offset: 0, given: 4 });
  const searchResults = useSearchQuery(searchText, page);

  const onDelete = uuid => {
    setState({
      syllabus: state.syllabus.filter(item => item.uuid !== uuid)
    });
  };

  console.log(state.syllabus);

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
                <Autocomplete
                  freeSolo
                  options={searchResults.resultItems}
                  getOptionLabel={option => option.name}
                  onChange={(_, { uuid, name }) =>
                    setState({
                      syllabus: [...state.syllabus, { uuid, name }]
                    })
                  }
                  renderInput={params => (
                    <TextField
                      {...params}
                      placeholder="Search Lessons or Tests"
                      onChange={inp => {
                        setSearchText(inp.target.value);
                      }}
                      InputProps={{
                        ...params.InputProps,
                        startAdornment: (
                          <InputAdornment position="start">
                            <SearchIcon />
                          </InputAdornment>
                        )
                      }}
                    />
                  )}
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <SuggestedTable
              title="Suggested Lessons based on Tags"
              lessons={searchResults.resultItems.slice(0, 3)}
              onAdd={({ uuid, name }) =>
                setState({ syllabus: [...state.syllabus, { uuid, name }] })
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
