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
  },
  margin: {
    margin: `${theme.spacing(2)}px 0`,
  },
  marginBottom: {
    marginBottom: '100px',
  },
  padding: {
    padding: theme.spacing(2),
  }
}));

const GET_MODULES = gql`
  query SearchModules {
    modules {
      edges {
        uuid
        name
        type
        syllabus {
          uuid
          name
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
    $excludeModule: Boolean
  ) {
    searchSyllabus(
      filter: {
        name: $name
        excludeLesson: $excludeLesson
        excludeTest: $excludeTest
        excludeModule: $excludeModule
      }
    ) {
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
    type: item?.type ?? '',
    text: item?.text ?? '',
    syllabus: item?.syllabus ?? ''
  }));

  return resultItems;
}

function useModulesQuery() {
  const { error, data } = useQuery(GET_MODULES);
  return cleanResult(data && data.modules, error);
}

function useSearchQuery(text, filter) {
  const { error, data } = useQuery(SEARCH, {
    variables: {
      name: text,
      excludeLesson: filter.excludeLesson,
      excludeTest: filter.excludeTest,
      excludeModule: filter.excludeModule,
    }
  });
  return cleanResult(data && data.searchSyllabus, error);
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
  const searchResults = useSearchQuery(searchText, searchFilters);
  const resultItems = useModulesQuery();

  React.useEffect(() => {
    state.syllabus.forEach(element => {
      // Add module syllabus and name
      if (element.type === 'module' && element.syllabus === undefined) {
        const module = resultItems.find(x => x.uuid === element.uuid);
        const index = module && state.syllabus.findIndex(x => x.uuid === module.uuid);
        if (index > -1) {
          const syllabus = [...state.syllabus];
          syllabus[index] = module;
          setState({ syllabus });
        }
      }
    });
  }, [resultItems]);

  const onDelete = (uuid) => {
    const newStructure = [...state.syllabus];
    const index = state.syllabus.findIndex(x => x.uuid === uuid);
    newStructure.splice(index, 1);
    setState({ syllabus: newStructure });
  };

  const newItem = (item) => ({
    uuid: item.uuid,
    item,
    items: item.syllabus ? item.syllabus : [],
    component: (
      item.type === 'module' ?
        <ReoderableDropdown
          className={classes.margin}
          title={item.name}
          onDelete={() => onDelete(item.uuid)}
          items={
            item.syllabus &&
            item.syllabus.map(child => ({
              uuid: child.uuid,
              item: child.item,
              component: (
                <ReoderableListItem
                  className={classes.margin}
                  uuid={child.uuid}
                  text={child.name}
                />
              )
            }))
          }
        />
      :
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
                  searchResults={searchResults}
                  setSearchText={setSearchText}
                  onChange={(item) => {
                    // No duplicates
                    if (module && !state.syllabus.find(x => x.uuid === item.uuid)) {
                      setState({ syllabus: [...state.syllabus, item] });
                    }
                  }}
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <SuggestedTable
              title="Suggested Modules based on Tags"
              suggestions={resultItems.slice(0, 3)}
              onAdd={(module) => {
                const toAdd = resultItems.find(x => x.uuid === module.uuid);
                // No duplicates
                if (module && !state.syllabus.find(x => x.uuid === module.uuid)) {
                  setState({ syllabus: [...state.syllabus, toAdd] });
                }
              }}
            />
          </Grid>
          <Grid item className={classes.marginBottom}>
            <Card>
              <CardHeader title="Course Structure" />
              <Divider />
                <ReoderableList
                  className={classes.padding}
                  newItem={newItem}
                  items={state.syllabus.map(item => newItem(item))}
                  setItems={items => {
                    setState({ syllabus: items.map(({item}) => item)});
                  }}
                />
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default CourseBuilder;
