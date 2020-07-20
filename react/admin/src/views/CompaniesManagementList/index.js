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

const GET_COMPANIES = gql`
  query GetCompanies {
    companies {
      edges {
        uuid
        name
        managers {
          edges {
            email
          }
          pageInfo {
            total
          }
        }
        delegates {
          edges {
            email
          }
          pageInfo {
            total
          }
        }
      }
    }
  }
`;

function CompaniesManagementList() {
  const classes = useStyles();

  const { loading, error, data } = useQuery(GET_COMPANIES);

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;

  const handleFilter = () => {};

  const handleSearch = () => {};
  const companies = data?.companies?.edges.map(comp => ({
    id: comp.uuid,
    name: comp.name,
    email: comp.managers?.edges[0]?.email,
    logo: 'https://cdn.cnn.com/cnnnext/dam/assets/180301124611-fedex-logo.png',
    noDelegates: 40,
    noManagers: comp.managers?.pageInfo?.total,
    paymentType: 'Contract'
  }));

  return (
    <Page className={classes.root} title="Companies Management List">
      <Container maxWidth={false}>
        <Header />
        <SearchBar onFilter={handleFilter} onSearch={handleSearch} />
        {companies && (
          <Results className={classes.results} companies={companies} />
        )}
      </Container>
    </Page>
  );
}

export default CompaniesManagementList;
