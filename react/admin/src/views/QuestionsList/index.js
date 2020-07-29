import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Button } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import { Link as RouterLink } from 'react-router-dom';
import Page from 'src/components/Page';
import SearchBar from 'src/components/SearchBar';
import Results from 'src/components/Results';
import Header from './Header';

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
        total
        offset
        limit
        given
      }
    }
  }
`;

function QuestionsList({ match, history }) {
  const classes = useStyles();

  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_QUESTIONS, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage,
      },
      filter: {
        text: searchText
      },
      orderBy: {
        ascending: false,
        field: 'created_at'
      }
    },
    fetchPolicy: 'cache-and-network'
  });

  if (error) return <div>{error.message}</div>;

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  // Results table
  const headers = ['Name', 'Tests Linked', 'Possible Answers', 'Actions'];
  const cells = [
    { field: 'text' }, { field: 'numTestsUsedIn' }, { field: 'numAnswers' },
    { component: (result) => (
      <Button
        color="primary"
        component={RouterLink}
        size="small"
        to={`/question/${result.uuid}/overview`}
        variant="outlined"
      >
        Edit
      </Button>
    )}
  ]

  return (
    <Page className={classes.root} title="Questions">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/questions/create/overview');
          }}
        />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.questions}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default QuestionsList;
