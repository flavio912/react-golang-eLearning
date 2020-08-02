import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Button } from '@material-ui/core';
import { gql } from 'apollo-boost';
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

const GET_TESTS = gql`
  query GetTests($page: Page, $filter: TestFilter, $orderBy: OrderBy) {
    tests(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        name
        attemptsAllowed
        passPercentage
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

function TestsList({ match, history }) {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_TESTS, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage
      },
      filter: {
        name: searchText
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
  const headers = ['Name', 'Pass Percentage', 'Attempts Allowed', 'Actions'];
  const cells = [
    { field: 'name' },
    { field: 'passPercentage' },
    { field: 'attemptsAllowed' },
    {
      component: result => (
        <Button
          color="primary"
          component={RouterLink}
          size="small"
          to={`/test/${result.uuid}/overview`}
          variant="outlined"
        >
          Edit
        </Button>
      )
    }
  ];

  return (
    <Page className={classes.root} title="Tests">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/test/create/overview');
          }}
        />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.tests}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default TestsList;
