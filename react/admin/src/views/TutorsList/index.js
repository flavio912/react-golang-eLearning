import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
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

function TutorsList() {
  const classes = useStyles();

  // const handleFilter = () => {};

  const handleSearch = () => {};

  const tutors = [
    {
      fullName: 'Alexandria Tutor',
      email: 'alexandria@tutor.com',
      lastLogin: {
        date: '10/2/2020'
      },
      createdAt: '4/1/2020',
      cin: '1000345323898'
    }
  ];

  return (
    <Page className={classes.root} title="Tutors">
      <Container maxWidth={false}>
        <Header />
        <SearchBar onFilter={false} onSearch={handleSearch} />
        {tutors && <Results className={classes.results} tutors={tutors} />}
      </Container>
    </Page>
  );
}

export default TutorsList;
