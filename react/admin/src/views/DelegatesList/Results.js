import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
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

function Results({ delegates, className, ...rest }) {
  const classes = useStyles();
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
            <TableCell>Job Title</TableCell>
            <TableCell>Telephone</TableCell>
            <TableCell>Company</TableCell>
            <TableCell>Created At</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {delegates.map(delegate => (
            <TableRow key={delegate.uuid}>
              <TableCell>
                <div className={classes.nameCell}>
                  <Avatar className={classes.avatar} src={delegate.logo}>
                    {delegate.firstName.charAt(0).toUpperCase()}
                    {delegate.lastName.charAt(0).toUpperCase()}
                  </Avatar>
                  <div>
                    <Link
                      color="inherit"
                      component={RouterLink}
                      to={`/delegates/${delegate.uuid}/overview`}
                      variant="h6"
                    >
                      {delegate.firstName} {delegate.lastName}
                    </Link>
                    <div>{delegate.email}</div>
                  </div>
                </div>
              </TableCell>
              <TableCell>{delegate.jobTitle}</TableCell>
              <TableCell>{delegate.telephone}</TableCell>
              <TableCell>{delegate.company.name}</TableCell>
              <TableCell>{moment(delegate.createdAt).format('LLL')}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={delegates.length}
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
