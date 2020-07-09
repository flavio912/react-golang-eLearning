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
  },
  termsInput: {
    width: '100%'
  }
}));

function Pricing({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title={'Terms and conditions'} />
                <Divider />
                <CardContent>
                  <TextField
                    label=""
                    multiline
                    className={classes.termsInput}
                    rows={5}
                    value={state.terms}
                    onChange={inp => {
                      setState('terms', inp.target.value);
                    }}
                    placeholder={'Terms and conditions'}
                    variant="outlined"
                  />
                </CardContent>
              </Card>
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

export default Pricing;
