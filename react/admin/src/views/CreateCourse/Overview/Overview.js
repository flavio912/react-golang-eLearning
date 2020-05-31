import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  Typography,
  TextField,
  CardContent,
  Container,
  Button,
  Grid,
  Chip,
  ButtonGroup,
  GridList,
  Radio,
  RadioGroup,
  InputLabel,
  FormControlLabel,
  FormLabel,
  Divider,
  Switch
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import CourseFeatures from './CourseFeatures';
import About from './About';

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
          <About />
        </Grid>
        <Grid item xs={4}>
          <CourseFeatures />
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
