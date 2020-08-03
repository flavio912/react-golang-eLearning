import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Typography,
  Switch,
  InputAdornment,
  TextField
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import TagsInput from 'src/components/TagsInput';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  }
}));

function Overview({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title={'Test Information'} />
                <Divider />
                <CardContent>
                  <Grid container spacing={2} direction={'column'}>
                    <Grid item>
                      <TextField
                        label="Test Name"
                        fullWidth
                        rows={5}
                        value={state.name}
                        onChange={inp => {
                          setState('name', inp.target.value);
                        }}
                        placeholder={'Test name'}
                        variant="outlined"
                      />
                    </Grid>
                    <Grid item>
                      <TagsInput />
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Card>
            <CardHeader title={'Test settings'} />
            <CardContent>
              <Grid container spacing={2} direction={'column'}>
                <Grid item>
                  <TextField
                    label="Attempts allowed"
                    fullWidth
                    value={state.attemptsAllowed}
                    onChange={inp => {
                      setState('attemptsAllowed', inp.target.value);
                    }}
                    type="number"
                    placeholder={'e.g 3'}
                    variant="outlined"
                  />
                </Grid>
                <Grid item>
                  <TextField
                    label="Pass Percentage"
                    fullWidth
                    value={state.passPercentage.toString()}
                    onChange={inp => {
                      try {
                        setState(
                          'passPercentage',
                          parseFloat(inp.target.value)
                        );
                      } catch (err) {}
                    }}
                    InputProps={{
                      endAdornment: (
                        <InputAdornment position="end">%</InputAdornment>
                      )
                    }}
                    type="number"
                    placeholder={'e.g 30'}
                    variant="outlined"
                  />
                </Grid>
                <Grid item>
                  <TextField
                    label="Num questions to answer"
                    fullWidth
                    value={state.questionsToAnswer.toString()}
                    onChange={inp => {
                      try {
                        setState(
                          'questionsToAnswer',
                          parseInt(inp.target.value)
                        );
                      } catch {}
                    }}
                    type="number"
                    placeholder={'e.g 3'}
                    variant="outlined"
                  />
                </Grid>
                <Grid item>
                  <Typography variant="h6">Randomise Answers</Typography>
                  <Switch
                    checked={state.randomiseAnswers}
                    onChange={evt => {
                      setState('randomiseAnswers', evt.target.checked);
                    }}
                    name="checkedA"
                  />
                </Grid>
              </Grid>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
