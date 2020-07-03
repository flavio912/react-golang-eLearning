import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Container } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import Page from 'src/components/Page';
import Header from './Header';
import Filter from './Filter';
import Results from './Results';

const useStyles = makeStyles(theme => ({
  root: {
    paddingTop: theme.spacing(3),
    paddingBottom: theme.spacing(3)
  },
  header: {
    marginBottom: theme.spacing(3)
  },
  filter: {
    marginTop: theme.spacing(3)
  },
  results: {
    marginTop: theme.spacing(6)
  }
}));

const GET_COURSES = gql`
  query GetCourses {
    courses {
      edges {
        id
        type
        name
        price
        category {
          name
        }
      }
      pageInfo {
        given
      }
    }
  }
`;

function CoursesView() {
  const classes = useStyles();

  const { loading, error, data } = useQuery(GET_COURSES);

  if (loading) return <div>Loading</div>;
  if (error) return <div>{error.message}</div>;
  console.log(data);
  return (
    <Page className={classes.root} title="All Courses">
      <Container maxWidth="lg">
        <Header className={classes.header} />
        <Filter className={classes.filter} />
        <Results className={classes.results} courses={data.courses.edges} />
      </Container>
    </Page>
  );
}

export default CoursesView;
