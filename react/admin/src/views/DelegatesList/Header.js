import React, { useState } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid, Typography, Button } from '@material-ui/core';
import DelegateCreateModal from './DelegateCreateModal';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Header({ className, onCreate, ...rest }) {
  const classes = useStyles();
  const [openModal, setOpenModal] = useState(false);

  const handleNewDelegate = data => {
    if (data) onCreate(data);
    setOpenModal(false);
  };

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="flex-end" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h2" gutterBottom variant="overline">
            Certificates
          </Typography>
          <Typography component="h1" variant="h3">
            Delegates
          </Typography>
        </Grid>
        <Grid item>
          <Button
            color="primary"
            variant="contained"
            onClick={() => setOpenModal(true)}
          >
            Add delegate
          </Button>
        </Grid>
      </Grid>
      <DelegateCreateModal
        onClose={data => handleNewDelegate(data)}
        open={openModal}
      />
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string,
  onCreate: PropTypes.func
};

export default Header;
