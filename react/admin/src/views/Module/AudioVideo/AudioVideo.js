import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
  Typography,
  Select,
  MenuItem
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import UploadFile from 'src/components/UploadFile';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  padding: {
    padding: theme.spacing(3),
    paddingTop: 0
  },
  margin: {
    margin: `0 ${theme.spacing(3)}px`
  },
  thinInput: {
    width: '30%'
  },
  shortDescription: {
    width: '100%'
  },
  filename: {
    display: 'inline',
    fontWeight: '400',
    fontStyle: 'italic',
    marginLeft: theme.spacing(1)
  }
}));

const UPLOAD_REQUEST = gql`
  mutation UploadRequest($fileType: String!, $contentLength: Int!) {
    voiceoverUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      url
      successToken
    }
  }
`;

const voiceoverOptions = [{ title: 'Upload Voiceover', uuid: '0' }];
const videoOptions = [{ title: 'Wistia URL', uuid: '0' }];

function AudioVideo({ state, setState }) {
  const classes = useStyles();
  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Card>
              <CardHeader title="Audio Voiceover" />
              <Grid
                container
                spacing={2}
                direction={'column'}
                className={classes.padding}
              >
                <Grid item>
                  <Select
                    value={voiceoverOptions[0]}
                    variant="outlined"
                    className={classes.thinInput}
                  >
                    {voiceoverOptions.map(option => (
                      <MenuItem key={option.uuid} value={option}>
                        {option.title}
                      </MenuItem>
                    ))}
                  </Select>
                </Grid>
                <Grid item>
                  <UploadFile
                    title="Upload MP3"
                    uploadMutation={UPLOAD_REQUEST}
                    acceptType="audio/mp3"
                    onUploaded={(successToken, url) => {
                      setState({
                        voiceoverSuccessToken: successToken,
                        voiceoverURL: url
                      });
                    }}
                  />
                </Grid>
              </Grid>
              <Divider className={classes.margin} />
              <CardHeader title="Video Source" />
              <Grid
                container
                spacing={2}
                direction={'column'}
                className={classes.padding}
              >
                <Grid item>
                  <Select
                    value={videoOptions[0]}
                    variant="outlined"
                    className={classes.thinInput}
                  >
                    {videoOptions.map(option => (
                      <MenuItem key={option.uuid} value={option}>
                        {option.title}
                      </MenuItem>
                    ))}
                  </Select>
                </Grid>
                <Grid item>
                  <Typography
                    variant="body2"
                    color="textPrimary"
                    className={classes.filename}
                  >
                    Select your preferred video type (.mp4, Youtube Viemo, etc).
                    Note when adding a video both the Audio Player + Module
                    Banner Image will be hidden
                  </Typography>
                </Grid>
                <Grid item>
                  <TextField
                    fullWidth
                    label=""
                    name="modulename"
                    onChange={inp => {
                      setState({
                        video: { type: 'WISTIA', url: inp.target.value }
                      });
                    }}
                    placeholder="Enter Wisita Video URL"
                    value={state.video.url}
                    variant="outlined"
                  />
                </Grid>
              </Grid>
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
                    setState({ transcript: inp.target.value });
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
