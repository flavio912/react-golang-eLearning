import React, { useState } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import {
  Card,
  CardHeader,
  CardContent,
  Button,
  Divider,
  Typography,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle
} from '@material-ui/core';
import DeleteIcon from '@material-ui/icons/DeleteOutline';

const useStyles = makeStyles(theme => ({
  root: {},
  mainActions: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start'
  },
  notice: {
    marginTop: theme.spacing(1)
  },
  deleteButton: {
    marginTop: theme.spacing(1),
    color: theme.palette.common.white,
    backgroundColor: theme.palette.error.main,
    '&:hover': {
      backgroundColor: theme.palette.error.dark
    }
  },
  buttonIcon: {
    marginRight: theme.spacing(1)
  }
}));

const handleDeleteAccount = () => {};

function OtherActions({ className, ...rest }) {
  const classes = useStyles();
  const [openDialog, setOpenDialog] = useState(false);

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Other actions" />
      <Divider />
      <CardContent>
        <div className={classes.mainActions}></div>
        <Typography className={classes.notice} variant="body2">
          Note: Once deleted data cannot be retrieved
        </Typography>
        <Button
          className={classes.deleteButton}
          onClick={() => {
            setOpenDialog(true);
          }}
        >
          <DeleteIcon className={classes.buttonIcon} />
          Delete Admin Account
        </Button>
      </CardContent>
      <Dialog
        open={openDialog}
        onClose={() => {
          setOpenDialog(false);
        }}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">
          {'Delete Admin account?'}
        </DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Are you sure to delete this admin account?
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => {
              setOpenDialog(false);
            }}
          >
            cancel
          </Button>
          <Button onClick={handleDeleteAccount} color="primary" autoFocus>
            Agree
          </Button>
        </DialogActions>
      </Dialog>
    </Card>
  );
}

OtherActions.propTypes = {
  className: PropTypes.string
};

export default OtherActions;
