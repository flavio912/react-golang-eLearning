import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import {
  Card,
  CardActions,
  Table,
  TableBody,
  TableCell,
  TablePagination,
  TableHead,
  TableRow,
  Button,
  Link
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
  },
  moduleName: {}
}));

function Results({ className, companies, onApprove, ...rest }) {
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
            <TableCell>Name</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {companies.map(company => (
            <TableRow key={company.uuid}>
              <TableCell className={classes.name}>
                <Link
                  color="inherit"
                  component={RouterLink}
                  to={`/companies/${company.uuid}`}
                  variant="h6"
                >
                  {company.name}
                </Link>
              </TableCell>
              <TableCell>
                <Button
                  color="primary"
                  component={RouterLink}
                  size="small"
                  onClick={() => onApprove(company.uuid)}
                  variant="outlined"
                >
                  Approve
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={companies.length}
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
  companies: PropTypes.object
};

export default Results;
