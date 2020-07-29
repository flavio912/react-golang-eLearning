import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Typography, Grid, Button } from '@material-ui/core';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Header({ className, companyName, approved, onApprove, ...rest }) {
  const classes = useStyles();

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="center" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h2" gutterBottom variant="overline">
            Companies
          </Typography>
          <Typography component="h1" variant="h3">
            {companyName}
          </Typography>
        </Grid>
        <Grid item>
          {!approved && (
            <Button
              color="primary"
              size="small"
              onClick={onApprove}
              variant="outlined"
            >
              Approve
            </Button>
          )}
        </Grid>
      </Grid>
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string
};

export default Header;
