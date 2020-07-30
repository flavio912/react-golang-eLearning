import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import {
  Card,
  CardHeader,
  CardContent,
  Divider,
  Table,
  TableBody,
  TableRow,
  TableCell,
  Grid
} from '@material-ui/core';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';

const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  actions: {
    flexDirection: 'column',
    alignItems: 'flex-start',
    '& > * + *': {
      marginLeft: 0
    }
  },
  buttonIcon: {
    marginRight: theme.spacing(1)
  }
}));

const GET_MY_COURSES = gql`
  query GetCourses($uuid: UUID!) {
    individual(uuid: $uuid) {
      myCourses {
        status
        course {
          name
        }
        minutesTracked
        enrolledAt
        upTo
      }
    }
  }
`;

function Courses({ individual, className, ...rest }) {
  const classes = useStyles();
  console.log('ind', individual);
  const { error, data } = useQuery(GET_MY_COURSES, {
    variables: {
      uuid: individual.uuid,
    },
    fetchPolicy: 'cache-and-network'
  });
  console.log('data', data);

  if (error) return <div>{error.message}</div>;

  const tableData = [
    {header: 'Status', field: 'status'},
    {header: 'Minutes', field: 'minutesTracked'},
    {header: 'Enrolled At', field: 'enrolledAt'},
    {header: 'Up to', field: 'upTo'},
  ];

  return (
    <Grid container spacing={4} {...rest} className={clsx(classes.root, className)}>
      {data && data.individual.myCourses.map(course => (
        <Grid item>
          <Card>
            <CardHeader title={course.course.name} />
            <Divider />
            <CardContent className={classes.content}>
              <Table>
                <TableBody>
                  {tableData.map(row => (
                    <TableRow>
                      <TableCell>{row.header}</TableCell>
                      <TableCell>{data[row.field]}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </CardContent>
          </Card>
        </Grid>
      ))}
    </Grid>
  );
}

Courses.propTypes = {
  className: PropTypes.string,
  individual: PropTypes.object.isRequired
};

export default Courses;
