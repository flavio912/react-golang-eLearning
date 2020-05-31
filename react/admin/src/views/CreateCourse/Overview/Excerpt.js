import React from 'react';
import { Card, CardHeader, CardContent, Divider } from '@material-ui/core';
import RichEditor from 'src/components/RichEditor';

function Excerpt() {
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
