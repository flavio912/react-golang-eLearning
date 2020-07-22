import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import TagsInput from 'src/components/TagsInput';
import FilesDropzone from 'src/components/FilesDropzone';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  shortDescription: {
    width: '100%'
  }
}));

const UPLOAD_REQUEST = gql`
  mutation UploadRequest($fileType: String!, $contentLength: Int!) {
    moduleBannerImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      successToken
    }
  }
`;

function Overview({ state, setState }) {
  const classes = useStyles();

  const [uploadRequest] = useMutation(UPLOAD_REQUEST);

  const mutationName =
    UPLOAD_REQUEST?.definitions[0].selectionSet?.selections[0].name?.value;
  if (!mutationName) {
    console.error('UploadFile: Could not find mutation name');
    return null;
  }

  const onUpload = async (files) => {
    // Attempt to get upload request
    console.log(files)
    const file = files[0];
    const fType = file.type.replace('image/', '');

    try {
      const resp = await uploadRequest({
        variables: {
          fileType: fType,
          contentLength: file.size
        }
      });

      const data = resp.data[mutationName];
      // Upload to S3
      const uploadResp = await fetch(data.url, {
        method: 'PUT',
        body: file
      });

      if (uploadResp.status !== 200) {
        console.log('Unable to upload');
      }

      setState('bannerImageSuccessToken', data.successToken);
    } catch (err) {
      console.error(err);
    }
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
                            setState('name', inp.target.value);
                        }}
                        placeholder="e.g. Fire Safety Module 1"
                        value={state.name}
                        variant="outlined"
                      />
                    </Grid>
                    <Grid item>
                      <TagsInput
                        allowNew
                        onChange={(tags) => setState('tags', tags)}
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
                      setState('description', inp.target.value);
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
              <FilesDropzone
                onUpload={onUpload}
                limit={1}
              />
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
