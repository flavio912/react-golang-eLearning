import React, { useState } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid, Typography, Button } from '@material-ui/core';
import AdminCreateModal from './AdminCreateModal';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Header({ className, ...rest }) {
  const classes = useStyles();
  const [openCreate, setOpenCreate] = useState(false);

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="flex-end" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h2" gutterBottom variant="overline">
            Certificates
          </Typography>
          <Typography component="h1" variant="h3">
            Admins
          </Typography>
        </Grid>
        <Grid item>
          <Button
            color="primary"
            variant="contained"
            onClick={() => setOpenCreate(true)}
          >
            Add admin
          </Button>
        </Grid>
      </Grid>
      <AdminCreateModal
        onClose={() => setOpenCreate(false)}
        open={openCreate}
      />
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string
};

export default Header;
