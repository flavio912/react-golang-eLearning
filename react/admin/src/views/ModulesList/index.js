import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import { gql } from 'apollo-boost';
import { useQuery } from '@apollo/react-hooks';
import Page from 'src/components/Page';
import Header from './Header';
import Results from './Results';
import Filter from 'src/components/Filter';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(3)
  },
  filter: {
    marginTop: theme.spacing(3)
  },
}));

const GET_MODULES = gql`
  query GetModules($page: Page, $filter: ModuleFilter, $orderBy: OrderBy) {
    modules(page: $page, filter: $filter, orderBy: $orderBy) {
      edges {
        uuid
        name
      }
      pageInfo {
        given
        total
      }
    }
  }
`;

function ModulesList({ match, history }) {
  const classes = useStyles();

  // const handleFilter = () => {};

  //const handleSearch = () => {};
  const { loading, error, data } = useQuery(GET_MODULES, {
    variables: {
      page: {
        offset: 0,
        limit: 100
      },
      filter: {},
      orderBy: {
        ascending: false,
        field: 'created_at'
      }
    }
  });
  console.log(data, error);

  const modules = [
    {
      uuid: '1231231231',
      name: 'test module',
      numCoursesUsedIn: 3,
      numLessons: 5,
      tags: [
        {
          name: 'cool tag',
          color: '#123'
        }
      ]
    }
  ];

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;

  return (
    <Page className={classes.root} title="Modules">
      <Container maxWidth={false}>
        <Header onAdd={() => {
            history.push('/modules/create/overview');
          }}/>
        <Filter className={classes.filter} />
        {modules && <Results className={classes.results} modules={data.modules.edges} />}
      </Container>
    </Page>
  );
}

export default ModulesList;
