import React, { useState } from 'react';
import { useHistory } from 'react-router';
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
import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';

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

const DELETE_DELEGATE = gql`
  mutation DeleteDelegate($uuid: UUID!) {
    deleteDelegate(input: { uuid: $uuid })
  }
`;

function OtherActions({ delegate, className, ...rest }) {
  const classes = useStyles();
  const history = useHistory();
  const [openDialog, setOpenDialog] = useState(false);
  const [deleteDelegate] = useMutation(DELETE_DELEGATE);

  const handleDeleteAccount = async event => {
    try {
      await deleteDelegate({
        variables: {
          uuid: delegate.uuid
        }
      });
      setOpenDialog(false);
      history.push('/delegates');
    } catch (err) {
      console.warn(err);
    }
  };

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
          Delete Delegate Account
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
          {'Delete Delegate account?'}
        </DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Are you sure to delete this delegate account?
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
            Delete
          </Button>
        </DialogActions>
      </Dialog>
    </Card>
  );
}

OtherActions.propTypes = {
  className: PropTypes.string,
  delegate: PropTypes.object.isRequired
};

export default OtherActions;
