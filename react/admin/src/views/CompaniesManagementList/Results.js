import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import PropTypes from 'prop-types';
import PerfectScrollbar from 'react-perfect-scrollbar';
import { makeStyles } from '@material-ui/styles';
import {
  Avatar,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Checkbox,
  Divider,
  Button,
  Link,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TablePagination,
  TableRow,
  Typography
} from '@material-ui/core';
import getInitials from 'src/utils/getInitials';
import GenericMoreButton from 'src/components/GenericMoreButton';
import TableEditBar from 'src/components/TableEditBar';
import Label from 'src/components/Label';

const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  inner: {
    minWidth: 700
  },
  nameCell: {
    display: 'flex',
    alignItems: 'center'
  },
  avatar: {
    height: 42,
    width: 42,
    marginRight: theme.spacing(1)
  },
  actions: {
    padding: theme.spacing(1),
    justifyContent: 'flex-end'
  }
}));

function Results({ className, companies, ...rest }) {
  const classes = useStyles();
  const [selectedCompanies, setSelectedCompanies] = useState([]);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  const handleSelectAll = event => {
    const selectedCompanies = event.target.checked
      ? companies.map(company => company.id)
      : [];

    setSelectedCompanies(selectedCompanies);
  };

  const handleSelectOne = (event, id) => {
    const selectedIndex = selectedCompanies.indexOf(id);
    let newSelectedCompanies = [];

    if (selectedIndex === -1) {
      newSelectedCompanies = newSelectedCompanies.concat(selectedCompanies, id);
    } else if (selectedIndex === 0) {
      newSelectedCompanies = newSelectedCompanies.concat(
        selectedCompanies.slice(1)
      );
    } else if (selectedIndex === selectedCompanies.length - 1) {
      newSelectedCompanies = newSelectedCompanies.concat(
        selectedCompanies.slice(0, -1)
      );
    } else if (selectedIndex > 0) {
      newSelectedCompanies = newSelectedCompanies.concat(
        selectedCompanies.slice(0, selectedIndex),
        selectedCompanies.slice(selectedIndex + 1)
      );
    }

    setSelectedCompanies(newSelectedCompanies);
  };

  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Typography color="textSecondary" gutterBottom variant="body2">
        {companies.length} Records found. Page {page + 1} of{' '}
        {Math.ceil(companies.length / rowsPerPage)}
      </Typography>
      <Card>
        <CardHeader action={<GenericMoreButton />} title="All companies" />
        <Divider />
        <CardContent className={classes.content}>
          <PerfectScrollbar>
            <div className={classes.inner}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell padding="checkbox">
                      <Checkbox
                        checked={selectedCompanies.length === companies.length}
                        color="primary"
                        indeterminate={
                          selectedCompanies.length > 0 &&
                          selectedCompanies.length < companies.length
                        }
                        onChange={handleSelectAll}
                      />
                    </TableCell>
                    <TableCell>Name</TableCell>
                    <TableCell>No. Delegates</TableCell>
                    <TableCell>No. Managers</TableCell>
                    <TableCell>Payment Type</TableCell>
                    <TableCell align="right">Actions</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {companies.slice(0, rowsPerPage).map(company => (
                    <TableRow
                      hover
                      key={company.id}
                      selected={selectedCompanies.indexOf(company.id) !== -1}
                    >
                      <TableCell padding="checkbox">
                        <Checkbox
                          checked={selectedCompanies.indexOf(company.id) !== -1}
                          color="primary"
                          onChange={event => handleSelectOne(event, company.id)}
                          value={selectedCompanies.indexOf(company.id) !== -1}
                        />
                      </TableCell>
                      <TableCell>
                        <div className={classes.nameCell}>
                          <Avatar className={classes.avatar} src={company.logo}>
                            {getInitials(company.name)}
                          </Avatar>
                          <div>
                            <Link
                              color="inherit"
                              component={RouterLink}
                              to={`/companies/${company.id}`}
                              variant="h6"
                            >
                              {company.name}
                            </Link>
                            <div>{company.email}</div>
                          </div>
                        </div>
                      </TableCell>
                      <TableCell>{company.noDelegates}</TableCell>
                      <TableCell>{company.noManagers}</TableCell>
                      <TableCell>
                        <Label color={'rgb(119, 125, 156)'}>
                          {company.paymentType}
                        </Label>
                      </TableCell>
                      <TableCell align="right">
                        <Button
                          color="primary"
                          component={RouterLink}
                          size="small"
                          to={`/companies/${company.id}`}
                          variant="outlined"
                        >
                          View
                        </Button>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </div>
          </PerfectScrollbar>
        </CardContent>
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
      <TableEditBar selected={selectedCompanies} />
    </div>
  );
}

Results.propTypes = {
  className: PropTypes.string,
  companies: PropTypes.array
};

Results.defaultProps = {
  companies: []
};

export default Results;
