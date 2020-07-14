import React from 'react';
import { TextField, Button, Grid, Chip } from '@material-ui/core';
import { Autocomplete } from '@material-ui/lab';
import Proptypes from 'prop-types';

function TagsInput({ onChange = () => {} }) {
  // TODO: Get from DB
  const tagOptions = [{ title: 'Aviation Security', value: 'avsec' }];

  return (
    <Grid container spacing={1} alignItems={'center'}>
      <Grid item xs={11}>
        <Autocomplete
          multiple
          options={tagOptions}
          getOptionLabel={option => option.title}
          onChange={(event, newValue) => {
            onChange(newValue);
          }}
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
  );
}

TagsInput.propTypes = {
  onChange: Proptypes.func
};

export default TagsInput;
