import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Link } from '@material-ui/core';
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
  }
}));

function TutorsList() {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  // TODO: Add tutors query
  // const { error, data, refetch } = useQuery(GET_INDIVIDUALS, {
  //   variables: {
  //     name: searchText,
  //     page: {
  //       offset: page,
  //       limit: rowsPerPage,
  //     },
  //   },
  //   fetchPolicy: 'cache-and-network'
  // });
  // if (error) return <div>{error.message}</div>;

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  // Results table
  const headers = ['User', 'CIN Number', 'Last Login', 'Created At'];
  const cells = [
    {
      component: (result) => (
        <div>
          <Link
            color="inherit"
            component={RouterLink}
            to="/companies/1"
            variant="h6"
          >
            {result.fullName}
          </Link>
          <div>{result.email}</div>
        </div>
      )
    },
    { field: 'cin' },
    { component: (result) => moment(result.lastLogin.date).format('LLL')},
    { component: (result) => moment(result.createdAt).format('LLL')}
  ]

  const data = {
    tutors: {
      edges: [{
        fullName: 'Alexandria Tutor',
        email: 'alexandria@tutor.com',
        lastLogin: {
          date: '10/2/2020'
        },
        createdAt: '4/1/2020',
        cin: '1000345323898'
      }],
      pageInfo: {
        given: 1,
        total: 1,
        offset: 0,
        limit: 10,
      }
    }
  }

  return (
    <Page className={classes.root} title="Tutors">
      <Container maxWidth={false}>
        <Header />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.tutors}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default TutorsList;
