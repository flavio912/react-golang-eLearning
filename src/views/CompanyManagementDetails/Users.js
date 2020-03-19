import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import PerfectScrollbar from 'react-perfect-scrollbar';
import { makeStyles } from '@material-ui/styles';
import getInitials from 'src/utils/getInitials';
import {
  Button,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Table,
  TableBody,
  Link,
  TableCell,
  Avatar,
  Input,
  TableHead,
  TableRow,
  colors
} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import axios from 'src/utils/axios';
import Label from 'src/components/Label';
import GenericMoreButton from 'src/components/GenericMoreButton';

const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  inner: {
    minWidth: 1150
  },
  search: {
    padding: theme.spacing(2, 3),
    display: 'flex',
    alignItems: 'center'
  },
  searchIcon: {
    color: theme.palette.text.secondary
  },
  searchInput: {
    marginLeft: theme.spacing(1),
    color: theme.palette.text.secondary,
    fontSize: '14px'
  },
  nameCell: {
    display: 'flex',
    alignItems: 'center'
  },
  avatar: {
    height: 42,
    width: 42,
    marginRight: theme.spacing(2)
  }
}));

function Users({ className, ...rest }) {
  const classes = useStyles();

  const exampleUsers = [
    {
      fullName: 'Tom Emmerson',
      userId: 'tom_emmerson',
      roles: ['Manager'],
      email: 'tom@tom.com',
      noValidCerts: 4,
      noExpiringCerts: 2,
      lastLogin: {
        date: '02/01/2020'
      },
      createdAt: '05/01/2020'
    }
  ];

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Managers and Delegates" />
      <Divider />
      <div className={classes.search}>
        <SearchIcon className={classes.searchIcon} color="inherit" />
        <Input
          className={classes.searchInput}
          disableUnderline
          placeholder="Search users"
        />
      </div>
      <Divider />
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>User</TableCell>
            <TableCell>Roles</TableCell>
            <TableCell>Valid Certificates</TableCell>
            <TableCell>Expiring Certificates</TableCell>
            <TableCell>Last Login</TableCell>
            <TableCell>Created At</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {exampleUsers.map(user => (
            <TableRow>
              <TableCell>
                <div className={classes.nameCell}>
                  <Avatar className={classes.avatar} src={user.logo}>
                    {getInitials(user.fullName)}
                  </Avatar>
                  <div>
                    <Link
                      color="inherit"
                      component={RouterLink}
                      to="/companies/1"
                      variant="h6"
                    >
                      {user.fullName}
                    </Link>
                    <div>{user.email}</div>
                  </div>
                </div>
              </TableCell>
              <TableCell>
                {user.roles.map(role => (
                  <Label>{role}</Label>
                ))}
              </TableCell>
              <TableCell>{user.noValidCerts}</TableCell>
              <TableCell>{user.noExpiringCerts}</TableCell>
              <TableCell>{moment(user.lastLogin.date).format('LLL')}</TableCell>
              <TableCell>{moment(user.createdAt).format('LLL')}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </Card>
  );
}

Users.propTypes = {
  className: PropTypes.string
};

export default Users;
