import React, { useState } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import { Grid, Typography, Button } from '@material-ui/core';
import IndividualCreateModal from './IndividualCreateModal';

const useStyles = makeStyles(() => ({
  root: {}
}));

function Header({ className, onCreateNewIndividual, ...rest }) {
  const classes = useStyles();
  const [openCreate, setOpenCreate] = useState(false);

  const handleNewIndividual = data => {
    if (data) onCreateNewIndividual(data);
    setOpenCreate(false);
  };

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="flex-end" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h1" variant="h3">
            Individuals
          </Typography>
        </Grid>
        <Grid item>
          <Button
            color="primary"
            variant="contained"
            onClick={() => setOpenCreate(true)}
          >
            Add individual
          </Button>
        </Grid>
      </Grid>
      <IndividualCreateModal
        onClose={data => handleNewIndividual(data)}
        open={openCreate}
      />
    </div>
  );
}

Header.propTypes = {
  className: PropTypes.string,
  onCreateNewIndividual: PropTypes.func
};

export default Header;
