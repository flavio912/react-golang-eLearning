import React from 'react';
import uuid from 'uuid/v1';
import { Redirect } from 'react-router-dom';
import { makeStyles } from '@material-ui/styles';
import { Container, Tabs, Tab, Divider } from '@material-ui/core';
import Page from 'src/components/Page';
import Header from './Header';
import About from './About';
import Overview from './Overview';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  divider: {
    marginBottom: theme.spacing(3)
  },
  tabs: {
    marginTop: theme.spacing(3)
  }
}));
function CreateCourse({ match, history }) {
  const classes = useStyles();

  const { id, tab: currentTab } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'about', label: 'About' }
  ];

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  if (!currentTab) {
    return <Redirect to={`/courses/create/overview`} />;
  }

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  return (
    <Page className={classes.root} title="Create Course">
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
          {currentTab === 'overview' && <Overview />}
          {currentTab === 'about' && <About />}
        </div>
      </Container>
    </Page>
  );
}

export default CreateCourse;
