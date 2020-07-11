import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid } from '@material-ui/core';
import AdminInfo from './AdminInfo';
import OtherActions from './OtherActions';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Overview({ admin, className, ...rest }) {
  const classes = useStyles();

  if (!admin) {
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
        <AdminInfo admin={admin} />
      </Grid>
      <Grid item lg={4} md={6} xl={3} xs={12}>
        <OtherActions admin={admin} />
      </Grid>
    </Grid>
  );
}

Overview.propTypes = {
  className: PropTypes.string
};

export default Overview;
