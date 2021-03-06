import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid, Typography, Button } from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  root: {},
  publish: {
    marginLeft: theme.spacing(2)
  }
}));

function Header({ className, onSave, title, ...rest }) {
  const classes = useStyles();

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="flex-end" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h2" gutterBottom variant="overline">
            Courses
          </Typography>
          <Typography component="h1" variant="h3">
            {title}
          </Typography>
        </Grid>
        <Grid item>
          <Button color="primary" variant="contained" onClick={onSave}>
            Save Test
          </Button>
        </Grid>
      </Grid>
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string
};

export default Header;
