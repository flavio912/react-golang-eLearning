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
  tabs,
  state,
  setState,
  currentTab,
  error,
  onSave,
  history,
  title
}) {
  const classes = useStyles();

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  return (
    <Page className={classes.root} title={title}>
      <ErrorModal error={error} />
      <Container maxWidth={false}>
        <Header onSave={onSave} title={title} />
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
            <Overview state={state} setState={setState} />
          )}
          {currentTab === 'audiovideo' && (
            <AudioVideo state={state} setState={setState} />
          )}
          {currentTab === 'modulebuilder' && (
            <ModuleBuilder state={state} setState={setState} />
          )}
        </div>
      </Container>
    </Page>
  );
}

export default ModulePage;
