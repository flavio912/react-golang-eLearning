import React from 'react';
import { Card, CardHeader, CardContent, Divider } from '@material-ui/core';
import RichEditor from 'src/components/RichEditor';

function EditorCard({ title, ...props }) {
  return (
    <Card>
      <CardHeader title={title} />
      <Divider />
      <CardContent>
        <RichEditor {...props} />
      </CardContent>
    </Card>
  );
}

export default EditorCard;
