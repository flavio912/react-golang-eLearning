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
import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';
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

const CREATE_DELEGATE = gql`
  mutation CreateDelegate(
    $firstName: String!
    $lastName: String!
    $jobTitle: String!
    $email: String
    $companyUUID: UUID
    $generatePassword: Boolean
  ) {
    createDelegate(
      input: {
        firstName: $firstName
        lastName: $lastName
        jobTitle: $jobTitle
        email: $email
        companyUUID: $companyUUID
        generatePassword: $generatePassword
      }
    ) {
      delegate {
        uuid
        firstName
        lastName
        jobTitle
        email
      }
    }
  }
`;

function Delegates({ className, company, onUpdateCompany, ...rest }) {
  const classes = useStyles();

  const [page, setPage] = useState(0);
  const [addUserModalOpen, setAddUserModalOpen] = useState(false);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [createDelegate] = useMutation(CREATE_DELEGATE);

  const handleChangePage = (event, page) => {
    setPage(page);
  };

  const handleChangeRowsPerPage = event => {
    setRowsPerPage(event.target.value);
  };

  const onAddUserModalClose = async values => {
    setAddUserModalOpen(false);
    if (!values) return;

    await createDelegate({
      variables: {
        firstName: values.firstName,
        lastName: values.lastName,
        email: values.email,
        jobTitle: 'delegate',
        companyUUID: company.uuid,
        generatePassword: true
      }
    });
    onUpdateCompany();
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
        <CardHeader title="Delegates" />
        <Divider />
        <div className={classes.searchRow}>
          <div className={classes.search}>
            <SearchIcon className={classes.searchIcon} color="inherit" />
            <Input
              className={classes.searchInput}
              disableUnderline
              placeholder="Search delegates"
            />
          </div>
        </div>
        <Divider />
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>User</TableCell>
              <TableCell>Valid Certificates</TableCell>
              <TableCell>Expiring Certificates</TableCell>
              <TableCell>Last Login</TableCell>
              <TableCell>Created At</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {company?.delegates?.edges &&
              company?.delegates?.edges.length > 0 &&
              company?.delegates?.edges.map(user => (
                <TableRow key={user.uuid}>
                  <TableCell>
                    <div className={classes.nameCell}>
                      <Avatar className={classes.avatar} src={user.logo}>
                        {getInitials(user.fullName)}
                      </Avatar>
                      <div>
                        <Link
                          color="inherit"
                          component={RouterLink}
                          to={`/users/${user.uuid}`}
                          variant="h6"
                        >
                          {user.firstName} {user.lastName}
                        </Link>
                        <div>{user.email}</div>
                      </div>
                    </div>
                  </TableCell>
                  <TableCell>{user.noValidCerts}</TableCell>
                  <TableCell>{user.noExpiringCerts}</TableCell>
                  <TableCell>
                    {moment(user.lastLogin.date).format('LLL')}
                  </TableCell>
                  <TableCell>{moment(user.createdAt).format('LLL')}</TableCell>
                </TableRow>
              ))}
          </TableBody>
        </Table>
        <CardActions className={classes.actions}>
          <TablePagination
            component="div"
            count={company?.delegates?.edges.length}
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
        userType="delegate"
      />
    </div>
  );
}

Delegates.propTypes = {
  className: PropTypes.string,
  company: PropTypes.object,
  onUpdateCompany: PropTypes.func
};

export default Delegates;
