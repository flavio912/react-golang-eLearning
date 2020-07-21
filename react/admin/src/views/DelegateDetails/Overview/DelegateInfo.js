import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import clsx from 'clsx';
import { makeStyles } from '@material-ui/styles';
import {
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Button,
  Divider,
  Table,
  TableBody,
  TableRow,
  TableCell,
  colors
} from '@material-ui/core';
import EditIcon from '@material-ui/icons/Edit';
import gql from 'graphql-tag';
import { useMutation } from '@apollo/react-hooks';
import Label from 'src/components/Label';
import DelegateEditModal from './DelegateEditModal';

const useStyles = makeStyles(theme => ({
  root: {},
  content: {
    padding: 0
  },
  actions: {
    flexDirection: 'column',
    alignItems: 'flex-start',
    '& > * + *': {
      marginLeft: 0
    }
  },
  buttonIcon: {
    marginRight: theme.spacing(1)
  }
}));

const UPDATE_DELEGATE = gql`
  mutation UpdateDelegate(
    $uuid: UUID!
    $firstName: String!
    $lastName: String!
    $jobTitle: String!
    $telephone: String
    $email: String
    $profileImageUrl: String
    $companyUUID: UUID
    $newPassword: String
  ) {
    updateDelegate(
      input: {
        uuid: $uuid
        firstName: $firstName
        lastName: $lastName
        jobTitle: $jobTitle
        telephone: $telephone
        email: $email
        profileImageUploadToken: $profileImageUrl
        companyUUID: $companyUUID
        newPassword: $newPassword
      }
    ) {
      uuid
      firstName
      lastName
      jobTitle
      telephone
      email
      profileImageUrl
      company {
        uuid
        name
      }
    }
  }
`;

function DelegateInfo({ delegate, className, ...rest }) {
  const classes = useStyles();
  const [delegateInfo, setDelegateInfo] = useState(delegate);
  const [openEdit, setOpenEdit] = useState(false);
  const [saveDelegate] = useMutation(UPDATE_DELEGATE);

  useEffect(() => {
    setDelegateInfo({
      ...delegate,
      isValid: true,
      errorMsg: null
    });
  }, [delegate]);

  const handleEditOpen = () => {
    setOpenEdit(true);
  };

  const handleEditClose = async values => {
    if (!values) {
      setOpenEdit(false);
      return;
    }

    try {
      const { data } = await saveDelegate({
        variables: {
          uuid: values.uuid,
          firstName: values.firstName,
          lastName: values.lastName,
          jobTitle: values.jobTitle,
          telephone: values.telephone,
          email: values.email,
          companyUUID: values.company.uuid,
          generatePassword: values.generatePassword
        }
      });
      setDelegateInfo({ ...data.updateDelegate });
      setOpenEdit(false);
    } catch (err) {
      setDelegateInfo({
        ...delegate,
        isValid: false,
        errorMsg: err?.graphQLErrors[0]?.message
      });
    }
  };

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Delegate info" />
      <Divider />
      <CardContent className={classes.content}>
        <Table>
          <TableBody>
            <TableRow>
              <TableCell>Email</TableCell>
              <TableCell>
                {delegateInfo.email}
                <div>
                  <Label
                    color={
                      delegateInfo.verified
                        ? colors.green[600]
                        : colors.orange[600]
                    }
                  >
                    {delegateInfo.verified
                      ? 'Email verified'
                      : 'Email not verified'}
                  </Label>
                </div>
              </TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>First Name</TableCell>
              <TableCell>{delegateInfo.firstName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Last Name</TableCell>
              <TableCell>{delegateInfo.lastName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Job Title</TableCell>
              <TableCell>{delegateInfo.jobTitle}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Telephone</TableCell>
              <TableCell>{delegateInfo.telephone}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Company</TableCell>
              <TableCell>{delegateInfo.company.name}</TableCell>
            </TableRow>
          </TableBody>
        </Table>
      </CardContent>
      <CardActions className={classes.actions}>
        <Button onClick={handleEditOpen}>
          <EditIcon className={classes.buttonIcon} />
          Edit
        </Button>
      </CardActions>
      <DelegateEditModal
        delegate={delegateInfo}
        onClose={values => handleEditClose(values)}
        open={openEdit}
      />
    </Card>
  );
}

DelegateInfo.propTypes = {
  className: PropTypes.string,
  delegate: PropTypes.object.isRequired
};

export default DelegateInfo;
