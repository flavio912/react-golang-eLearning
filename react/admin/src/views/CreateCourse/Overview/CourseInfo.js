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

function CourseInfo() {
  const [title, setTitle] = useState();

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
                setTitle(inp.target.value);
              }}
              placeholder="e.g Dangerous Goods"
              value={title}
              variant="outlined"
            />
          </Grid>
          <Grid item>
            <Grid container spacing={1} alignItems={'center'}>
              <Grid item xs={11}>
                <Autocomplete
                  options={categoryOptions}
                  getOptionLabel={option => option.title}
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
            <Grid container spacing={1} alignItems={'center'}>
              <Grid item xs={11}>
                <Autocomplete
                  multiple
                  options={categoryOptions}
                  getOptionLabel={option => option.title}
                  renderTags={(value, getTagProps) =>
                    value.map((option, index) => (
                      <Chip
                        variant="outlined"
                        label={option.title}
                        {...getTagProps({ index })}
                      />
                    ))
                  }
                  renderInput={params => (
                    <TextField {...params} label="Tags" variant="outlined" />
                  )}
                />
              </Grid>
              <Grid item xs={1}>
                <Button>+ Add</Button>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default CourseInfo;
