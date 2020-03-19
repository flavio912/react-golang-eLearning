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
  CardActions,
  CardContent,
  Divider,
  ButtonGroup,
  Table,
  TableBody,
  Link,
  TableCell,
  Avatar,
  TablePagination,
  Input,
  TableHead,
  TableRow,
  colors
} from '@material-ui/core';
import SearchIcon from '@material-ui/icons/Search';
import axios from 'src/utils/axios';
import Label from 'src/components/Label';
import GenericMoreButton from 'src/components/GenericMoreButton';
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

function Results({ className, tutors, ...rest }) {
  const classes = useStyles();

  const [toggle, setToggle] = useState('all');
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
            <TableCell>User</TableCell>
            <TableCell>CIN Number</TableCell>
            <TableCell>Last Login</TableCell>
            <TableCell>Created At</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {tutors.map(tutor => (
            <TableRow key={tutor.uuid}>
              <TableCell>
                <div className={classes.nameCell}>
                  <Avatar className={classes.avatar} src={tutor.logo}>
                    {getInitials(tutor.fullName)}
                  </Avatar>
                  <div>
                    <Link
                      color="inherit"
                      component={RouterLink}
                      to="/companies/1"
                      variant="h6"
                    >
                      {tutor.fullName}
                    </Link>
                    <div>{tutor.email}</div>
                  </div>
                </div>
              </TableCell>
              <TableCell>{tutor.cin}</TableCell>
              <TableCell>
                {moment(tutor.lastLogin.date).format('LLL')}
              </TableCell>
              <TableCell>{moment(tutor.createdAt).format('LLL')}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={tutors.length}
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
  className: PropTypes.string,
  tutors: PropTypes.object
};

export default Results;
