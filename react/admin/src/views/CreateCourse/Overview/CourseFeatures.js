import React from 'react';
import {
  Card,
  CardHeader,
  Typography,
  CardContent,
  Grid,
  Radio,
  RadioGroup,
  FormControlLabel,
  Divider,
  Switch
} from '@material-ui/core';

function CourseFeatures({ state, setState }) {
  return (
    <Card>
      <CardHeader title={'Course Features'} />
      <CardContent>
        <Grid container direction="column" spacing={2}>
          <Grid item>
            <Typography variant={'overline'}>COURSE TYPE</Typography>
            <RadioGroup
              aria-label="coursetype"
              name="course type"
              value={state.courseType}
              onChange={(evt, value) => {
                setState({ courseType: value });
              }}
            >
              <FormControlLabel
                value="online"
                control={<Radio />}
                label="Online Course"
              />
              <FormControlLabel
                value="classroom"
                control={<Radio />}
                label="Classroom course"
              />
            </RadioGroup>
          </Grid>
          <Grid item>
            <Typography variant={'overline'}>ACCESS TYPE</Typography>
            <RadioGroup
              aria-label="access type"
              name="accessType"
              value={state.accessType}
              onChange={(evt, value) => {
                setState({ accessType: value });
              }}
            >
              <FormControlLabel
                value="restricted"
                control={<Radio />}
                label="Restricted"
              />
              <FormControlLabel
                value="open"
                control={<Radio />}
                label="Open Access"
              />
            </RadioGroup>
          </Grid>
          <Grid item>
            <Divider />
          </Grid>
          <Grid item>
            <Grid container direction={'column'} spacing={1}>
              <Grid item>
                <Typography variant={'overline'}>BACKGROUND CHECK</Typography>
              </Grid>
              <Grid item>
                <Typography variant={'h5'}>Background Check</Typography>
              </Grid>
              <Grid item>
                <Typography variant={'body2'}>
                  Do managers need to provide a 5 year background check of their
                  delegates to take this course
                </Typography>
              </Grid>
              <Grid item>
                <Switch
                  checked={state.backgroundCheck}
                  onChange={(evt, checked) => {
                    setState({ backgroundCheck: checked });
                  }}
                  color="primary"
                  name="checkedB"
                  inputProps={{ 'aria-label': 'primary checkbox' }}
                />
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default CourseFeatures;
