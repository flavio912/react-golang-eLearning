import React, { useState } from 'react';
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
import { gql } from 'apollo-boost';
import { useMutation } from '@apollo/react-hooks';
import ErrorModal from 'src/components/ErrorModal';

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

const UPDATE_COMPANY = gql`
  mutation UpdateCompany(
    $uuid: UUID!
    $name: String
    $contactEmail: String
    $addressLine1: String
    $addressLine2: String
    $county: String
    $postCode: String
    $country: String
    $isContract: Boolean
  ) {
    updateCompany(
      input: {
        uuid: $uuid
        companyName: $name
        contactEmail: $contactEmail
        addressLine1: $addressLine1
        addressLine2: $addressLine2
        county: $county
        postCode: $postCode
        country: $country
        isContract: $isContract
      }
    ) {
      uuid
    }
  }
`;

function CompanyEditModal({ open, onClose, company, className, ...rest }) {
  const classes = useStyles();
  const [values, setValues] = useState({
    ...company,
    ...company.address
  });

  const [updateCompany, { error: mutationErr }] = useMutation(UPDATE_COMPANY);

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

  const onUpdate = async () => {
    try {
      await updateCompany({
        variables: {
          ...values
        }
      });
      onClose();
    } catch (err) {}
  };

  if (!open) {
    return null;
  }

  return (
    <Modal onClose={onClose} open={open}>
      <>
        <ErrorModal error={mutationErr} />
        <Card {...rest} className={clsx(classes.root, className)}>
          <form>
            <CardHeader title="Edit Company" />
            <Divider />
            <CardContent>
              <Grid container spacing={3}>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="Contact Email"
                    name="contactEmail"
                    onChange={handleFieldChange}
                    value={values.contactEmail}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="Company Name"
                    name="name"
                    onChange={handleFieldChange}
                    value={values.name}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="Phone number"
                    name="phone"
                    onChange={handleFieldChange}
                    value={values.phone}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="Address 1"
                    name="addressLine1"
                    onChange={handleFieldChange}
                    value={values.addressLine1}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="Address 2"
                    name="addressLine2"
                    onChange={handleFieldChange}
                    value={values.addressLine2}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="County"
                    name="county"
                    onChange={handleFieldChange}
                    value={values.county}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="County"
                    name="postCode"
                    onChange={handleFieldChange}
                    value={values.postCode}
                    variant="outlined"
                  />
                </Grid>
                <Grid item md={6} xs={12}>
                  <TextField
                    fullWidth
                    label="Country"
                    name="country"
                    onChange={handleFieldChange}
                    value={values.country}
                    variant="outlined"
                  />
                </Grid>
                <Grid item />
                <Grid item md={6} xs={12}>
                  <Typography variant="h5">Is Contract</Typography>
                  <Typography variant="body2">
                    Purchases will no longer need to be paid for through the
                    site
                  </Typography>
                  <Switch
                    checked={values.isContract}
                    color="secondary"
                    edge="start"
                    name="isContract"
                    onChange={handleFieldChange}
                    value={values.isContract}
                  />
                </Grid>
              </Grid>
            </CardContent>
            <Divider />
            <CardActions className={classes.actions}>
              <Button onClick={onClose}>Close</Button>
              <Button color="primary" onClick={onUpdate} variant="contained">
                Save
              </Button>
            </CardActions>
          </form>
        </Card>
      </>
    </Modal>
  );
}

CompanyEditModal.propTypes = {
  className: PropTypes.string,
  company: PropTypes.any,
  onClose: PropTypes.func,
  open: PropTypes.bool
};

CompanyEditModal.defaultProps = {
  open: false,
  onClose: () => {}
};

export default CompanyEditModal;
