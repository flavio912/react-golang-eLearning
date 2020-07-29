import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid } from '@material-ui/core';
import CompanyInfo from './CompanyInfo';
import SendEmails from './SendEmails';
import OtherActions from './OtherActions';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Summary({ className, company, onUpdate, ...rest }) {
  const classes = useStyles();

  return (
    <Grid
      {...rest}
      className={clsx(classes.root, className)}
      container
      spacing={3}
    >
      <Grid item lg={4} md={6} xl={3} xs={12}>
        <CompanyInfo company={company} onUpdate={onUpdate} />
      </Grid>
      {/* {company.approved && (
        <Grid item lg={4} md={6} xl={3} xs={12}>
          <SendEmails company={company} />
        </Grid>
      )} */}
      <Grid item lg={4} md={6} xl={3} xs={12}>
        <OtherActions />
      </Grid>
    </Grid>
  );
}

Summary.propTypes = {
  className: PropTypes.string
};

export default Summary;
