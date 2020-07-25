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
    moduleBannerImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      successToken
      url
    }
  }
`;

function Overview({ state, setState }) {
  const classes = useStyles();

  const onUploaded = (token, url) => {
    setState({ bannerImageURL: url, bannerImageSuccessToken: token });
  };
  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title="About this Module" />
                <Divider />
                <CardContent>
                  <Grid container spacing={4} direction={'column'}>
                    <Grid item>
                      <TextField
                        fullWidth
                        label="Module Name"
                        name="modulename"
                        onChange={inp => {
                          setState({ name: inp.target.value });
                        }}
                        placeholder="e.g. Fire Safety Module 1"
                        value={state.name}
                        variant="outlined"
                      />
                    </Grid>
                    <Grid item>
                      <TagsInput
                        allowNew
                        onChange={tags => {
                          setState({ tags: tags });
                          console.log('tasgs: ', tags);
                        }}
                      />
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
            <Grid item>
              <Card>
                <CardHeader title="Module Description" />
                <Divider />
                <CardContent>
                  <TextField
                    label=""
                    multiline
                    className={classes.shortDescription}
                    rows={5}
                    value={state.description}
                    onChange={inp => {
                      setState({ description: inp.target.value });
                    }}
                    placeholder="Description"
                    variant="outlined"
                  />
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Card>
            <CardHeader title="Module Banner Image" />
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
