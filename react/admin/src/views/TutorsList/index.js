import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Link } from '@material-ui/core';
import gql from 'graphql-tag';
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
  },
  imageContainer: {
    display: 'flex',
    justifyContent: 'left'
  },
  image: {
    maxWidth: 200,
    width: 'auto',
    maxHeight: 70,
    height: 'auto'
  }
}));

const GET_TUTORS = gql`
  query GetTutors($page: Page, $filter: TutorFilter, $orderBy: OrderBy) {
    tutors(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        name
        cin
        signatureURL
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

function TutorsList() {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_TUTORS, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage
      },
      filter: {
        name: searchText
      },
      orderBy: {
        ascending: true,
        field: 'name'
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
  const headers = ['User', 'CIN Number', 'Signature URL'];
  const cells = [
    {
      component: result => (
        <div>
          <Link
            color="inherit"
            component={RouterLink}
            to={`/tutors/${result.uuid}/overview`}
            variant="h6"
          >
            {result.name}
          </Link>
        </div>
      )
    },
    { field: 'cin' },
    {
      component: result => (
        <div className={classes.imageContainer}>
          <img
            alt="Signature"
            className={classes.image}
            src={result.signatureURL}
          />
        </div>
      )
    }
  ];

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
