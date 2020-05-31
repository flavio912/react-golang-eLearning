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

function Excerpt() {
  const classes = useStyles();

  const editorState = EditorState.createEmpty();
  return (
    <Card>
      <CardHeader title={'About this course (70 Words)'} />
      <Divider />
      <CardContent>
        <RichEditor inlineOnly />
      </CardContent>
    </Card>
  );
}

export default Excerpt;
