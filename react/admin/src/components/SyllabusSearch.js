import React from 'react';
import { Grid, Typography, TextField, Chip } from '@material-ui/core';
import Autocomplete from '@material-ui/lab/Autocomplete';
import SearchIcon from '@material-ui/icons/Search';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  chip: {},
  result: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    width: '100%'
  }
}));

function SyllabusSearch({
  placeholder,
  searchFilters,
  setSearchFilters,
  searchResults,
  setSearchText,
  onChange
}) {
  const classes = useStyles();
  return (
    <Autocomplete
      multiple
      options={[...searchFilters.filters, ...searchResults]}
      getOptionLabel={option => option.name}
      renderOption={option => (
        <div className={classes.result}>
          <div>
            <Typography variant="body1" color="textPrimary">
              {option.name}
            </Typography>
          </div>
          <Chip
            label={option.type.replace(/^\w/, c => c.toUpperCase())}
            className={classes.chip}
          />
        </div>
      )}
      onChange={(_, values) => {
        values &&
          values.forEach((value, index) => {
            if (value.isFilter) {
              const newFilters = { ...searchFilters };
              newFilters[value.isFilter] = !searchFilters[value.isFilter];
              setSearchFilters(newFilters);
            } else {
              onChange(value);
              values.splice(index, 1);
            }
          });
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
                const newFilters = { ...searchFilters };
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
            <SearchIcon style={{ marginTop: 5 }} />
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
  );
}

export default SyllabusSearch;
