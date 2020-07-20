import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

import Page from 'src/components/Page';
import SearchBar from 'src/components/SearchBar';
import Header from './Header';
import Results from './Results';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  }
}));

const GET_DELEGATES = gql`
  query GetDelegates {
    delegates {
      edges {
        uuid
        email
        firstName
        lastName
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

function DelegatesList() {
  const classes = useStyles();
  const { loading, error, data, refetch } = useQuery(GET_DELEGATES, {
    fetchPolicy: 'cache-and-network'
  });
  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;
  const delegates = data?.delegates?.edges;

  const handleSearch = () => {};

  const handleNewDelegate = data => {
    if (!data.createDelegate) return;

    refetch();
  };

  return (
    <Page className={classes.root} title="Delegates">
      <Container maxWidth={false}>
        <Header onCreate={handleNewDelegate} />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        <Results className={classes.results} delegates={delegates} />
      </Container>
    </Page>
  );
}

export default DelegatesList;
