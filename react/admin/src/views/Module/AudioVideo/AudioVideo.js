import React, { useState } from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { Autocomplete } from '@material-ui/lab';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  thinInput: {
    width: '30%'
  },
  shortDescription: {
    width: '100%'
  }
}));

function AudioVideo({}) {
  const classes = useStyles();
  const onChange=()=>{};
  const [name, setName] = React.useState();
  const [description, setDescription] = React.useState();
  const categoryOptions = [{ title: 'Aviation Security', value: 'avsec' }];

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Card>
              <CardContent>
                <Grid container spacing={2} direction={'column'}>
                  <Grid item>
                    <CardHeader title="Audio Voiceover" />
                  </Grid>
                  <Grid item>
                    <Autocomplete
                      options={categoryOptions}
                      getOptionLabel={option => option.title}
                      onChange={(event, newValue) => {
                          onChange(newValue.value);
                      }}
                      renderInput={params => (
                          <TextField
                              {...params}
                              label=""
                              variant="outlined"
                              className={classes.thinInput}
                          />
                      )}
                    />
                  </Grid>
                  <Grid item>
                    <Divider />
                    <CardHeader title="Video Source" />
                  </Grid>
                  <Grid item>
                    <Autocomplete
                      options={categoryOptions}
                      getOptionLabel={option => option.title}
                      onChange={(event, newValue) => {
                          onChange(newValue.value);
                      }}
                      renderInput={params => (
                          <TextField
                              {...params}
                              label=""
                              variant="outlined"
                              className={classes.thinInput}
                          />
                      )}
                    />
                  </Grid>
                  <Grid item>
                    <div>Select</div>
                  </Grid>
                  <Grid item>
                    <TextField
                      fullWidth
                      label=""
                      name="modulename"
                      onChange={inp => {
                          onChange(inp.target.value);
                      }}
                      placeholder="Enter Wisita Video URL"
                      value={name}
                      variant="outlined"
                    />
                  </Grid>
                </Grid>
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Video Transcript" />
              <Divider />
              <CardContent>
                <TextField
                  label=""
                  multiline
                  className={classes.shortDescription}
                  rows={5}
                  value={description}
                  onChange={inp => {
                      setDescription(inp.target.value);
                  }}
                  placeholder="Description"
                  variant="outlined"
                />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default AudioVideo;
