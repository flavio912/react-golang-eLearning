import React from 'react';
import { Redirect } from 'react-router-dom';
import PropTypes from 'prop-types';
import uuid from 'uuid/v4';
import { makeStyles } from '@material-ui/styles';
import { Container, Tabs, Tab, Divider, colors } from '@material-ui/core';
import Page from 'src/components/Page';
import Header from './Header';
import Summary from './Summary';
import Invoices from './Invoices';
import Users from './Users';
import Logs from './Logs';

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

function CompanyManagementDetails({ match, history }) {
  const classes = useStyles();
  const { id, tab: currentTab } = match.params;
  const tabs = [
    { value: 'summary', label: 'Summary' },
    { value: 'Courses', label: 'Courses' },
    { value: 'users', label: 'Users' },
    { value: 'invoices', label: 'Invoices' },
    { value: 'logs', label: 'Logs' }
  ];

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  if (!currentTab) {
    return <Redirect to={`/companies/${id}/summary`} />;
  }

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  const company = {
    id: uuid(),
    name: 'FedEx',
    email: 'kate@fedex.com',
    logo: 'https://cdn.cnn.com/cnnnext/dam/assets/180301124611-fedex-logo.png',
    noDelegates: 40,
    noManagers: 1,
    paymentType: 'Contract'
  };

  return (
    <Page className={classes.root} title="Company Management Details">
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
          {currentTab === 'summary' && <Summary />}
          {currentTab === 'invoices' && <Invoices />}
          {currentTab === 'users' && <Users company={company} />}
          {currentTab === 'logs' && <Logs />}
        </div>
      </Container>
    </Page>
  );
}

CompanyManagementDetails.propTypes = {
  history: PropTypes.object.isRequired,
  match: PropTypes.object.isRequired
};

export default CompanyManagementDetails;
