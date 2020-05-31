import React from 'react';
import { Card, CardHeader, CardContent, Divider } from '@material-ui/core';
import RichEditor from 'src/components/RichEditor';

function Introduction() {
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
