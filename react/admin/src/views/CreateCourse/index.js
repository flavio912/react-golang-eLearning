import React, { useState } from 'react';
import { Redirect } from 'react-router-dom';
import { makeStyles } from '@material-ui/styles';
import { Container, Tabs, Tab, Divider } from '@material-ui/core';
import Page from 'src/components/Page';
import Header from './Header';
import About from './About';
import Overview from './Overview';
import Pricing from './Pricing';
import CourseBuilder from './CourseBuilder';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';

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

const SAVE_ONLINE_COURSE = gql`
  mutation SaveOnlineCourse(
    $id: Int
    $name: String!
    $excerpt: String
    $backgroundCheck: Boolean
    $accessType: AccessType
  ) {
    saveOnlineCourse(
      input: {
        id: $id
        name: $name
        excerpt: $excerpt
        backgroundCheck: $backgroundCheck
        accessType: $accessType
      }
    ) {
      id
    }
  }
`;

function CreateCourse({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'about', label: 'About' },
    { value: 'builder', label: 'Course Builder' },
    { value: 'pricing', label: 'Pricing' }
  ];

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  const [state, setState] = useState({
    name: '',
    primaryCategory: {},
    secondaryCategory: {},
    tags: [],
    excerpt: '',
    courseType: 'online',
    accessType: 'restricted',
    backgroundCheck: false
  });

  const updateState = (item, value) => {
    var updatedState = { ...state, [item]: value };
    setState(updatedState);
    console.log(updatedState);
  };

  const [saveOnlineCourse, { data }] = useMutation(SAVE_ONLINE_COURSE);

  const saveDraft = () => {
    if (state.courseType == 'online') {
      saveOnlineCourse({
        variables: {
          name: state.name,
          excerpt: state.excerpt,
          backgroundCheck: state.backgroundCheck,
          accessType: state.accessType
        }
      });
    }
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
        <Header onSaveDraft={saveDraft} />
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
            <Overview state={state} setState={updateState} />
          )}
          {currentTab === 'about' && <About />}
          {currentTab === 'pricing' && <Pricing />}
          {currentTab === 'builder' && <CourseBuilder />}
        </div>
      </Container>
    </Page>
  );
}

export default CreateCourse;
