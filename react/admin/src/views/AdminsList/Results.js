import React, { useState } from 'react';
import { useSelector } from 'react-redux';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import getInitials from 'src/utils/getInitials';
import {
  Card,
  CardActions,
  Table,
  TableBody,
  Link,
  TableCell,
  Avatar,
  TablePagination,
  TableHead,
  TableRow
} from '@material-ui/core';

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

function Results({ className, ...rest }) {
  const classes = useStyles();
  const adminState = useSelector(state => state.admin);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Email</TableCell>
            <TableCell>First Name</TableCell>
            <TableCell>Last Name</TableCell>
            <TableCell>Created At</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {adminState.listLoaded &&
            adminState.list.length > 0 &&
            adminState.list.map(admin => (
              <TableRow key={admin.uuid}>
                <TableCell>
                  <div className={classes.nameCell}>
                    <Avatar className={classes.avatar} src={admin.logo}>
                      {/* {getInitials(admin.fullName)} */}
                    </Avatar>
                    <Link
                      color="inherit"
                      component={RouterLink}
                      to={`/admins/${admin.uuid}/overview`}
                      variant="h6"
                    >
                      {admin.email}
                    </Link>
                  </div>
                </TableCell>
                <TableCell>{admin.firstName}</TableCell>
                <TableCell>{admin.lastName}</TableCell>
                <TableCell>{moment(admin.createdAt).format('LLL')}</TableCell>
              </TableRow>
            ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={adminState.list.length}
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

Results.propTypes = {
  className: PropTypes.string
};

export default Results;
