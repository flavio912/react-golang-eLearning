import React, { useState } from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
}));

function Overview({}) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Card>
              <CardHeader title="About this Module" />
              <Divider />
              <CardContent>
                <div>TBA</div>
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Module Description" />
              <Divider />
              <CardContent>
                <div>TBA</div>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
