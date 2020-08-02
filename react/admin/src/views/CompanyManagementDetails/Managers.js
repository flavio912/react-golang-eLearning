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
  Button,
  TableCell,
  Avatar,
  TablePagination,
  Input,
  TableHead,
  TableRow
} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import AddUser from './AddUser';

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
  },
  options: {
    display: 'flex',
    justifyContent: 'flex-end',
    marginBottom: theme.spacing(2)
  }
}));

function Managers({ className, company, ...rest }) {
  const classes = useStyles();
  const [page, setPage] = useState(0);
  const [addUserModalOpen, setAddUserModalOpen] = useState(false);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  const { managers } = company;
  console.log('ma', managers);
  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  const onAddUserModalClose = () => {
    setAddUserModalOpen(false);
  };

  const openAddUserModal = () => {
    setAddUserModalOpen(true);
  };

  return (
    <div>
      <div className={classes.options}>
        <Button variant="contained" color="primary" onClick={openAddUserModal}>
          Add User
        </Button>
      </div>
      <Card {...rest} className={clsx(classes.root, className)}>
        <CardHeader title="Managers" />
        <Divider />
        <div className={classes.searchRow}>
          <div className={classes.search}>
            <SearchIcon className={classes.searchIcon} color="inherit" />
            <Input
              className={classes.searchInput}
              disableUnderline
              placeholder="Search managers"
            />
          </div>
        </div>
        <Divider />
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>User</TableCell>
              <TableCell>Last Login</TableCell>
              <TableCell>Created At</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {(managers?.edges ?? []).map(manager => (
              <TableRow key={manager.uuid}>
                <TableCell>
                  <div className={classes.nameCell}>
                    <Avatar className={classes.avatar} src={manager.logo}>
                      {getInitials(`${manager.firstName} ${manager.lastName}`)}
                    </Avatar>
                    <div>
                      <Link
                        color="inherit"
                        component={RouterLink}
                        to="/managers/1"
                        variant="h6"
                      >
                        {`${manager.firstName} ${manager.lastName}`}
                      </Link>
                      <div>{manager.email}</div>
                    </div>
                  </div>
                </TableCell>
                <TableCell>
                  {manager.lastLogin
                    ? moment(manager.lastLogin).format('LLL')
                    : 'Never'}
                </TableCell>
                <TableCell>{moment(manager.createdAt).format('LLL')}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <CardActions className={classes.actions}>
          <TablePagination
            component="div"
            count={managers.length}
            onChangePage={handleChangePage}
            onChangeRowsPerPage={handleChangeRowsPerPage}
            page={page}
            rowsPerPage={rowsPerPage}
            rowsPerPageOptions={[5, 10, 25]}
          />
        </CardActions>
      </Card>
      <AddUser
        open={addUserModalOpen}
        onClose={values => onAddUserModalClose(values)}
        company={company}
        userType="manager"
      />
    </div>
  );
}

Managers.propTypes = {
  className: PropTypes.string
};

export default Managers;
