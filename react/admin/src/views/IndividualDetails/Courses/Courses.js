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
  }
}));

const GET_MY_COURSES = gql`
  query GetCourses($uuid: UUID!) {
    individual(uuid: $uuid) {
      myCourses {
        status
        course {
          id
          name
        }
        minutesTracked
        enrolledAt
        upTo
        certificateURL
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

  const extractCourses = array => {
    const myCourses = array?.individual?.myCourses ?? [];
    return { edges: myCourses, pageInfo: array?.individuals?.pageInfo };
  };

  // Results methods
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  const headers = ['Course', 'Status', 'Enrolled At', 'Certificate'];
  const cells = [
    {
      component: result => (
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
    { field: 'status' },
    { field: 'enrolledAt' },
    {
      component: result =>
        result.certificateURL ? (
          <Link
            href={result.certificateURL}
            target={'_blank'}
            className={classes.certLink}
          >
            Certificate
          </Link>
        ) : (
          <div>Not Available</div>
        )
    }
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
