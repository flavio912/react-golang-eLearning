import React from 'react';
import { Grid } from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import SideOptions from './SideOptions';
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

function Pricing() {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <EditorCard title={'Standard Terms'} />
            </Grid>
            <Grid item>
              <EditorCard title={'Specific Terms'} />
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Grid container direction={'column'} spacing={2}>
            <Grid item>
              <SideOptions />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default Pricing;