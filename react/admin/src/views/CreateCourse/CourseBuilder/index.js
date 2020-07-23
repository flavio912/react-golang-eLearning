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
import ReoderableList from 'src/components/ReorderableList/ReorderableList';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  heading: {
    margin: theme.spacing(2)
  }
}));

const GET_MODULES = gql`
  query SearchModules($page: Page!) {
    modules(page: $page) {
      edges {
        uuid
        name
        type
        syllabus {
          uuid
          name
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
  query SearchSyllabus($name: String!, $page: Page!) {
    searchSyllabus(filter: { name: $name }, page: $page) {
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

function useModulesQuery(page) {
  const { error, data } = useQuery(GET_MODULES, {
    variables: {
      page,
    }
  });

  if (!data || !data.modules || !data.modules.edges || !data.modules.pageInfo) {
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

  const resultItems = data.modules.edges.map(module => ({
    uuid: module?.uuid ?? '',
    name: module?.name ?? '',
    type: module?.type ?? '',
    syllabus: module?.syllabus ?? ''
  }));

  const pageInfo = {
    total: data.modules.pageInfo?.total,
    offset: data.modules.pageInfo.offset,
    limit: data.modules.pageInfo.limit,
    given: data.modules.pageInfo.given
  };

  return { resultItems, pageInfo };
}

function useSearchQuery(text, page) {
  const { error, data } = useQuery(SEARCH, {
    variables: {
      name: text,
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

function CourseBuilder({ state, setState }) {
  const classes = useStyles();

  const [searchText, setSearchText] = React.useState('');
  const page = { total: 4, offset: 0, given: 4 };
  const searchResults = useSearchQuery(searchText, page);

  const [courseStructure, setCourseStructure] = React.useState([]);
  const { resultItems } = useModulesQuery(page);

  // Add module syllabus and name
  const addModule = (array) => {
    array.map((element) => {
      if (element.type === 'module') {
        const module = resultItems.find(x => x.uuid === element.uuid);
        if (module && !courseStructure.find(x => x.uuid === module.uuid)) {        
          setCourseStructure([
            ...courseStructure,
            module,
          ])
        }
      }
    })
  }

  console.log('strut: ', courseStructure)
  React.useEffect(() => {
    if (state.syllabus) {
      addModule(state.syllabus);
    } else if (state.structure) {
      addModule(state.structure);
    }
  })

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
              Build this Course
            </Typography>
            <Typography
              variant="body1"
              color="textPrimary"
            >
              All Courses are comprised of <strong>Modules</strong>, <strong>Lessons</strong> and <strong>Tests</strong><br />
              Add your first Module or Lesson below to get started
            </Typography>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader
                title="Search Modules"
                className={classes.noPadding}
              />
              <CardContent>
                <Autocomplete
                  freeSolo
                  options={searchResults.resultItems}
                  getOptionLabel={option => option.name}
                  onChange={(_, { uuid, name, type }) =>
                    setState({
                      syllabus: [...state.syllabus, { uuid, name, type }]
                    })
                  }
                  renderInput={params => (
                    <TextField
                      {...params}
                      placeholder="Search Modules, Lessons or Tests"
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
              title="Suggested Modules based on Tags"
              suggestions={resultItems.slice(0, 3)}
              onAdd={({ uuid, name }) =>
                setState({ syllabus: [...state.syllabus, { uuid, name, type: 'lesson' }] })
              }
            />
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Course Structure" />
              <Divider />
              <CardContent>
                <ReoderableList
                  items={courseStructure.map((module) => ({
                    id: module.uuid,
                    component: (
                      <ReoderableDropdown
                        title={module.name}
                        items={module.syllabus && module.syllabus.map((item) => ({
                          id: item.uuid,
                          component: (
                            <ReoderableListItem
                              uuid={item.uuid}
                              text={item.name}
                              onDelete={onDelete}
                            />
                          )
                        }))}
                        setItems={(items) => {
                          const newModule = {
                            ...module,
                            syllabus: items,
                          };
                          setCourseStructure([
                            ...courseStructure,
                            ...newModule
                          ])
                        }}
                      />
                    )
                  }))}
                  setItems={setCourseStructure}
                />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default CourseBuilder;
