import React, { useState } from 'react';
import {
  Grid,
  TextField,
  Card,
  CardHeader,
  CardContent,
  Divider,
  FormControl,
  MenuItem,
  InputLabel,
  Select,
  Typography,
  Button,
  CircularProgress
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';

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
  input: {
    display: 'none'
  },
  previewImage: {
    width: 200,
    maxHeight: 200,
    marginLeft: 16
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

function AnswerInput({ answer, onSave, onChange }) {
  const classes = useStyles();

  const [uploadRequest] = useMutation(UPLOAD_REQUEST);
  const [uploadText, setUploadText] = useState('Upload Image');

  const uploadChange = async evt => {
    // Attempt to get upload request
    console.log(evt.target.files);
    const file = evt.target.files[0];
    const fType = file.type.replace('image/', '');

    setUploadText(<CircularProgress />);
    try {
      const resp = await uploadRequest({
        variables: {
          fileType: fType,
          contentLength: file.size
        }
      });

      const data = resp.data.answerImageUploadRequest;
      // Upload to S3
      const uploadResp = await fetch(data.url, {
        method: 'PUT',
        body: file
      });

      if (uploadResp.status !== 200) {
        console.log('Unable to upload');
      }

      setUploadText(file.name);
      const newAns = {
        ...answer,
        imageToken: data.successToken,
        imageURL: URL.createObjectURL(file),
        isCorrect: false
      };
      onChange(newAns);
    } catch (err) {
      setUploadText('Unable to upload, please try again');
    }
  };

  return (
    <Card>
      <CardHeader title={'Answers'} />
      <Divider />
      <CardContent>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Typography variant="h6">Answer Type</Typography>
          </Grid>
          <Grid item>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel id="demo-simple-select-outlined-label">
                Answer type
              </InputLabel>
              <Select
                labelId="demo-simple-select-outlined-label"
                id="demo-simple-select-outlined"
                value={answer.answerType || 'TEXT'}
                onChange={({ target }) => {
                  const value = target.value;

                  var newAns = { ...answer };
                  switch (value) {
                    case 'TEXT':
                      newAns.imageToken = undefined;
                      newAns.imageURL = undefined;
                      break;
                    case 'IMAGE':
                      newAns.text = undefined;
                      break;
                    default:
                      console.error('Unable to find onChange type');
                      break;
                  }

                  newAns.answerType = value;
                  onChange(newAns);
                }}
                label="Answer Type"
              >
                <MenuItem value={'TEXT'}>Text</MenuItem>
                <MenuItem value={'IMAGE'}>Image</MenuItem>
                <MenuItem value={'TEXT_IMAGE'}>Text + Image</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          {(answer.answerType === 'TEXT' ||
            answer.answerType === 'TEXT_IMAGE') && (
            <>
              <Grid item>
                <Typography variant="h6">Answer Text</Typography>
              </Grid>
              <Grid item>
                <TextField
                  fullWidth
                  label="Answer title"
                  name="answer"
                  onChange={inp => {
                    var newAns = { ...answer };
                    newAns.text = inp.target.value;
                    onChange(newAns);
                  }}
                  placeholder="Answer text"
                  value={answer.text}
                  variant="outlined"
                />
              </Grid>
            </>
          )}
          {(answer.answerType === 'IMAGE' ||
            answer.answerType === 'TEXT_IMAGE') && (
            <>
              <Grid item>
                <Typography variant="h6">Answer Image</Typography>
              </Grid>
              {answer.imageURL && (
                <img
                  src={answer.imageURL}
                  className={classes.previewImage}
                  alt="preview"
                />
              )}
              <Grid item>
                <input
                  accept="image/*"
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
              </Grid>
            </>
          )}
          <Grid item>
            <Button variant="contained" color="primary" onClick={onSave}>
              Save Answer
            </Button>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default AnswerInput;
