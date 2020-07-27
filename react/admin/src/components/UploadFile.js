import React, { useState } from 'react';
import { Button, CircularProgress } from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { useMutation } from '@apollo/react-hooks';

const useStyles = makeStyles(theme => ({
  input: {
    display: 'none'
  }
}));

export default function UploadFile({
  title = 'Upload Image',
  uploadMutation,
  onUploaded,
  acceptType = 'image/*'
}) {
  const classes = useStyles();

  const [uploadRequest] = useMutation(uploadMutation);
  const [uploadText, setUploadText] = useState(title);

  const mutationName =
    uploadMutation?.definitions[0].selectionSet?.selections[0].name?.value;
  if (!mutationName) {
    console.error('UploadFile: Could not find mutation name');
    return null;
  }

  const uploadChange = async evt => {
    // Attempt to get upload request
    const file = evt.target.files[0];
    const split = file.name.split('.');
    const fType = split[split.length - 1];
    setUploadText(<CircularProgress />);
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
        setUploadText('Unable to upload, try again');
        return;
      }

      setUploadText(file.name);

      onUploaded(data.successToken, URL.createObjectURL(file));
    } catch (err) {
      setUploadText('Unable to upload, please try again');
    }
  };

  return (
    <>
      <input
        accept={acceptType}
        className={classes.input}
        id="contained-button-file"
        onChange={uploadChange}
        type="file"
      />
      <label htmlFor="contained-button-file">
        <Button variant="contained" component="span">
          {uploadText}
        </Button>
      </label>
    </>
  );
}
