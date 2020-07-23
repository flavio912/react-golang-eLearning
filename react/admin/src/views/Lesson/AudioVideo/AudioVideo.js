import React, { useState } from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
  InputLabel,
  Select,
  MenuItem,
  FormControl,
  Typography,
  Box
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';

import UploadFile from 'src/components/UploadFile';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  formControl: {
    minWidth: 240
  },
  cardContent: {
    paddingTop: 0
  },
  uploadButton: {
    marginTop: theme.spacing(2)
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

function AudioVideo({ state, setState }) {
  const classes = useStyles();
  const [voiceOver, setVoiceOver] = useState('');
  const [wistiaUrl, setWistiaUrl] = useState('');

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item sm={12} spacing={4}>
          <Card>
            <CardHeader title={'Audio Voiceover'} />
            <CardContent className={classes.cardContent}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-simple-select-outlined-label">
                  Upload Voiceover
                </InputLabel>
                <Select
                  labelId="demo-simple-select-outlined-label"
                  id="demo-simple-select-outlined"
                  value={voiceOver}
                  onChange={event => {
                    setVoiceOver(event.target.value);
                  }}
                  label="Upload Voiceover"
                >
                  <MenuItem value={10}>Upload Voiceover</MenuItem>
                </Select>
              </FormControl>
              <Grid item sm={12} className={classes.uploadButton}>
                <UploadFile
                  uploadMutation={UPLOAD_REQUEST}
                  onUploaded={(successToken, url) => {
                    setState({ voiceoverSuccessToken: '' });
                  }}
                />
              </Grid>
            </CardContent>
            <Divider />
            <CardHeader title={'Video Source'} />
            <CardContent className={classes.cardContent}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel id="demo-simple-select-outlined-label">
                  Wistia URL
                </InputLabel>
                <Select
                  labelId="demo-simple-select-outlined-label"
                  id="demo-simple-select-outlined"
                  value={wistiaUrl}
                  onChange={event => {
                    setWistiaUrl(event.target.value);
                  }}
                  label="Wistia URL"
                >
                  <MenuItem value={10}>Wistia URL</MenuItem>
                </Select>
              </FormControl>
              <Typography>
                <Box fontStyle="italic" my={2}>
                  Select your preferred video type (.mp4, YouTube, Vimeo etc.).
                  Not when adding a video both the Audio Player + Module Banner
                  Image will be hidden.
                </Box>
              </Typography>
              <TextField
                fullWidth
                label="Wisita Video URL"
                name="wisita"
                onChange={inp => {}}
                placeholder="Enter Wisita Video URL"
                variant="outlined"
              />
            </CardContent>
          </Card>
        </Grid>
        <Grid item sm={12}>
          <Card>
            <CardHeader title={'Video Transcript'} />
            <Divider />
            <CardContent>
              <Grid item>
                <TextField
                  fullWidth
                  multiline
                  rows={8}
                  label="Transcript"
                  name="transcript"
                  onChange={inp => {
                    setState({ transcript: inp.target.value });
                  }}
                  variant="outlined"
                />
              </Grid>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </div>
  );
}

export default AudioVideo;
