import React from 'react';
import { Card, CardHeader, CardContent } from '@material-ui/core';
import { gql } from 'apollo-boost';
import { makeStyles } from '@material-ui/styles';

import UploadFile from 'src/components/UploadFile';

const UPLOAD_REQUEST = gql`
  mutation CourseBannerUploadRequest($fileType: String!, $contentLength: Int!) {
    courseBannerImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      url
      successToken
    }
  }
`;

const useStyles = makeStyles(theme => ({
  bannerImagePreview: {
    width: '100%'
  }
}));

function CourseBannerUpload({ state, setState }) {
  const classes = useStyles();

  return (
    <Card>
      <CardHeader title={'Course Banner Image'} />
      <CardContent>
        {state.bannerImageURL && (
          <img
            src={state.bannerImageURL}
            className={classes.bannerImagePreview}
            alt="Banner preview"
          />
        )}
        <UploadFile
          uploadMutation={UPLOAD_REQUEST}
          onUploaded={(token, url) => {
            setState({ bannerImageURL: url, bannerImageSuccess: token });
          }}
        />
      </CardContent>
    </Card>
  );
}

export default CourseBannerUpload;
