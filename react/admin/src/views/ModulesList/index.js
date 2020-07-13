import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import { gql } from 'apollo-boost';
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

const GET_MODULES = gql`
  query GetModules($id: Int!) {
    modules(id: $id) {
      id
      name
      excerpt
      price
      accessType
      backgroundCheck
      type
      howToComplete
      hoursToComplete
      whatYouLearn
      requirements
    }
  }
`;

function ModulesList({ match, history }) {
  const classes = useStyles();

  // const handleFilter = () => {};

  const handleSearch = () => {};

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

  return (
    <Page className={classes.root} title="Modules">
      <Container maxWidth={false}>
        <Header onAdd={() => {
            history.push('/modules/create/overview');
          }}/>
        <SearchBar onFilter={false} onSearch={handleSearch} />
        {modules && <Results className={classes.results} modules={modules} />}
      </Container>
    </Page>
  );
}

export default ModulesList;
