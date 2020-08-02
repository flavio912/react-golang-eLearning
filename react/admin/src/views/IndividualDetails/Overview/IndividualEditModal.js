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
  TextField,
  Button
} from '@material-ui/core';
import validate from 'validate.js';

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

const schema = {
  firstName: {
    presence: { allowEmpty: false, message: 'is required' },
    length: {
      maximum: 32
    }
  },
  lastName: {
    presence: { allowEmpty: false, message: 'is required' },
    length: {
      maximum: 32
    }
  },
  jobTitle: {
    presence: { allowEmpty: false, message: 'is required' },
    length: {
      maximum: 32
    }
  },
  telephone: {
    presence: { allowEmpty: false, message: 'is required' },
    length: {
      maximum: 13
    }
  },
  email: {
    presence: { allowEmpty: false, message: 'is required' },
    email: true,
    length: {
      maximum: 64
    }
  },
  password: {
    presence: { allowEmpty: false, message: 'is required' },
    length: {
      maximum: 128
    }
  }
};

function IndividualEditModal({ open, onClose, individual, className, ...rest }) {
  const classes = useStyles();
  const [formState, setFormState] = useState({
    isValid: false,
    values: { ...individual },
    touched: {},
    errors: {}
  });

  const hasError = field =>
    !!(formState.touched[field] && formState.errors[field]);

  useEffect(() => {
    const errors = validate(formState.values, schema);

    setFormState(prevFormState => ({
      ...prevFormState,
      isValid: !errors,
      errors: errors || {}
    }));
  }, [formState.values]);

  useEffect(() => {
    if (!individual.isValid && individual.errorMsg) {
      setFormState(prevFormState => ({
        ...prevFormState,
        isValid: false,
        errors: {
          email: [individual.errorMsg]
        },
        touched: {
          ...prevFormState.touched,
          email: true
        }
      }));
      console.warn(individual.errorMsg);
    } else {
      setFormState(prevFormState => ({
        ...prevFormState,
        values: { ...individual }
      }));
    }
  }, [individual]);

  const handleFieldChange = event => {
    event.persist();

    setFormState(prevFormState => ({
      ...prevFormState,
      values: {
        ...prevFormState.values,
        [event.target.name]:
          event.target.type === 'checkbox'
            ? event.target.checked
            : event.target.value
      },
      touched: {
        ...prevFormState.touched,
        [event.target.name]: true
      }
    }));
  };

  if (!open) {
    return null;
  }

  return (
    <Modal onClose={() => onClose(null)} open={open}>
      <Card {...rest} className={clsx(classes.root, className)}>
        <form>
          <CardHeader title="Edit Individual" />
          <Divider />
          <CardContent>
            <Grid container spacing={3}>
              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('firstName')}
                  fullWidth
                  helperText={
                    hasError('firstName') ? formState.errors.firstName[0] : null
                  }
                  label="First Name"
                  name="firstName"
                  onChange={handleFieldChange}
                  value={formState.values.firstName || ''}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('lastName')}
                  fullWidth
                  helperText={
                    hasError('lastName') ? formState.errors.lastName[0] : null
                  }
                  label="Last Name"
                  name="lastName"
                  onChange={handleFieldChange}
                  value={formState.values.lastName || ''}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('jobTitle')}
                  fullWidth
                  helperText={
                    hasError('jobTitle') ? formState.errors.jobTitle[0] : null
                  }
                  label="Job Title"
                  name="jobTitle"
                  onChange={handleFieldChange}
                  value={formState.values.jobTitle || ''}
                  variant="outlined"
                />
              </Grid>

              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('telephone')}
                  fullWidth
                  helperText={
                    hasError('telephone') ? formState.errors.telephone[0] : null
                  }
                  label="Phone Number"
                  name="telephone"
                  onChange={handleFieldChange}
                  value={formState.values.telephone || ''}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('email')}
                  fullWidth
                  helperText={
                    hasError('email') ? formState.errors.email[0] : null
                  }
                  label="Email address"
                  name="email"
                  onChange={handleFieldChange}
                  value={formState.values.email || ''}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('password')}
                  fullWidth
                  helperText={
                    hasError('password') ? formState.errors.password[0] : null
                  }
                  label="Password"
                  name="password"
                  onChange={handleFieldChange}
                  type="password"
                  value={formState.values.password || ''}
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
              onClick={() => onClose(formState.values)}
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

IndividualEditModal.propTypes = {
  className: PropTypes.string,
  individual: PropTypes.any,
  onClose: PropTypes.func,
  open: PropTypes.bool
};

IndividualEditModal.defaultProps = {
  open: false,
  onClose: () => {}
};

export default IndividualEditModal;
