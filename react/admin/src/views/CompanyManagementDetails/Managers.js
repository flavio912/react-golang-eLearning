import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import PerfectScrollbar from 'react-perfect-scrollbar';
import { makeStyles } from '@material-ui/styles';
import {
  Button,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  colors
} from '@material-ui/core';
import axios from 'src/utils/axios';
import Label from 'src/components/Label';
import GenericMoreButton from 'src/components/GenericMoreButton';

const useStyles = makeStyles(() => ({
  root: {},
  content: {
    padding: 0
  },
  inner: {
    minWidth: 1150
  }
}));

function Invoices({ className, ...rest }) {
  const classes = useStyles();
  const [invoices, setInvoices] = useState([]);

  useEffect(() => {
    let mounted = true;

    const fetchInvoices = () => {
      axios.get('/api/management/customers/1/invoices').then(response => {
        if (mounted) {
          setInvoices(response.data.invoices);
        }
      });
    };

    fetchInvoices();

    return () => {
      mounted = false;
    };
  }, []);

  const statusColors = {
    pending: colors.orange[600],
    paid: colors.green[600],
    rejected: colors.red[600]
  };

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Card>
        <CardHeader action={<GenericMoreButton />} title="Company Managers" />
        <Divider />
        <CardContent className={classes.content}>
          <PerfectScrollbar>
            <div className={classes.inner}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell>Email</TableCell>
                    <TableCell>Job Title</TableCell>
                    <TableCell>Telephone</TableCell>
                    <TableCell>Last Login</TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {managers.map(manager => (
                    <TableRow key={manager.id}>
                      <TableCell>{`${manager.firstName} ${manager.lastName}`}</TableCell>
                      <TableCell>
                        {moment(manager.lastLogin).format('DD/MM/YYYY | HH:MM')}
                      </TableCell>

                      <TableCell align="right">
                        <Button
                          color="primary"
                          component={RouterLink}
                          size="small"
                          to={'/management/invoices/1'}
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
      </Card>
    </div>
  );
}

Invoices.propTypes = {
  className: PropTypes.string
};

export default Invoices;
