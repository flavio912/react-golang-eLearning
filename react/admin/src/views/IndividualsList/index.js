import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Avatar, Link } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
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
  },
  nameCell: {
    display: 'flex',
    alignItems: 'center'
  },
  avatar: {
    height: 42,
    width: 42,
    marginRight: theme.spacing(2)
  }
}));

const GET_INDIVIDUALS = gql`
  query Getindividuals($name: String!, $page: Page!) {
    individuals(filter: { name: $name }, page: $page) {
      edges {
        uuid
        email
        firstName
        lastName
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

function IndividualsList() {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data, refetch } = useQuery(GET_INDIVIDUALS, {
    variables: {
      name: searchText,
      page: {
        offset: page,
        limit: rowsPerPage,
      },
    },
    fetchPolicy: 'cache-and-network'
  });

  if (error) return <div>{error.message}</div>;

  const handleNewIndividual = data => {
    if (!data.createIndividual) return;
    refetch();
  };

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  // Results table
  const headers = ['Email', 'First Name', 'Last Name', 'Created At'];
  const cells = [
    {
      component: (result) => (
        <div className={classes.nameCell}>
          <Avatar className={classes.avatar} src={result.logo} />
          <Link
            color="inherit"
            component={RouterLink}
            to={`/individuals/${result.uuid}/overview`}
            variant="h6"
          >
            {result.email}
          </Link>
        </div>
      )
    },
    { field: 'firstName' }, { field: 'lastName' },
    { component: (result) => moment(result.createdAt).format('LLL')}
  ]

  return (
    <Page className={classes.root} title="Individuals">
      <Container maxWidth={false}>
        <Header onCreateNewIndividual={handleNewIndividual} />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.individuals}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default IndividualsList;
