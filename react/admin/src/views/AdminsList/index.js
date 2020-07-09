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

  const { loading, error, data } = useQuery(GET_ADMINS);

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;

  const handleFilter = () => {};
  const handleSearch = () => {};
  const admins = data?.admins?.edges;

  // const admins = [
  //   {
  //     fullName: 'Alexandria Tutor',
  //     email: 'alexandria@tutor.com',
  //     lastLogin: {
  //       date: '10/2/2020'
  //     },
  //     createdAt: '4/1/2020',
  //     cin: '1000345323898'
  //   }
  // ];

  return (
    <Page className={classes.root} title="Admins">
      <Container maxWidth={false}>
        <Header />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        {admins && <Results className={classes.results} admins={admins} />}
      </Container>
    </Page>
  );
}

export default AdminsList;
