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
import Excerpt from './Excerpt';
import Introduction from './Introduction';
import Learn from './Learn';

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
              <About />
            </Grid>
            <Grid item>
              <Excerpt />
            </Grid>
            <Grid item>
              <Introduction />
            </Grid>
            <Grid item>
              <Learn />
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
