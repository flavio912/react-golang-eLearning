import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Button, Chip } from '@material-ui/core';
import CreateOutlinedIcon from '@material-ui/icons/CreateOutlined';
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
  },
  row: {
    display: 'flex',
  },
  icon: {
    marginTop: 7,
    marginRight: 5,
  }
}));

const GET_MODULES = gql`
  query SearchModules($name: String!, $page: Page!) {
    modules(filter: { name: $name }, page: $page) {
      edges {
        uuid
        name
        syllabus {
          name
          uuid
        }
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

function ModulesList({ match, history }) {
  const classes = useStyles();
  const [searchText, setSearchText] = React.useState('');
  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_MODULES, {
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

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  // Results table
  const headers = ['Name', 'Courses Linked', 'Lessons', 'Tags', 'Actions'];
  const cells = [
    { field: 'name' }, { field: 'numCoursesUsedIn' },
    { component: (result) =>  result.syllabus.length},
    { component: (result) => (
      result.tags && result.tags.map(tag => (
        <Chip style={{ backgroundColor: tag.color }} label={tag.name} />
      ))
    )},
    { component: (result) => (
      <Button
        color="default"
        component={RouterLink}
        size="small"
        to={`/modules/${result.uuid}/overview`}
      >
        <CreateOutlinedIcon />
      </Button>
    )}
  ]

  return (
    <Page className={classes.root} title="Modules">
      <Container maxWidth={false}>
        <Header
          onAdd={() => {
            history.push('/modules/create/overview');
          }}
        />
        <SearchBar setSearchText={setSearchText} />
        <Results
          className={classes.results}
          results={data?.modules}
          headers={headers}
          cells={cells}
          handleChangePage={handleChangePage}
          handleChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Container>
    </Page>
  );
}

export default ModulesList;
