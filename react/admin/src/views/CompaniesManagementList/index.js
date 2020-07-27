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
  query GetCompanies($page: Page, $filter: CompanyFilter, $orderBy: OrderBy) {
    companies(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        name
        isContract
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
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { loading, error, data } = useQuery(GET_COMPANIES, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage,
      },
      filter: {
        name: searchText
      },
      orderBy: {
        ascending: false,
        field: 'name'
      }
    }
  });

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;

  const companies = data?.companies?.edges.map(comp => ({
    id: comp.uuid,
    name: comp.name,
    email: comp.managers?.edges[0]?.email,
    logo: 'https://cdn.cnn.com/cnnnext/dam/assets/180301124611-fedex-logo.png',
    noDelegates: comp.delegates?.pageInfo?.total,
    noManagers: comp.managers?.pageInfo?.total,
    paymentType: comp.isContract ? 'Contract' : 'Pay as you go'
  }));

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  return (
    <Page className={classes.root} title="Companies Management List">
      <Container maxWidth={false}>
        <Header />
        <SearchBar setSearchText={setSearchText} />
        {companies && (
          <Results
            className={classes.results}
            companies={companies}
            page={page}
            handleChangePage={handleChangePage}
            rowsPerPage={rowsPerPage}
            handleChangeRowsPerPage={handleChangeRowsPerPage}
          />
        )}
      </Container>
    </Page>
  );
}

export default CompaniesManagementList;
