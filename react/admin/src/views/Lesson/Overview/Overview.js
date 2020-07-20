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
  }
}));

const UPLOAD_REQUEST = gql`
  mutation UploadRequest($fileType: String!, $contentLength: Int!) {
    answerImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      url
      successToken
    }
  }
`;

function Overview({ state, setState }) {
  const classes = useStyles();

  const handleTags = () => {};

  return (
    <div className={classes.root}>
      <Grid container spacing={4}>
        <Grid item sm={8}>
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
                    setState('name', inp.target.value);
                  }}
                  placeholder="Lesson Name"
                  variant="outlined"
                />
              </Grid>
              <TagsInput onChange={handleTags} />
            </CardContent>
          </Card>
        </Grid>
        <Grid item sm={4}>
          <Card>
            <CardHeader title={'Lessons banner Image'} />
            <Divider />
            <CardContent>
              <UploadFile
                uploadMutation={UPLOAD_REQUEST}
                onUploaded={(successToken, url) => {}}
              />
            </CardContent>
          </Card>
        </Grid>
        <Grid item sm={8}>
          <Card>
            <CardHeader title={'Lesson Description'} />
            <Divider />
            <CardContent>
              <Grid item>
                <TextField
                  fullWidth
                  multiline
                  rows={8}
                  label="Description"
                  name="description"
                  value={state.text}
                  onChange={inp => {
                    setState('description', inp.target.value);
                  }}
                  placeholder="Lesson Description"
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

export default Overview;
