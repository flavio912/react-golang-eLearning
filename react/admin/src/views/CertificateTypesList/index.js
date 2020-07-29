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

const GET_CERTIFICATE_TYPES = gql`
  query GetCertificateTypes($page: Page, $filter: CertificateTypeFilter) {
    certificateTypes(page: $page, filter: $filter) {
      edges {
        uuid
        name
        createdAt
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

function CertificateTypes({ match, history }) {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_CERTIFICATE_TYPES, {
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
  const headers = ['Name', 'Actions'];
  const cells = [
    { field: 'name' },
    { component: (result) => (
      <Button
        color="primary"
        component={RouterLink}
        size="small"
        to={`/certificateTypes/${result.uuid}/overview`}
        variant="outlined"
      >
        Edit
      </Button>
    )}
  ]

  return (
    <Page className={classes.root} title="Certificate Types">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/certificateTypes/create/overview');
          }}
        />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.certificateTypes}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default CertificateTypes;
