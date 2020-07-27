import React, { useState } from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
  Typography,
  Switch
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import TagsInput from 'src/components/TagsInput';
import UploadFile from 'src/components/UploadFile';
import { gql } from 'apollo-boost';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  buttonText: {
    color: '#4a4a4a',
    fontSize: 11,
    fontWeight: 'weight: 700'
  },
  shortDescription: {
    width: '100%'
  },
  answerItem: {
    border: '1px solid gainsboro',
    borderRadius: 3,
    padding: '6px 21px',
    justifyContent: 'space-between',
    alignItems: 'center',
    display: 'flex'
  },
  formControl: {
    width: '100%'
  },
  previewImage: {
    width: 200,
    maxHeight: 200
  },
  editItems: {
    alignItems: 'center'
  }
}));

const UPLOAD_REQUEST = gql`
  mutation UploadRequest($fileType: String!, $contentLength: Int!) {
    certificateBodyImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      url
      successToken
    }
  }
`;

function Overview({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={2} direction={'column'}>
            <Grid item>
              <Card>
                <CardHeader title={'Certificate Info'} />
                <Divider />
                <CardContent>
                  <Grid container direction={'column'} spacing={2}>
                    <Grid item>
                      <TextField
                        label="Certificate Type Name"
                        value={state.name}
                        onChange={inp => {
                          setState({ name: inp.target.value });
                        }}
                        fullWidth
                        variant="outlined"
                      />
                    </Grid>
                    <Grid item>
                      <TextField
                        label="Regulation Text (optional)"
                        value={state.regulationText}
                        onChange={inp => {
                          setState({ regulationText: inp.target.value });
                        }}
                        fullWidth
                        variant="outlined"
                      />
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4} container spacing={2} direction={'column'}>
          <Grid item>
            <Card>
              <CardHeader title={'Additional options'} />
              <Divider />
              <CardContent>
                <Grid container direction={'column'} spacing={3}>
                  <Grid item>
                    <Typography variant="overline">Regulator Image</Typography>
                    <Typography variant="body2">
                      The image shown in the top left of a certificate
                    </Typography>
                    {state.certificateBodyImageURL && (
                      <img
                        alt="preview"
                        src={state.certificateBodyImageURL}
                        className={classes.previewImage}
                      />
                    )}
                    <UploadFile
                      uploadMutation={UPLOAD_REQUEST}
                      onUploaded={(token, url) => {
                        setState({
                          certificateBodyToken: token,
                          certificateBodyImageURL: url
                        });
                      }}
                    />
                  </Grid>
                  <Grid item>
                    <Typography variant="overline">
                      Requires CAA Number
                    </Typography>
                    <Typography variant="body2">
                      If checked certificates will have a CAA number attached.
                    </Typography>
                    <Switch
                      checked={state.requiresCAANo}
                      color="secondary"
                      edge="start"
                      name="RequiresCAA"
                      onChange={(evt, checked) => {
                        setState({ requiresCAANo: checked });
                      }}
                      value={state.requiresCAANo}
                    />
                  </Grid>
                  <Grid item>
                    <Typography variant="overline">
                      Show training section
                    </Typography>
                    <Typography variant="body2">
                      Adds a second page to the PDF that shows modules taken
                    </Typography>
                    <Switch
                      checked={state.showTrainingSection}
                      color="secondary"
                      edge="start"
                      name="trainingSection"
                      onChange={(evt, checked) => {
                        setState({ showTrainingSection: checked });
                      }}
                      value={state.showTrainingSection}
                    />
                  </Grid>
                </Grid>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
