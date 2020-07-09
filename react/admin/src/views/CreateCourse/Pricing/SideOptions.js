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

function SideOptions({ state, setState }) {
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
              value={state.priceType}
              onChange={(evt, value) => {
                if (value === 'free') {
                  setState('price', parseFloat(0));
                }
                setState('priceType', value);
              }}
            >
              <FormControlLabel
                value="paid"
                control={<Radio />}
                label={
                  <TextField
                    label="Price"
                    InputProps={{
                      startAdornment: (
                        <InputAdornment position="start">£</InputAdornment>
                      )
                    }}
                    type="number"
                    value={state.price}
                    onChange={evt => {
                      try {
                        setState('price', parseFloat(evt.target.value));
                      } catch (err) {
                        return;
                      }
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
