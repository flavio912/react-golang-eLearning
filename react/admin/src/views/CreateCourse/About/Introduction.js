import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  Typography,
  CardContent,
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
      <CardHeader title={'How to complete this Course (500 Words)'} />
      <Divider />
      <CardContent>
        <RichEditor minHeight={400} />
      </CardContent>
    </Card>
  );
}

export default Introduction;
