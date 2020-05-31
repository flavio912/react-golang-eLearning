import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  Typography,
  CardContent,
  Grid,
  Radio,
  RadioGroup,
  FormControlLabel,
  FormLabel,
  Divider,
  Switch
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles(theme => ({}));

function CourseFeatures() {
  const classes = useStyles();

  return (
    <Card>
      <CardHeader title={'Course Features'} />
      <CardContent>
        <Grid container direction="column" spacing={2}>
          <Grid item>
            <FormLabel component="legend">Course Type</FormLabel>
            <RadioGroup
              aria-label="gender"
              name="gender1"
              value={'online'}
              onChange={() => {}}
            >
              <FormControlLabel
                value="online"
                control={<Radio />}
                label="Online Course"
              />
              <FormControlLabel
                value="classroom"
                control={<Radio />}
                label="Classroom course"
              />
            </RadioGroup>
          </Grid>
          <Grid item>
            <FormLabel component="legend">Access Type</FormLabel>
            <RadioGroup
              aria-label="gender"
              name="gender1"
              value={'open'}
              onChange={() => {}}
            >
              <FormControlLabel
                value="restricted"
                control={<Radio />}
                label="Restricted"
              />
              <FormControlLabel
                value="open"
                control={<Radio />}
                label="Open Access"
              />
            </RadioGroup>
          </Grid>
          <Grid item>
            <Divider />
          </Grid>
          <Grid item>
            <Typography variant={'overline'}>BACKGROUND CHECK</Typography>
            <Typography variant={'h5'}>Background Check</Typography>
            <Typography variant={'body2'}>
              Do managers need to provide a 5 year background check of their
              delegates to take this course
            </Typography>
            <Switch
              checked={true}
              onChange={() => {}}
              color="primary"
              name="checkedB"
              inputProps={{ 'aria-label': 'primary checkbox' }}
            />
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default CourseFeatures;
