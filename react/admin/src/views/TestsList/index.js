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
        given
        total
      }
    }
  }
`;

function TestsList({ match, history }) {
  const classes = useStyles();

  // const handleFilter = () => {};

  const handleSearch = () => {};

  const { loading, error, data } = useQuery(GET_TESTS, {
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
    <Page className={classes.root} title="Tests">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/test/create/overview');
          }}
        />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        {data.tests && (
          <Results className={classes.results} tests={data.tests.edges} />
        )}
      </Container>
    </Page>
  );
}

export default TestsList;
