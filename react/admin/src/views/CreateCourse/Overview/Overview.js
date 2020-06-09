import React from 'react';
import { Grid } from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import CourseFeatures from './CourseFeatures';
import CourseInfo from './CourseInfo';
import EditorCard from 'src/components/EditorCard';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  buttonText: {
    color: '#4a4a4a',
    fontSize: 11,
    fontWeight: 'weight: 700'
  }
}));

function Overview() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <CourseInfo />
            </Grid>
            <Grid item>
              <EditorCard inlineOnly title={'About this course (70 words)'} />
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <CourseFeatures />
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
