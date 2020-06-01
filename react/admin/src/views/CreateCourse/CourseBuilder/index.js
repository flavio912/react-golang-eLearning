import React from 'react';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  }
}));

function CourseBuilder() {
  const classes = useStyles();

  return <div className={classes.root}></div>;
}

export default CourseBuilder;
