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
  Button,
  FormControl,
  FormControlLabel,
  Checkbox,
  InputLabel,
  Select,
  MenuItem
} from '@material-ui/core';
import validate from 'validate.js';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import UploadFile from 'src/components/UploadFile';

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
  }
};

const GET_COMPANIES = gql`
  query GetCompanies {
    companies {
      edges {
        uuid
        name
        managers {
          edges {
            email
          }
          pageInfo {
            total
          }
        }
        delegates {
          edges {
            email
          }
          pageInfo {
            total
          }
        }
      }
    }
  }
`;

const UPLOAD_REQUEST = gql`
  mutation UploadRequest($fileType: String!, $contentLength: Int!) {
    profileImageUploadRequest(
      input: { fileType: $fileType, contentLength: $contentLength }
    ) {
      url
      successToken
    }
  }
`;

function DelegateEditModal({ open, onClose, delegate, className, ...rest }) {
  const classes = useStyles();
  const [formState, setFormState] = useState({
    isValid: false,
    values: { ...delegate },
    touched: {},
    errors: {}
  });
  const [companies, setCompanies] = useState([]);
  const { loading, error, data } = useQuery(GET_COMPANIES);

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
    if (!delegate.isValid && delegate.errorMsg) {
      setFormState(prevFormState => ({
        ...prevFormState,
        isValid: false,
        errors: {
          email: [delegate.errorMsg]
        },
        touched: {
          ...prevFormState.touched,
          email: true
        }
      }));
      console.warn(delegate.errorMsg);
    } else {
      setFormState(prevFormState => ({
        ...prevFormState,
        values: { ...delegate }
      }));
    }
  }, [delegate]);

  useEffect(() => {
    if (loading || error || !data) return;

    setCompanies(data?.companies?.edges);
  }, [loading, error, data]);

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
          <CardHeader title="Edit Delegate" />
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
                <FormControl
                  variant="outlined"
                  className={classes.formControl}
                  style={{ width: '100%' }}
                >
                  <InputLabel id="demo-simple-select-outlined-label">
                    Company
                  </InputLabel>
                  <Select
                    labelId="demo-simple-select-outlined-label"
                    id="demo-simple-select-outlined"
                    onChange={handleFieldChange}
                    value={formState.values.company.uuid || ''}
                    name="company"
                    label="Company"
                  >
                    {companies.length > 0 &&
                      companies.map((company, index) => {
                        return (
                          <MenuItem value={company.uuid} key={index}>
                            {company.name}
                          </MenuItem>
                        );
                      })}
                  </Select>
                </FormControl>
              </Grid>
              <Grid item md={6} xs={12}>
                <TextField
                  error={hasError('password')}
                  fullWidth
                  helperText={
                    hasError('password') ? formState.errors.password[0] : null
                  }
                  label="New Password"
                  name="password"
                  onChange={handleFieldChange}
                  type="password"
                  value={formState.values.password || ''}
                  variant="outlined"
                />
              </Grid>
              <Grid item md={6} xs={12}>
                <UploadFile
                  uploadMutation={UPLOAD_REQUEST}
                  onUploaded={(token, url) => {
                    setFormState(prevFormState => ({
                      ...prevFormState,
                      values: {
                        ...prevFormState.values,
                        profileUploadToken: token
                      }
                    }));
                  }}
                />
              </Grid>
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

DelegateEditModal.propTypes = {
  className: PropTypes.string,
  delegate: PropTypes.any,
  onClose: PropTypes.func,
  open: PropTypes.bool
};

DelegateEditModal.defaultProps = {
  open: false,
  onClose: () => {}
};

export default DelegateEditModal;
