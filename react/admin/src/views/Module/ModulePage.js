import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container, Tabs, Tab, Divider } from '@material-ui/core';
import Page from 'src/components/Page';
import Header from './Header';
import Overview from './Overview';
import AudioVideo from './AudioVideo';
import ModuleBuilder from './ModuleBuilder';
import ErrorModal from 'src/components/ErrorModal';

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
  },
  centerProgress: {
    position: 'absolute',
    top: '50%',
    left: '50%'
  }
}));

function ModulePage({
    history,
    match
}) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [
      { value: 'overview', label: 'Overview' },
      { value: 'audiovideo', label: 'Audio/Video' },
      { value: 'modulebuilder', label: 'Module Builder' }
    ];

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  const onSave = () => {};

  return (
    <Page className={classes.root} title="Test">
      <ErrorModal />
      <Container maxWidth={false}>
        <Header onSave={onSave} title="Test" />
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
          {currentTab === 'overview' && (
            <Overview />
          )}
          {currentTab === 'audiovideo' && (
            <AudioVideo />
          )}
          {currentTab === 'modulebuilder' && (
            <ModuleBuilder />
          )}
        </div>
      </Container>
    </Page>
  );
}

export default ModulePage;
