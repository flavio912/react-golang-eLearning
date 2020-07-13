import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
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
  TableCell,
  colors
} from '@material-ui/core';
import EditIcon from '@material-ui/icons/Edit';
import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';
import Label from 'src/components/Label';
import AdminEditModal from './AdminEditModal';

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

const UPDATE_ADMIN = gql`
  mutation UpdateAdmin(
    $uuid: UUID!
    $email: String
    $firstName: String
    $lastName: String
  ) {
    updateAdmin(
      input: {
        uuid: $uuid
        email: $email
        firstName: $firstName
        lastName: $lastName
      }
    ) {
      uuid
      email
      firstName
      lastName
    }
  }
`;

function AdminInfo({ admin, className, ...rest }) {
  const classes = useStyles();
  const [adminInfo, setAdminInfo] = useState(admin);
  const [openEdit, setOpenEdit] = useState(false);
  const [saveAdmin] = useMutation(UPDATE_ADMIN);

  useEffect(() => {
    setAdminInfo({
      ...admin,
      isValid: true,
      errorMsg: null
    });
  }, [admin]);

  const handleEditOpen = () => {
    setOpenEdit(true);
  };

  const handleEditClose = async values => {
    if (!values) {
      setOpenEdit(false);
      return;
    }

    try {
      const { data } = await saveAdmin({
        variables: {
          uuid: values.uuid,
          email: values.email,
          firstName: values.firstName,
          lastName: values.lastName,
          password: values.password
        }
      });
      setAdminInfo({ ...data.updateAdmin });
      setOpenEdit(false);
    } catch (err) {
      setAdminInfo({
        ...admin,
        isValid: false,
        errorMsg: err?.graphQLErrors[0]?.message
      });
    }
  };

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Admin info" />
      <Divider />
      <CardContent className={classes.content}>
        <Table>
          <TableBody>
            <TableRow>
              <TableCell>Email</TableCell>
              <TableCell>
                {adminInfo.email}
                <div>
                  <Label
                    color={
                      adminInfo.verified
                        ? colors.green[600]
                        : colors.orange[600]
                    }
                  >
                    {adminInfo.verified
                      ? 'Email verified'
                      : 'Email not verified'}
                  </Label>
                </div>
              </TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>First Name</TableCell>
              <TableCell>{adminInfo.firstName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Last Name</TableCell>
              <TableCell>{adminInfo.lastName}</TableCell>
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
      <AdminEditModal
        admin={adminInfo}
        onClose={values => handleEditClose(values)}
        open={openEdit}
      />
    </Card>
  );
}

AdminInfo.propTypes = {
  className: PropTypes.string,
  admin: PropTypes.object.isRequired
};

export default AdminInfo;
