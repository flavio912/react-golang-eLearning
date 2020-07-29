import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid } from '@material-ui/core';
import IndividualInfo from './IndividualInfo';
import OtherActions from './OtherActions';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Overview({ individual, className, ...rest }) {
  const classes = useStyles();

  if (!individual) {
    return null;
  }

  return (
    <Grid
      {...rest}
      className={clsx(classes.root, className)}
      container
      spacing={3}
    >
      <Grid item lg={8} md={6} xl={3} xs={12}>
        <IndividualInfo individual={individual} />
      </Grid>
      <Grid item lg={4} md={6} xl={3} xs={12}>
        <OtherActions individual={individual} />
      </Grid>
    </Grid>
  );
}

Overview.propTypes = {
  className: PropTypes.string,
  individual: PropTypes.object.isRequired
};

export default Overview;