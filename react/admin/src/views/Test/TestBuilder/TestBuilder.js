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
  marginBottom: {
    marginBottom: '100px',
  },
  padding: {
    padding: theme.spacing(2),
  }
}));

const GET_QUESTIONS = gql`
  query SearchQuestions {
    questions {
      edges {
        uuid
        text
        answers {
          uuid
          text
        }
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
    text: item?.text ?? '',
    tags: item?.tags ?? '',
  }));

  return resultItems;
}

function useSearchQuery() {
  const { error, data } = useQuery(GET_QUESTIONS);
  return cleanResult(data && data.questions, error);
}

function TestBuilder({ state, setState }) {
  const classes = useStyles();

  const [searchFilters, setSearchFilters] = React.useState({filters: []});
  const [searchText, setSearchText] = React.useState('');
  const searchResults = useSearchQuery(searchText, searchFilters);

  const onDelete = (uuid) => {
    const newStructure = [...state.questions];
    const index = state.questions.findIndex(x => x.uuid === uuid);
    newStructure.splice(index, 1);
    setState('questions', newStructure);
  };

  const newItem = (item) => ({
    uuid: item.uuid,
    item,
    items: [],
    component: (
      <ReoderableListItem
        uuid={item.uuid}
        text={item.text}
        onDelete={onDelete}
      />
    )
  })

  console.log('stru: ', state.questions)

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
              Build this Test
            </Typography>
            <Typography variant="body1" color="textPrimary">
              All Tests are comprised of <strong>Questions</strong> and <strong>Answers</strong>
              <br />
              Add your first Question below to get started
            </Typography>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader
                title="Search Questions"
                className={classes.noPadding}
              />
              <CardContent>
                <SyllabusSearch
                  placeholder="Search Questions"
                  searchFilters={searchFilters}
                  setSearchFilters={setSearchFilters}
                  searchResults={searchResults.map((result) =>
                    ({ ...result, name: result.text, type: 'Question'})
                  )}
                  setSearchText={setSearchText}
                  onChange={(item) => {
                    // No duplicates
                    if (!state.questions.find(x => x.uuid === item.uuid)) {
                      setState('questions', [...state.questions, item]);
                    }
                  }}
                />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <SuggestedTable
              title="Suggested Questions based on Tags"
              suggestions={searchResults.slice(0, 3).map((result) =>
                ({ ...result, name: result.text, type: 'Question'})
              )}
              onAdd={(item) => {
                // No duplicates
                if (!state.questions.find(x => x.uuid === item.uuid)) {
                  setState('questions', [...state.questions, item]);
                }
              }}
            />
          </Grid>
          <Grid item className={classes.marginBottom}>
            <Card>
              <CardHeader title="Test Structure" />
              <Divider />
              <ReoderableList
                className={classes.padding}
                newItem={newItem}
                items={state.questions.map(item => newItem(item))}
                setItems={items => {
                  setState('questions', items.map(({item}) => item));
                }}
              />
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default TestBuilder;
