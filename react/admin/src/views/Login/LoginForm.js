/* eslint-disable no-unused-vars */
import React, { useState, useEffect } from 'react';
import { useHistory } from 'react-router';
import validate from 'validate.js';
import clsx from 'clsx';
import PropTypes from 'prop-types';
import { useDispatch } from 'react-redux';
import { makeStyles } from '@material-ui/styles';
import { Button, TextField, Typography } from '@material-ui/core';

import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';

const ADMIN_LOGIN = gql`
  mutation AdminLogin($email: String!, $pass: String!) {
    adminLogin(input: { email: $email, password: $pass }) {
      token
    }
  }
`;

const schema = {
  email: {
    presence: { allowEmpty: false, message: 'is required' },
    email: true
  },
  password: {
    presence: { allowEmpty: false, message: 'is required' }
  }
};

const useStyles = makeStyles(theme => ({
  root: {},
  fields: {
    margin: theme.spacing(-1),
    display: 'flex',
    flexWrap: 'wrap',
    '& > *': {
      flexGrow: 1,
      margin: theme.spacing(1)
    }
  },
  submitButton: {
    marginTop: theme.spacing(2),
    width: '100%'
  }
}));

function LoginForm({ className, ...rest }) {
  const classes = useStyles();
  const history = useHistory();
  const [formState, setFormState] = useState({
    isValid: false,
    values: {},
    touched: {},
    errors: {}
  });
  const [
    adminLogin,
    { loading: isAuthenticating, error: authError }
  ] = useMutation(ADMIN_LOGIN);

  const handleChange = event => {
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

  const handleSubmit = async event => {
    event.preventDefault();
    try {
      const resp = await adminLogin({
        variables: {
          email: formState.values.email,
          pass: formState.values.password
        }
      });
      localStorage.setItem('auth', resp.data?.adminLogin?.token);
      console.log(resp);
      history.push('/');
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

  return (
    <form
      {...rest}
      className={clsx(classes.root, className)}
      onSubmit={handleSubmit}
    >
      <div className={classes.fields}>
        <TextField
          error={hasError('email')}
          fullWidth
          helperText={hasError('email') ? formState.errors.email[0] : null}
          label="Email address"
          name="email"
          onChange={handleChange}
          value={formState.values.email || ''}
          variant="outlined"
        />
        <TextField
          error={hasError('password')}
          fullWidth
          helperText={
            hasError('password') ? formState.errors.password[0] : null
          }
          label="Password"
          name="password"
          onChange={handleChange}
          type="password"
          value={formState.values.password || ''}
          variant="outlined"
        />
      </div>
      <Button
        className={classes.submitButton}
        color="secondary"
        disabled={!formState.isValid}
        size="large"
        type="submit"
        variant="contained"
      >
        Sign in
      </Button>
    </form>
  );
}

LoginForm.propTypes = {
  className: PropTypes.string
};

export default LoginForm;
