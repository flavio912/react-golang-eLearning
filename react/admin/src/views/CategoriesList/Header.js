import React, { useState } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid, Typography, Button } from '@material-ui/core';
import CategorySaveModal from './CategorySaveModal';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Header({ className, onAddCategory, ...rest }) {
  const classes = useStyles();

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="flex-end" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h1" variant="h3">
            Categories
          </Typography>
        </Grid>
        <Grid item>
          <Button color="primary" variant="contained" onClick={onAddCategory}>
            Add category
          </Button>
        </Grid>
      </Grid>
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string,
  onCreateNewIndividual: PropTypes.func
};

export default Header;
