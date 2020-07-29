import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Avatar, Link } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import { Link as RouterLink } from 'react-router-dom';
import Page from 'src/components/Page';
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

const GET_ADMINS = gql`
  query GetAdmins($page: Page!) {
    admins(page: $page) {
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
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data, refetch } = useQuery(GET_ADMINS, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage,
      },
    },
    fetchPolicy: 'cache-and-network'
  });

  if (error) return <div>{error.message}</div>;

  const handleNewAdmin = data => {
    if (!data.createAdmin) return;
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
  const headers = ['Email', 'First Name', 'Last Name'];
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
    { field: 'firstName' },
    { field: 'lastName' },
  ]

  return (
    <Page className={classes.root} title="Admins">
      <Container maxWidth={false}>
        <Header onCreateNewAdmin={handleNewAdmin} />
        <Results
          className={classes.results}
          results={data?.admins}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default AdminsList;
