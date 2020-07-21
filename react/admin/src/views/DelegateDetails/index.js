import React from 'react';
import { Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import { Container, Tabs, Tab, Divider, colors } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

import Page from 'src/components/Page';
import Header from './Header';
import Overview from './Overview';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  tabs: {
    marginTop: theme.spacing(3)
  },
  divider: {
    backgroundColor: colors.grey[300]
  },
  content: {
    marginTop: theme.spacing(3)
  }
}));

const GET_DELEGATE = gql`
  query GetDelegate($id: UUID!) {
    delegate(uuid: $id) {
      uuid
      firstName
      lastName
      jobTitle
      email
      telephone
      company {
        uuid
        name
      }
    }
  }
`;

function DelegateDetails({ match, history }) {
  const classes = useStyles();
  const { id, tab: currentTab } = match.params;
  const tabs = [{ value: 'overview', label: 'Overview' }];

  const { loading, error, data } = useQuery(GET_DELEGATE, {
    variables: {
      id: id
    },
    fetchPolicy: 'cache-and-network',
    skip: !id
  });
  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;
  const delegate = data?.delegate;

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  if (!currentTab) {
    return <Redirect to={`/delegates/${id}/overview`} />;
  }

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  return (
    <Page className={classes.root} title="Delegate Details">
      <Container maxWidth={false}>
        <Header />
        <Tabs
          className={classes.tabs}
          onChange={handleTabsChange}
          scrollButtons="auto"
          value={currentTab}
          variant="scrollable"
        >
          {tabs.map(tab => (
            <Tab key={tab.value} label={tab.label} value={tab.value} />
          ))}
        </Tabs>
        <Divider className={classes.divider} />
        <div className={classes.content}>
          {currentTab === 'overview' && <Overview delegate={delegate} />}
        </div>
      </Container>
    </Page>
  );
}

DelegateDetails.propTypes = {
  history: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired
};

export default DelegateDetails;
