import React from 'react';
import {
  Card,
  CardHeader,
  CardContent,
  TextField,
  InputAdornment,
  Divider
} from '@material-ui/core';

function Introduction() {
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
