import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Button, Chip } from '@material-ui/core';
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

const GET_LESSONS = gql`
  query GetLessons($page: Page, $filter: LessonFilter, $orderBy: OrderBy) {
    lessons(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        name
        tags {
          uuid
          name
          color
        }
      }
      pageInfo {
        total
        limit
        offset
        given
      }
    }
  }
`;

function LessonsList({ history }) {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_LESSONS, {
    variables: {
      page: {
        offset: page,
        limit: rowsPerPage
      },
      filter: {
        name: searchText,
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
    const headers = ['Name', 'Courses Linked', 'Modules Linked', 'Tags', 'Actions'];
    const cells = [
      { field: 'name' },{ field: 'numCoursesLinked' },{ field: 'numModulesLinked' },
      {
        component: result =>
          result.tags &&
          result.tags.map(tag => (
            <Chip
              style={{ backgroundColor: tag.color, color: 'white' }}
              label={tag.name}
            />
          ))
      },
      {
        component: result => (
          <Button
            color="default"
            component={RouterLink}
            size="small"
            to={`/lesson/${result.uuid}/overview`}
          >
            Edit
          </Button>
        )
      }
    ];

  return (
    <Page className={classes.root} title="Lessons">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/lessons/create/overview');
          }}
        />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.lessons}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default LessonsList;
