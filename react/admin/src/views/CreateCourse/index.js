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
    $structure: [CourseItem!]
    $expiresInMonths: Int
    $expirationToEndMonth: Boolean
    $certificateType: UUID
    $specificTerms: String
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
        structure: $structure
        expiresInMonths: $expiresInMonths
        expirationToEndMonth: $expirationToEndMonth
        specificTerms: $specificTerms
        certificateType: $certificateType
      }
    ) {
      id
    }
  }
`;

const SET_PUBLISHED = gql`
  mutation SetPublished($courseID: Int!, $published: Boolean) {
    setCoursePublished(courseID: $courseID, published: $published)
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
      published
      expiresInMonths
      expirationToEndMonth
      specificTerms
      certificateType {
        uuid
        name
      }
      category {
        name
        uuid
      }
      syllabus {
        uuid
        name
        type
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
  published: false,
  bannerImageURL: undefined,
  bannerImageSuccess: undefined,
  categoryUUID: { title: '', value: '' },
  syllabus: [],
  certificateType: { name: '', uuid: undefined },
  expiresInMonths: false,
  expirationToEndMonth: false
};

function CreateCourse({ match, history }) {
  const classes = useStyles();

  const { tab: currentTab, ident } = match.params;
  const tabs = [
    { value: 'overview', label: 'Overview' },
    { value: 'about', label: 'About' },
    { value: 'builder', label: 'Course Builder' },
    { value: 'pricing', label: 'Pricing + Certificates' }
  ];

  const handleTabsChange = (_, value) => {
    history.push(value);
  };

  const [state, setState] = useState(initState);
  const [saveOnlineCourse, { error: saveError }] = useMutation(
    SAVE_ONLINE_COURSE
  );
  const [setPublishedMutation, { error: publishError }] = useMutation(
    SET_PUBLISHED
  );
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
      published: data.course.published,
      bannerImageURL: data.course.bannerImageURL,
      price: data.course.price,
      category: data.course.category,
      syllabus: data.course.syllabus,
      expiresInMonths: data.course.expiresInMonths,
      expirationToEndMonth: data.course.expirationToEndMonth,
      specificTerms: data.course.specificTerms,
      certificateType: {
        name: data.course.certificateType?.name,
        uuid: data.course.certificateType?.uuid
      }
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
          categoryUUID: state.category?.uuid,
          price: state.price,
          structure: state.syllabus.map(({type, uuid}) => ({type, uuid})),
          expiresInMonths: state.expiresInMonths,
          expirationToEndMonth: state.expirationToEndMonth,
          certificateType: state.certificateType?.uuid,
          specificTerms: state.specificTerms
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

  const setPublish = async published => {
    if (!ident) {
      alert('Must save the course before publishing');
      return;
    }

    await setPublishedMutation({
      variables: {
        published: published,
        courseID: parseInt(ident)
      }
    });
    refetch();
  };

  if (!tabs.find(tab => tab.value === currentTab)) {
    return <Redirect to="/errors/error-404" />;
  }

  return (
    <Page className={classes.root} title="Create Course">
      <ErrorModal error={saveError || publishError || error} />
      <Container maxWidth={false}>
        <Header
          onSaveDraft={saveDraft}
          setPublish={setPublish}
          isSaved={ident ?? undefined}
          published={state.published}
        />
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
          {currentTab === 'builder' && (
            <CourseBuilder state={state} setState={updateState} />
          )}
        </div>
      </Container>
    </Page>
  );
}

export default CreateCourse;
