import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  Typography,
  CardContent,
  TextField,
  InputAdornment,
  Divider
} from '@material-ui/core';
import { Editor, EditorState, convertFromRaw } from 'draft-js';
import { makeStyles } from '@material-ui/styles';
import RichEditor from 'src/components/RichEditor';

const useStyles = makeStyles(theme => ({}));

function Introduction() {
  const classes = useStyles();

  const editorState = EditorState.createEmpty();
  return (
    <Card>
      <CardHeader title={'Estimated time to complete'} />
      <Divider />
      <CardContent>
        <TextField
          label="Estimated time"
          InputProps={{
            endAdornment: <InputAdornment position="end">Hours</InputAdornment>
          }}
          variant="outlined"
        />
      </CardContent>
    </Card>
  );
}

export default Introduction;
