import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import {
  Modal,
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Grid,
  Divider,
  Typography,
  TextField,
  Switch,
  Button
} from '@material-ui/core';

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
    justifyContent: 'flex-end'
  }
}));

function AdminEditModal({ open, onClose, admin, className, ...rest }) {
  const classes = useStyles();
  const [values, setValues] = useState(admin);

  useEffect(() => {
    setValues(admin);
  }, [admin]);

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

  if (!open) {
    return null;
  }

  return (
    <Modal onClose={() => onClose(null)} open={open}>
      <Card {...rest} className={clsx(classes.root, className)}>
        <form>
          <CardHeader title="Edit Admin" />
          <Divider />
          <CardContent>
            <Grid container spacing={3}>
              <Grid item md={12} xs={12}>
                <TextField
                  error={!values.isValid && values.errorMsg}
                  fullWidth
                  helperText={
                    !values.isValid && values.errorMsg ? values.errorMsg : null
                  }
                  label="Email address"
                  name="email"
                  onChange={handleFieldChange}
                  value={values.email}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={!values.isValid && values.errorMsg}
                  fullWidth
                  helperText={
                    !values.isValid && values.errorMsg ? values.errorMsg : null
                  }
                  label="First Name"
                  name="firstName"
                  onChange={handleFieldChange}
                  value={values.firstName}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={!values.isValid && values.errorMsg}
                  fullWidth
                  helperText={
                    !values.isValid && values.errorMsg ? values.errorMsg : null
                  }
                  label="Last Name"
                  name="lastName"
                  onChange={handleFieldChange}
                  value={values.lastName}
                  variant="outlined"
                />
              </Grid>
              {/* <Grid item md={6} xs={12}>
                <Typography variant="h5">Email Verified</Typography>
                <Typography variant="body2">
                  Disabling this will automatically send the user a verification
                  email
                </Typography>
                <Switch
                  checked={values.verified}
                  color="secondary"
                  edge="start"
                  name="verified"
                  onChange={handleFieldChange}
                  value={values.verified}
                />
              </Grid> */}
            </Grid>
          </CardContent>
          <Divider />
          <CardActions className={classes.actions}>
            <Button onClick={() => onClose(null)}>Close</Button>
            <Button
              color="primary"
              onClick={() => onClose(values)}
              variant="contained"
            >
              Save
            </Button>
          </CardActions>
        </form>
      </Card>
    </Modal>
  );
}

AdminEditModal.propTypes = {
  className: PropTypes.string,
  admin: PropTypes.any,
  onClose: PropTypes.func,
  open: PropTypes.bool
};

AdminEditModal.defaultProps = {
  open: false,
  onClose: () => {}
};

export default AdminEditModal;
