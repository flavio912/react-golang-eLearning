import React, { useState } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import moment from 'moment';
import {
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Button,
  Divider,
  Table,
  TableBody,
  TableRow,
  TableCell
} from '@material-ui/core';
import EditIcon from '@material-ui/icons/Edit';
import CustomerEditModal from './CustomerEditModal';

const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  actions: {
    flexDirection: 'column',
    alignItems: 'flex-start',
    '& > * + *': {
      marginLeft: 0
    }
  },
  buttonIcon: {
    marginRight: theme.spacing(1)
  }
}));

function CustomerInfo({ customer, className, ...rest }) {
  const classes = useStyles();
  const [openEdit, setOpenEdit] = useState(false);

  const handleEditOpen = () => {
    setOpenEdit(true);
  };

  const handleEditClose = () => {
    setOpenEdit(false);
  };
  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="User info" />
      <Divider />
      <CardContent className={classes.content}>
        <Table>
          <TableBody>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>{`${customer.firstName} ${customer.lastName}`}</TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>Username</TableCell>
              <TableCell>{customer.username}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Roles</TableCell>
              <TableCell>{customer.phone}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Email</TableCell>
              <TableCell>{customer.email}</TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>Job Title</TableCell>
              <TableCell>{customer.jobTitle}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Phone</TableCell>
              <TableCell>{customer.phone}</TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>Last Login</TableCell>
              <TableCell>{moment(customer.lastLogin).format('LLL')}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Created At</TableCell>
              <TableCell>{moment(customer.createdAt).format('LLL')}</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
      <CardActions className={classes.actions}>
        <Button onClick={handleEditOpen}>
          <EditIcon className={classes.buttonIcon} />
          Edit
        </Button>
      </CardActions>
      <CustomerEditModal
        customer={customer}
        onClose={handleEditClose}
        open={openEdit}
      />
    </Card>
  );
}

CustomerInfo.propTypes = {
  className: PropTypes.string,
  customer: PropTypes.object.isRequired
};

export default CustomerInfo;
