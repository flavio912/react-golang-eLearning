import React from 'react';
import { TextField, Button, Grid, Chip } from '@material-ui/core';
import { Autocomplete } from '@material-ui/lab';
import Proptypes from 'prop-types';

function TagsInput({ onChange = () => {}, allowNew = false }) {
  // TODO: Get from DB
  const tagOptions = [{ name: 'Aviation Security', color: '#EGEGEG', uuid: '00000000-0000-0000-0000-000000000000' }];

  const onTagAdded = (newValue) => {
    setInputs(newValue);
    onChange(newValue);
  }

  // Allow custom user tags
  const [ newInput, setNewInput ] = React.useState('');
  const [ inputs, setInputs ] = React.useState([]);
  const onNewTag = () => {
    if (allowNew && newInput !== '') {
      onTagAdded([
        ...inputs,
        {
          name: newInput,
          color: '#ff68b4',
          uuid: '00000000-0000-0000-0000-000000000000'
        }
      ]);
      setNewInput('');
    }
  };

  return (
    <Grid container spacing={1} alignItems={'center'}>
      <Grid item xs={11}>
        <Autocomplete
          multiple
          options={tagOptions}
          value={inputs}
          getOptionLabel={option => option.name}
          onChange={(_, newValue) => {
            onTagAdded(newValue);
          }}
          renderTags={(value, getTagProps) =>
            value.map((option, index) => (
              <Chip
                label={option.name}
                style={{
                  backgroundColor: option.color
                }}
                {...getTagProps({ index })}
              />
            ))
          }
          renderInput={params => (
            <TextField
              {...params}
              label="Tags"
              variant="outlined"
              onChange={(inp) => {
                setNewInput(inp.target.value);
              }}
            />
          )}
        />
      </Grid>
      <Grid item xs={1}>
        <Button
          onClick={onNewTag}
        >
          + Add
        </Button>
      </Grid>
    </Grid>
  );
}

TagsInput.propTypes = {
  onChange: Proptypes.func,
  allowNew: Proptypes.bool,
};

export default TagsInput;
