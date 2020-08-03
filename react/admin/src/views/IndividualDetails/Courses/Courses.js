import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { Link as RouterLink } from 'react-router-dom';
import { makeStyles } from '@material-ui/styles';
import { Link } from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import Results from 'src/components/Results';
import Page from 'src/components/Page';

const useStyles = makeStyles(theme => ({
  root: {
    wisth: '100%'
  },
}));

const GET_MY_COURSES = gql`
  query GetCourses($uuid: UUID!, $page: Page) {
    individuals(filter: { uuid: $uuid }, page: $page) {
      edges {
        myCourses {
          status
          course {
            id
            name
          }
          minutesTracked
          enrolledAt
          upTo
        }
      }
      pageInfo {
        total
        limit
        offset
        given
      }
    }
  }
`;

function Courses({ individual, className, ...rest }) {
  const classes = useStyles();

  const [page, setPage] = React.useState(0);
  const [rowsPerPage, setRowsPerPage] = React.useState(10);
  const { error, data } = useQuery(GET_MY_COURSES, {
    variables: {
      uuid: individual.uuid,
      page: {
        offset: page,
        limit: rowsPerPage
      }
    },
    fetchPolicy: 'cache-and-network'
  });
  
  if (error) return <div>{error.message}</div>;

  const extractCourses = (array) => {
    const myCourses = [];
    const courses = array?.individuals?.edges.map((ind) => (
        ind.myCourses.map(({status,course,minutesTracked,enrolledAt,upTo}) => (
          myCourses.push({
            status,course,minutesTracked,enrolledAt,upTo
          })
        ))
    ));
    console.log(courses);
    return {edges: myCourses, pageInfo: array?.individuals?.pageInfo};
  };

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  const headers = ['Course', 'Status', 'Minutes', 'Enrolled At', 'Up to'];
  const cells = [
    {
      component: (result) => (
        <Link
          color="inherit"
          component={RouterLink}
          to={`/course/${result.course.id}/overview`}
          variant="h6"
        >
          {result.course.name}
        </Link>
      )
    },
    {field: 'status'},{field: 'minutesTracked'},
    {field: 'enrolledAt'},{field: 'upTo'},
  ];

  return (
    <Page {...rest} className={clsx(classes.root, className)}>
      <Results
        results={extractCourses(data)}
        headers={headers}
        cells={cells}
        handleChangePage={handleChangePage}
        handleChangeRowsPerPage={handleChangeRowsPerPage}
      />
    </Page>
  );
}

Courses.propTypes = {
  className: PropTypes.string,
  individual: PropTypes.object.isRequired
};

export default Courses;
