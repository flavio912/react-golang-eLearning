import React, { useState } from 'react';
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

const GET_ADMINS = gql`
  query GetAdmins {
    admins {
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

function AdminsList() {
  const classes = useStyles();
  const { loading, error, data, refetch } = useQuery(GET_ADMINS, {
    fetchPolicy: 'cache-and-network'
  });
  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;
  const admins = data?.admins?.edges;

  const handleFilter = () => {};
  const handleSearch = () => {};

  const handleNewAdmin = data => {
    if (!data.createAdmin) return;

    refetch();
  };

  return (
    <Page className={classes.root} title="Admins">
      <Container maxWidth={false}>
        <Header onCreateNewAdmin={handleNewAdmin} />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        <Results className={classes.results} admins={admins} />
      </Container>
    </Page>
  );
}

export default AdminsList;
