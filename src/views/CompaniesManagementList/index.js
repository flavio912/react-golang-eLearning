import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import axios from 'src/utils/axios';
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

function CompaniesManagementList() {
  const classes = useStyles();
  const [companies, setCompanies] = useState([]);

  const handleFilter = () => {};

  const handleSearch = () => {};

  useEffect(() => {
    let mounted = true;

    const fetchCompanies = () => {
      axios.get('/api/companies').then(response => {
        if (mounted) {
          setCompanies(response.data.companies);
        }
      });
    };

    fetchCompanies();

    return () => {
      mounted = false;
    };
  }, []);

  return (
    <Page className={classes.root} title="Companies Management List">
      <Container maxWidth={false}>
        <Header />
        <SearchBar onFilter={handleFilter} onSearch={handleSearch} />
        {companies && (
          <Results className={classes.results} companies={companies} />
        )}
      </Container>
    </Page>
  );
}

export default CompaniesManagementList;
