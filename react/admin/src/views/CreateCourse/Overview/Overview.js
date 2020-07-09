import React from 'react';
import {
  Grid,
  TextField,
  Card,
  CardHeader,
  CardContent,
  Divider
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import CourseFeatures from './CourseFeatures';
import CourseInfo from './CourseInfo';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  buttonText: {
    color: '#4a4a4a',
    fontSize: 11,
    fontWeight: 'weight: 700'
  },
  shortDescription: {
    width: '100%'
  }
}));

function Overview({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <CourseInfo state={state} setState={setState} />
            </Grid>
            <Grid item>
              <Card>
                <CardHeader title={'Short Description'} />
                <Divider />
                <CardContent>
                  <TextField
                    label=""
                    multiline
                    className={classes.shortDescription}
                    rows={5}
                    value={state.excerpt}
                    onChange={inp => {
                      setState('excerpt', inp.target.value);
                    }}
                    placeholder={'Short description'}
                    variant="outlined"
                  />
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <CourseFeatures state={state} setState={setState} />
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
