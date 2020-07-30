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
import { gql } from 'apollo-boost';

import TagsInput from 'src/components/TagsInput';
import UploadFile from 'src/components/UploadFile';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  lessonName: {
    marginBottom: 20
  },
  shortDescription: {
    width: '100%'
  },
  preview: {
    height: 300,
    width: '100%'
  }
}));

const UPLOAD_REQUEST = gql`
  mutation UploadRequest($fileType: String!, $contentLength: Int!) {
    lessonBannerImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      url
      successToken
    }
  }
`;

function Overview({ state, setState }) {
  const classes = useStyles();

  const onUploaded = (token, url) => {
    setState({ bannerImageURL: url, bannerImageToken: token });
  };

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item sm={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title={'About this Lesson'} />
                <Divider />
                <CardContent>
                  <Grid item className={classes.lessonName}>
                    <TextField
                      fullWidth
                      label="Name"
                      name="lesson"
                      value={state.name}
                      onChange={inp => {
                        setState({ name: inp.target.value });
                      }}
                      placeholder="Lesson Name"
                      variant="outlined"
                    />
                  </Grid>
                  <TagsInput
                    onChange={tags => {
                      setState({ tags });
                    }}
                  />
                </CardContent>
              </Card>
            </Grid>
            <Grid item>
              <Card>
                <CardHeader title="Lesson Description" />
                <Divider />
                <CardContent>
                  <Grid item>
                    <TextField
                      label="Description"
                      multiline
                      className={classes.shortDescription}
                      rows={5}
                      value={state.description}
                      onChange={inp => {
                        setState({ description: inp.target.value });
                      }}
                      placeholder="Lesson Description"
                      variant="outlined"
                    />
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Card>
            <CardHeader title="Lessons banner Image" />
            <Divider />
            <CardContent>
              {state.bannerImageURL && (
                <img
                  src={state.bannerImageURL}
                  className={classes.preview}
                  alt="preview"
                />
              )}
              <UploadFile
                uploadMutation={UPLOAD_REQUEST}
                onUploaded={onUploaded}
              />
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
