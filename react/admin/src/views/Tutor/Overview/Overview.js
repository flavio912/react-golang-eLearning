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
import UploadFile from 'src/components/UploadFile';

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
    tutorSignatureImageUploadRequest(
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
    setState({ signatureURL: url, signatureToken: token });
  };

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
          <Grid container spacing={4} direction="column">
            <Grid item>
              <Card>
                <CardHeader title="Tutor Information" />
                <Divider />
                <CardContent>
                  <Grid container direction="column" spacing={2}>
                    <Grid item>
                      <TextField
                        fullWidth
                        label="Tutor Name"
                        value={state.name}
                        variant="outlined"
                        onChange={inp => {
                          setState({ name: inp.target.value });
                        }}
                      />
                    </Grid>
                    <Grid item>
                      <TextField
                        fullWidth
                        label="CIN Number"
                        value={state.cin}
                        variant="outlined"
                        onChange={inp => {
                          setState({ cin: inp.target.value });
                        }}
                      />
                    </Grid>
                  </Grid>
                </CardContent>
              </Card>
            </Grid>
          </Grid>
        </Grid>
        <Grid item xs={4}>
          <Card>
            <CardHeader title="Signature File" />
            <Divider />
            <CardContent>
              {state.signatureURL && (
                <img
                  src={state.signatureURL}
                  className={classes.previewImage}
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
