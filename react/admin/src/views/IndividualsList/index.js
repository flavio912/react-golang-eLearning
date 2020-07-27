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

const GET_INDIVIDUALS = gql`
  query Getindividuals {
    individuals {
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

function IndividualsList() {
  const classes = useStyles();
  const { loading, error, data, refetch } = useQuery(GET_INDIVIDUALS, {
    fetchPolicy: 'cache-and-network'
  });
  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;
  const individuals = data?.individuals?.edges;

  const handleSearch = () => {};

  const handleNewIndividual = data => {
    if (!data.createIndividual) return;

    refetch();
  };

  return (
    <Page className={classes.root} title="Individuals">
      <Container maxWidth={false}>
        <Header onCreateNewIndividual={handleNewIndividual} />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        <Results className={classes.results} individuals={individuals} />
      </Container>
    </Page>
  );
}

export default IndividualsList;
