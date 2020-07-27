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
import ReoderableList from 'src/components/ReorderableList/ReorderableList';
import SyllabusSearch from 'src/components/SyllabusSearch';

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
  query SearchSyllabus(
    $name: String!
    $excludeLesson: Boolean
    $excludeTest: Boolean
    $excludeModule: Boolean
    $page: Page!
  ) {
    searchSyllabus(
      filter: {
        name: $name
        excludeLesson: $excludeLesson
        excludeTest: $excludeTest
        excludeModule: $excludeModule
      }
      page: $page
    ) {
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
      page
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

function useSearchQuery(text, filter, page) {
  console.log('page: ', page);
  const { error, data } = useQuery(SEARCH, {
    variables: {
      name: text,
      excludeLesson: filter.excludeLesson,
      excludeTest: filter.excludeTest,
      excludeModule: filter.excludeModule,
      page: page
    }
  });

  if (
    !data ||
    !data.searchSyllabus ||
    !data.searchSyllabus.edges ||
    !data.searchSyllabus.pageInfo
  ) {
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

function CourseBuilder({ state, setState }) {
  const classes = useStyles();

  const [searchFilters, setSearchFilters] = React.useState({
    excludeLesson: false,
    excludeTest: false,
    excludeModule: false,
    filters: [
      { name: 'Exclude Tests', type: 'Filter', isFilter: 'excludeTest' },
      { name: 'Exclude Lessons', type: 'Filter', isFilter: 'excludeLesson' },
      { name: 'Exclude Modules', type: 'Filter', isFilter: 'excludeModule' }
    ]
  });
  const [searchText, setSearchText] = React.useState('');
  const page = { total: 4, offset: 0, given: 4 };
  const searchResults = useSearchQuery(searchText, searchFilters, page);

  const [courseStructure, setCourseStructure] = React.useState([]);
  const { resultItems } = useModulesQuery(page);

  // Add module syllabus and name
  const addModule = array => {
    array.forEach(element => {
      if (element.type === 'module') {
        const module = resultItems.find(x => x.uuid === element.uuid);
        if (module && !courseStructure.find(x => x.uuid === module.uuid)) {
          setCourseStructure([...courseStructure, module]);
        }
      }
    });
  };

  console.log('strut: ', courseStructure);
  React.useEffect(() => {
    if (state.syllabus) {
      addModule(state.syllabus);
    } else if (state.structure) {
      addModule(state.structure);
    }
  });

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
            <Typography variant="body1" color="textPrimary">
              All Courses are comprised of <strong>Modules</strong>,{' '}
              <strong>Lessons</strong> and <strong>Tests</strong>
              <br />
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
                <SyllabusSearch
                  placeholder="Search Modules, Lessons or Tests"
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
              title="Suggested Modules based on Tags"
              suggestions={resultItems.slice(0, 3)}
              onAdd={({ uuid, name }) =>
                setState({
                  syllabus: [...state.syllabus, { uuid, name, type: 'lesson' }]
                })
              }
            />
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Course Structure" />
              <Divider />
              <CardContent>
                <ReoderableList
                  items={courseStructure.map(module => ({
                    id: module.uuid,
                    component: (
                      <ReoderableDropdown
                        title={module.name}
                        items={
                          module.syllabus &&
                          module.syllabus.map(item => ({
                            id: item.uuid,
                            component: (
                              <ReoderableListItem
                                uuid={item.uuid}
                                text={item.name}
                                onDelete={onDelete}
                              />
                            )
                          }))
                        }
                        setItems={items => {
                          const newModule = {
                            ...module,
                            syllabus: items
                          };
                          setCourseStructure([
                            ...courseStructure,
                            ...newModule
                          ]);
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
