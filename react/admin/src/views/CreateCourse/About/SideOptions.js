import React from 'react';
import {
  Card,
  CardHeader,
  CardContent,
  TextField,
  InputAdornment,
  Divider
} from '@material-ui/core';

function SideOptions({ state, setState }) {
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
          type="number"
          value={state.hoursToComplete}
          onChange={evt => {
            try {
              setState('hoursToComplete', parseFloat(evt.target.value));
            } catch (err) {}
          }}
        />
      </CardContent>
    </Card>
  );
}

export default SideOptions;
