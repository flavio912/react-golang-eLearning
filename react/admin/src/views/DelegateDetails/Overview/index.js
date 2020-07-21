import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid } from '@material-ui/core';
import DelegateInfo from './DelegateInfo';
import OtherActions from './OtherActions';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Overview({ delegate, className, ...rest }) {
  const classes = useStyles();

  if (!delegate) {
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
        <DelegateInfo delegate={delegate} />
      </Grid>
      <Grid item lg={4} md={6} xl={3} xs={12}>
        <OtherActions delegate={delegate} />
      </Grid>
    </Grid>
  );
}

Overview.propTypes = {
  className: PropTypes.string,
  delegate: PropTypes.object.isRequired
};

export default Overview;
