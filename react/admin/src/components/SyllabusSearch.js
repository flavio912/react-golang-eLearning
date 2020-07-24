import React from 'react';
import {
  Grid,
  Typography,
  TextField,
  Chip
} from '@material-ui/core';
import Autocomplete from '@material-ui/lab/Autocomplete';
import SearchIcon from '@material-ui/icons/Search';

function SyllabusSearch({ placeholder, searchFilters, setSearchFilters, searchResults, setSearchText, onChange }) {
  return (
    <Autocomplete
      multiple
      options={[...searchFilters.filters, ...searchResults]}
      getOptionLabel={option => option.name}
      renderOption={(option) => (
        <Typography variant="body1" color="textPrimary">
          {option.name} - <strong>{option.type.replace(/^\w/, (c) => c.toUpperCase())}</strong>
        </Typography>
      )}
      onChange={(_, values) => {
        values &&
          values.map((value, index) => {
            if (value.isFilter) {
              const newFilters = {...searchFilters};
              newFilters[value.isFilter] = !searchFilters[value.isFilter];
              setSearchFilters(newFilters);
            } else {
              onChange(value);
              values.splice(index, 1);
            }
          })                 
      }}
      renderTags={(values, getTagProps) =>
      values.map((option, index) => (
        <Chip
          label={option.name}
          style={{
            backgroundColor: option.color
          }}
          {...getTagProps({ index })}
          onDelete={() => {
            if (option.isFilter) {
              const newFilters = {...searchFilters};
              newFilters[option.isFilter] = !searchFilters[option.isFilter];
              setSearchFilters(newFilters);
              values.splice(index, 1);
            }
          }}
        />
      ))
      }
      renderInput={params => (
      <Grid container spacing={1} justify="center">
        <Grid item>
          <SearchIcon style={{marginTop: 5}}/>
        </Grid>
        <Grid item xs={11}>
          <TextField
              {...params}
              placeholder={placeholder}
              onChange={inp => {
              setSearchText(inp.target.value);
            }}
          />
          </Grid>
      </Grid>
      )}
    />
  )
}

export default SyllabusSearch;