import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  TextField,
  CardContent,
  Button,
  Grid,
  Chip,
  Divider
} from '@material-ui/core';
import { Autocomplete } from '@material-ui/lab';
import TagsInput from 'src/components/TagsInput';

function CourseInfo({ state, setState }) {
  const categoryOptions = [{ title: 'Aviation Security', value: 'avsec' }];

  return (
    <Card>
      <CardHeader title={'Course Info'} />
      <Divider />
      <CardContent>
        <Grid container direction="column" spacing={2}>
          <Grid item>
            <TextField
              fullWidth
              label="Course Name"
              name="courseName"
              onChange={inp => {
                setState('name', inp.target.value);
              }}
              placeholder="e.g Dangerous Goods"
              value={state.name}
              variant="outlined"
            />
          </Grid>
          <Grid item>
            <Grid container spacing={1} alignItems={'center'}>
              <Grid item xs={11}>
                <Autocomplete
                  options={categoryOptions}
                  getOptionLabel={option => option.title}
                  onChange={(event, newValue) => {
                    setState('primaryCategory', newValue.value);
                  }}
                  renderInput={params => (
                    <TextField
                      {...params}
                      label="Primary Course Category"
                      variant="outlined"
                    />
                  )}
                />
              </Grid>
              <Grid item xs={1}>
                <Button>+ Add</Button>
              </Grid>
            </Grid>
          </Grid>
          <Grid item>
            <Grid container spacing={1} alignItems={'center'}>
              <Grid item xs={11}>
                <Autocomplete
                  options={categoryOptions}
                  getOptionLabel={option => option.title}
                  onChange={(event, newValue) => {
                    setState('secondaryCategory', newValue.value);
                  }}
                  renderInput={params => (
                    <TextField
                      {...params}
                      label="Secondary Course Category"
                      variant="outlined"
                    />
                  )}
                />
              </Grid>
              <Grid item xs={1}>
                <Button>+ Add</Button>
              </Grid>
            </Grid>
          </Grid>
          <Grid item>
            <TagsInput onChange={newVal => setState('tags', newVal)} />
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default CourseInfo;
