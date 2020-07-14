import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import BulletRepeater from './BulletRepeater';
import SideOptions from './SideOptions';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  buttonText: {
    color: '#4a4a4a',
    fontSize: 11,
    fontWeight: 'weight: 700'
  },
  toComplete: {
    width: '100%'
  }
}));

function About({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title={'How to complete this course (500 words)'} />
                <Divider />
                <CardContent>
                  <TextField
                    label=""
                    multiline
                    className={classes.toComplete}
                    rows={5}
                    value={state.howToComplete}
                    onChange={inp => {
                      setState('howToComplete', inp.target.value);
                    }}
                    placeholder={'How to complete this course (500 words)'}
                    variant="outlined"
                  />
                </CardContent>
              </Card>
            </Grid>
            <Grid item>
              <BulletRepeater
                title={"What you'll learn"}
                items={state['whatYouLearn']}
                onChange={items => setState('whatYouLearn', items)}
              />
            </Grid>
            <Grid item>
              <BulletRepeater
                title={'Requirements'}
                items={state['requirements']}
                onChange={items => setState('requirements', items)}
              />
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Grid container direction={'column'} spacing={2}>
            <Grid item>
              <SideOptions state={state} setState={setState} />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default About;
