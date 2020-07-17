import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
  Button,
  Typography,
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
  },
  filename: {
    display: 'inline',
    fontWeight: '600',
    marginLeft: theme.spacing(2)
  }
}));

function AudioVideo({ state, setState }) {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Card>
              <CardContent>
                <Grid container spacing={2} direction={'column'}>
                <CardHeader title="Audio Voiceover" />
                  <Grid item>
                    <Autocomplete
                      options={state.tags}
                      getOptionLabel={option => option.title}
                      onChange={(event, newValue) => {
                          setState('voiceoverSuccessToken', newValue.value);
                      }}
                      renderInput={params => (
                          <TextField
                              {...params}
                              label="Upload Voiceover"
                              variant="outlined"
                              className={classes.thinInput}
                          />
                      )}
                    />
                    
                  </Grid>
                  <Grid item>
                    <Button variant="contained">
                      Upload MP3
                    </Button>
                    <Typography
                      variant="body2"
                      color="textSecondary"
                      className={classes.filename}
                    >
                      {state.voiceoverSuccessToken}
                    </Typography>
                 </ Grid>
                  <Grid item>
                    <Divider />
                  </Grid>
                  <CardHeader title="Video Source" />
                  <Grid item>
                    <Autocomplete
                      options={state.tags}
                      getOptionLabel={option => option.title}
                      onChange={(event, newValue) => {
                          setState('video', {type: 'WISTIA', url: newValue.value});
                      }}
                      renderInput={params => (
                          <TextField
                              {...params}
                              label="Wisita URL"
                              variant="outlined"
                              className={classes.thinInput}
                          />
                      )}
                    />
                  </Grid>
                  <Grid item>
                  <Typography
                      variant="body2"
                      color="textSecondary"
                      className={classes.filename}
                    >
                      Select your preferred video type (.mp4, Youtube Viemo, etc). Note when adding a video both the Audio Player + Module Banner Image will be hidden
                    </Typography>
                  </Grid>
                  <Grid item>
                    <TextField
                      fullWidth
                      label=""
                      name="modulename"
                      onChange={inp => {
                          setState('voiceoverSuccessToken', inp.target.value);
                      }}
                      placeholder="Enter Wisita Video URL"
                      value={state.voiceoverSuccessToken}
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
                  value={state.transcript}
                  onChange={inp => {
                      setState('transcript', inp.target.value);
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
