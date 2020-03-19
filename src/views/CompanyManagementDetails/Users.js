import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import getInitials from 'src/utils/getInitials';
import {
  Card,
  CardHeader,
  CardActions,
  Divider,
  Table,
  TableBody,
  Link,
  TableCell,
  Avatar,
  TablePagination,
  Input,
  TableHead,
  TableRow
} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import Label from 'src/components/Label';
import { ToggleButtonGroup, ToggleButton } from '@material-ui/lab';

const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  inner: {
    minWidth: 1150
  },
  search: {
    display: 'flex',
    alignItems: 'center',
    flexGrow: 1
  },
  searchRow: {
    padding: theme.spacing(2, 3),
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between'
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
  actions: {
    padding: theme.spacing(1),
    justifyContent: 'flex-end'
  },
  avatar: {
    height: 42,
    width: 42,
    marginRight: theme.spacing(2)
  }
}));

function Users({ className, ...rest }) {
  const classes = useStyles();

  const [toggle, setToggle] = useState('all');
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

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
    },
    {
      fullName: 'John Doe',
      userId: 'john_doe2',
      roles: ['Manager', 'Delegate'],
      email: 'tom@tom.com',
      noValidCerts: 4,
      noExpiringCerts: 2,
      lastLogin: {
        date: '02/01/2020'
      },
      createdAt: '05/01/2020'
    }
  ];

  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Managers and Delegates" />
      <Divider />
      <div className={classes.searchRow}>
        <div className={classes.search}>
          <SearchIcon className={classes.searchIcon} color="inherit" />
          <Input
            className={classes.searchInput}
            disableUnderline
            placeholder="Search users"
          />
        </div>
        <ToggleButtonGroup
          exclusive
          value={toggle}
          onChange={(_, value) => {
            setToggle(value);
          }}
        >
          <ToggleButton value="all">Show All</ToggleButton>
          <ToggleButton value="manager">Managers Only</ToggleButton>
        </ToggleButtonGroup>
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
            <TableRow key={user.userId}>
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
                {user.roles.map((role, index) => (
                  <Label style={{ marginRight: 5 }} key={index}>
                    {role}
                  </Label>
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
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={exampleUsers.length}
          onChangePage={handleChangePage}
          onChangeRowsPerPage={handleChangeRowsPerPage}
          page={page}
          rowsPerPage={rowsPerPage}
          rowsPerPageOptions={[5, 10, 25]}
        />
      </CardActions>
    </Card>
  );
}

Users.propTypes = {
  className: PropTypes.string
};

export default Users;
