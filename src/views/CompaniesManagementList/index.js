import React, { useState, useEffect } from 'react';
import uuid from 'uuid/v1';
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

  const handleFilter = () => {};

  const handleSearch = () => {};

  const companies = [
    {
      id: uuid(),
      name: 'FedEx',
      email: 'kate@fedex.com',
      logo:
        'https://cdn.cnn.com/cnnnext/dam/assets/180301124611-fedex-logo.png',
      noDelegates: 40,
      noManagers: 1,
      paymentType: 'Contract'
    },
    {
      id: uuid(),
      name: 'Royal Mail',
      email: 'user@royalmail.co.uk',
      logo:
        'https://upload.wikimedia.org/wikipedia/en/thumb/e/ee/Royal_Mail.svg/1200px-Royal_Mail.svg.png',
      noDelegates: 23,
      noManagers: 3,
      paymentType: 'Pay as you go'
    },
    {
      id: uuid(),
      name: 'DHL',
      email: 'manager@dhl.com',
      avatar: '/images/avatars/avatar_4.png',
      noDelegates: 23,
      noManagers: 2,
      paymentType: 'Contract'
    }
  ];

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
