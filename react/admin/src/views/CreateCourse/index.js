import React, { useState, useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import { makeStyles } from '@material-ui/styles';
import {
  Container,
  Tabs,
  Tab,
  Divider,
  CircularProgress
} from '@material-ui/core';
import Page from 'src/components/Page';
import Header from './Header';
import About from './About';
import Overview from './Overview';
import Pricing from './Pricing';
import CourseBuilder from './CourseBuilder';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';

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

const SAVE_ONLINE_COURSE = gql`
  mutation SaveOnlineCourse(
    $id: Int
    $name: String
    $excerpt: String
    $price: Float
    $backgroundCheck: Boolean
    $accessType: AccessType
    $howToComplete: String
    $hoursToComplete: Float
    $whatYouLearn: [String!]
    $requirements: [String!]
    $bannerImageSuccess: String
    $categoryUUID: UUID
  ) {
    saveOnlineCourse(
      input: {
        id: $id
        name: $name
        excerpt: $excerpt
        price: $price
        backgroundCheck: $backgroundCheck
        accessType: $accessType
        howToComplete: $howToComplete
        hoursToComplete: $hoursToComplete
        whatYouLearn: $whatYouLearn
        requirements: $requirements
        bannerImageSuccess: $bannerImageSuccess
        categoryUUID: $categoryUUID
      }
    ) {
      id
    }
  }
`;

const GET_COURSE = gql`
  query GetCourse($id: Int!) {
    course(id: $id) {
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
      bannerImageURL
      category {
        name
        uuid
      }
    }
  }
`;

var initState = {
  name: '',
  category: {},
  secondaryCategory: {},
  tags: [],
  excerpt: '',
  courseType: 'online',
  accessType: 'open',
  backgroundCheck: false,
  howToComplete: '',
  whatYouLearn: [],
  requirements: [],
  hoursToComplete: 0,
  price: 0,
  priceType: 'paid',
  bannerImageURL: undefined,
  bannerImageSuccess: undefined,
  categoryUUID: undefined
};

function CreateCourse({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'about', label: 'About' },
    { value: 'builder', label: 'Course Builder' },
    { value: 'pricing', label: 'Pricing' }
  ];

  const handleTabsChange = (event, value) => {
    history.push(value);
  };

  const [state, setState] = useState(initState);
  const [saveOnlineCourse] = useMutation(SAVE_ONLINE_COURSE);
  const { loading, error, data, refetch } = useQuery(GET_COURSE, {
    variables: {
      id: parseInt(ident)
    },
    fetchPolicy: 'cache-and-network',
    skip: !ident
  });
  useEffect(() => {
    if (loading || error) return;
    if (!data) return;
    setState({
      ...initState,
      name: data.course.name,
      excerpt: data.course.excerpt,
      backgroundCheck: data.course.backgroundCheck,
      courseType: data.course.type,
      accessType: data.course.accessType,
      howToComplete: data.course.howToComplete,
      hoursToComplete: data.course.hoursToComplete,
      whatYouLearn: data.course.whatYouLearn,
      requirements: data.course.requirements,
      bannerImageURL: data.course.bannerImageURL,
      price: data.course.price,
      categoryUUID: data.course.categoryUUID
    });
  }, [data, loading, error]);

  const updateState = stateUpdates => {
    var updatedState = { ...state, ...stateUpdates };
    setState(updatedState);
  };

  if (ident) {
    if (loading) return <CircularProgress className={classes.centerProgress} />;
    if (error) return <div>{error.message}</div>;
  }

  const saveDraft = async () => {
    if (state.courseType === 'online') {
      console.log('updating', state);
      const { data } = await saveOnlineCourse({
        variables: {
          id: ident ? parseInt(ident) : undefined,
          name: state.name,
          excerpt: state.excerpt,
          backgroundCheck: state.backgroundCheck,
          accessType: state.accessType,
          howToComplete: state.howToComplete,
          hoursToComplete: state.hoursToComplete,
          whatYouLearn: state.whatYouLearn,
          requirements: state.requirements,
          bannerImageSuccess: state.bannerImageSuccess,
          price: state.price
        }
      });

      // If a new course go to edit page
      if (!ident) {
        history.push(`/course/${data.saveOnlineCourse.id}/overview`);
        return;
      }
    }
    refetch();
  };

  // if (!currentTab) {
  //   return <Redirect to={`/courses/create/overview`} />;
  // }

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
          {currentTab === 'about' && (
            <About state={state} setState={updateState} />
          )}
          {currentTab === 'pricing' && (
            <Pricing state={state} setState={updateState} />
          )}
          {currentTab === 'builder' && <CourseBuilder />}
        </div>
      </Container>
    </Page>
  );
}

export default CreateCourse;
