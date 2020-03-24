import React from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Typography } from '@material-ui/core';
import Label from 'src/components/Label';

const useStyles = makeStyles(theme => ({
  root: {},
  userName: {
    display: 'flex',
    alignItems: 'center'
  },
  companyTag: {
    marginLeft: theme.spacing(2)
  }
}));

function Header({ className, ...rest }) {
  const classes = useStyles();
  const customer = {
    name: 'John Doe'
  };

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Typography component="h2" gutterBottom variant="overline">
        Users
      </Typography>
      <div className={classes.userName}>
        <Typography component="h1" variant="h3">
          {customer.name}
        </Typography>
        <Label className={classes.companyTag} color={'#20b74e'}>
          FedEx
        </Label>
      </div>
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string
};

export default Header;
