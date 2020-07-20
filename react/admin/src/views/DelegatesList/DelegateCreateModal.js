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
import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';
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

const CREATE_DELEGATE = gql`
  mutation CreateDelegate(
    $firstName: String!
    $lastName: String!
    $jobTitle: String!
    $telephone: String
    $email: String
    $profileImageUploadToken: String
    $companyUUID: UUID
    $generatePassword: Boolean
  ) {
    createDelegate(
      input: {
        firstName: $firstName
        lastName: $lastName
        jobTitle: $jobTitle
        telephone: $telephone
        email: $email
        profileImageUploadToken: $profileImageUploadToken
        companyUUID: $companyUUID
        generatePassword: $generatePassword
      }
    ) {
      uuid
      firstName
      lastName
      jobTitle
      telephone
      email
      profileImageUploadToken
      companyUUID
    }
  }
`;

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
  }
};

function DelegateCreateModal({ open, onClose, className, ...rest }) {
  const classes = useStyles();
  const [formState, setFormState] = useState({
    isValid: false,
    values: {},
    touched: {},
    errors: {}
  });
  const [createDelegate] = useMutation(CREATE_DELEGATE);

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

  useEffect(() => {
    if (!open) return;

    setFormState({
      isValid: false,
      values: {},
      touched: {},
      errors: {}
    });
  }, [open]);

  const handleCreateDelegate = async event => {
    event.preventDefault();
    try {
      // const resp = await createAdmin({
      //   variables: {
      //     firstName: formState.values.firstName,
      //     lastName: formState.values.lastName,
      //     email: formState.values.email,
      //     password: formState.values.password
      //   }
      // });
      // onClose(resp.data);
    } catch (err) {
      setFormState(prevFormState => ({
        ...prevFormState,
        isValid: false,
        errors: {
          email: [err?.graphQLErrors[0]?.message]
        }
      }));
      console.warn(err);
    }
  };

  return (
    <Modal onClose={onClose} open={open}>
      <Card {...rest} className={clsx(classes.root, className)}>
        <form>
          <CardHeader title="Create Delegate" />
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
            </Grid>
          </CardContent>
          <Divider />
          <CardActions className={classes.actions}>
            <Button onClick={onClose}>Close</Button>
            <Button
              color="primary"
              onClick={handleCreateDelegate}
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

DelegateCreateModal.propTypes = {
  className: PropTypes.string,
  onClose: PropTypes.func,
  open: PropTypes.bool
};

DelegateCreateModal.defaultProps = {
  open: false,
  onClose: () => {}
};

export default DelegateCreateModal;
