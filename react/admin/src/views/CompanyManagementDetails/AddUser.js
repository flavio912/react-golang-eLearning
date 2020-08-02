import React, { useState } from 'react';
import { makeStyles } from '@material-ui/styles';
import {
  Button,
  Modal,
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Grid,
  Divider,
  TextField,
  Typography
} from '@material-ui/core';
import clsx from 'clsx';

const useStyles = makeStyles(theme => ({
  root: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    outline: 'none',
    boxShadow: theme.shadows[20],
    width: 700,
    maxHeight: '100%',
    overflowY: 'auto',
    maxWidth: '100%'
  },
  actions: {
    display: 'flex',
    justifyContent: 'flex-end'
  },
  spaceBottom: {
    marginBottom: theme.spacing(2)
  },
  spaceTop: {
    marginTop: theme.spacing(2)
  }
}));

export default function AddUser({
  onClose,
  open,
  className,
  company,
  userType,
  ...rest
}) {
  const classes = useStyles();
  const [values, setValues] = useState({});

  const handleFieldChange = event => {
    event.persist();
    setValues(currentValues => ({
      ...currentValues,
      [event.target.name]:
        event.target.type === 'checkbox'
          ? event.target.checked
          : event.target.value
    }));
  };

  return (
    <Modal onClose={onClose} open={open}>
      <Card {...rest} className={clsx(classes.root, className)}>
        <CardHeader title="Add User" />
        <Divider />
        <CardContent>
          <Grid container spacing={3} className={classes.spaceBottom}>
            <Grid item md={6} xs={12}>
              <Typography variant="overline">COMPANY NAME</Typography>
              <Typography variant="h4">{company.name}</Typography>
            </Grid>
            <Grid item md={6} xs={12}>
              <Typography variant="overline">COMPANY CONTACT</Typography>
              <Typography variant="body1">{company.email}</Typography>
            </Grid>
          </Grid>
          <Grid container spacing={3}>
            <Grid item md={6} xs={12}>
              <TextField
                fullWidth
                label="First Name"
                name="firstName"
                onChange={handleFieldChange}
                value={values.firstName}
                variant="outlined"
              />
            </Grid>
            <Grid item md={6} xs={12}>
              <TextField
                fullWidth
                label="Last Name"
                name="lastName"
                onChange={handleFieldChange}
                value={values.lastName}
                variant="outlined"
              />
            </Grid>
            <Grid item md={6} xs={12}>
              <TextField
                fullWidth
                label="Email Address"
                name="email"
                onChange={handleFieldChange}
                value={values.email}
                variant="outlined"
              />
            </Grid>
            <Grid item></Grid>
          </Grid>
        </CardContent>
        <Divider />
        <CardActions className={classes.actions}>
          <Button onClick={onClose}>Close</Button>
          <Button
            color="primary"
            onClick={() => onClose(values)}
            variant="contained"
          >
            Save
          </Button>
        </CardActions>
      </Card>
    </Modal>
  );
}
