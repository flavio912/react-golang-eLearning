import React from 'react';
import {
  Card,
  CardHeader,
  CardContent,
  Divider,
  Radio,
  RadioGroup,
  Typography,
  FormControlLabel,
  TextField,
  InputAdornment,
  Grid
} from '@material-ui/core';

function SideOptions() {
  return (
    <Card>
      <CardHeader title={'Course Pricing'} />
      <Divider />
      <CardContent>
        <Grid container direction={'column'} spacing={2}>
          <Grid item>
            <Typography variant={'overline'}>SET PRICING</Typography>
          </Grid>
          <Grid item>
            <RadioGroup
              aria-label="gender"
              name="gender1"
              value={'online'}
              onChange={() => {}}
            >
              <FormControlLabel
                value="online"
                control={<Radio />}
                label={
                  <TextField
                    label="Price"
                    InputProps={{
                      startAdornment: (
                        <InputAdornment position="start">£</InputAdornment>
                      )
                    }}
                    variant="outlined"
                  />
                }
              />
              <FormControlLabel
                value="free"
                control={<Radio />}
                label="FREE (CLASSROOM COURSE)"
              />
            </RadioGroup>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default SideOptions;
