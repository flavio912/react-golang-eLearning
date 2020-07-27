import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import Page from 'src/components/Page';
import SearchBar from 'src/components/SearchBar';
import Header from './Header';
import Results from './Results';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  }
}));

const GET_QUESTIONS = gql`
  query GetQuestions($page: Page, $filter: QuestionFilter, $orderBy: OrderBy) {
    questions(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        text
      }
      pageInfo {
        given
        total
      }
    }
  }
`;

function QuestionsList({ match, history }) {
  const classes = useStyles();

  // const handleFilter = () => {};

  const handleSearch = () => {};

  const { loading, error, data } = useQuery(GET_QUESTIONS, {
    variables: {
      page: {
        offset: 0,
        limit: 100
      },
      filter: {},
      orderBy: {
        ascending: false,
        field: 'created_at'
      }
    }
  });

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;

  return (
    <Page className={classes.root} title="Questions">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/questions/create/overview');
          }}
        />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        {data.questions && (
          <Results
            className={classes.results}
            questions={data.questions.edges}
          />
        )}
      </Container>
    </Page>
  );
}

export default QuestionsList;
