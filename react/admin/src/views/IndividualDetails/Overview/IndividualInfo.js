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
import IndividualEditModal from './IndividualEditModal';

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

const UPDATE_INDIVIDUAL = gql`
  mutation UpdateIndividual(
    $uuid: UUID!
    $email: String
    $firstName: String
    $lastName: String
  ) {
    updateIndividual(
      input: {
        uuid: $uuid
        email: $email
        firstName: $firstName
        lastName: $lastName
      }
    ) {
      uuid
      email
      firstName
      lastName
    }
  }
`;

function IndividualInfo({ individual, className, ...rest }) {
  const classes = useStyles();
  const [individualInfo, setIndividualInfo] = useState(individual);
  const [openEdit, setOpenEdit] = useState(false);
  const [saveIndividual] = useMutation(UPDATE_INDIVIDUAL);

  useEffect(() => {
    setIndividualInfo({
      ...individual,
      isValid: true,
      errorMsg: null
    });
  }, [individual]);

  const handleEditOpen = () => {
    setOpenEdit(true);
  };

  const handleEditClose = async values => {
    if (!values) {
      setOpenEdit(false);
      return;
    }

    try {
      const { data } = await saveIndividual({
        variables: {
          uuid: values.uuid,
          email: values.email,
          firstName: values.firstName,
          lastName: values.lastName,
          password: values.password
        }
      });
      setIndividualInfo({ ...data.updateIndividual });
      setOpenEdit(false);
    } catch (err) {
      setIndividualInfo({
        ...individual,
        isValid: false,
        errorMsg: err?.graphQLErrors[0]?.message
      });
    }
  };

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Individual info" />
      <Divider />
      <CardContent className={classes.content}>
        <Table>
          <TableBody>
            <TableRow>
              <TableCell>Email</TableCell>
              <TableCell>
                {individualInfo.email}
                <div>
                  <Label
                    color={
                      individualInfo.verified
                        ? colors.green[600]
                        : colors.orange[600]
                    }
                  >
                    {individualInfo.verified
                      ? 'Email verified'
                      : 'Email not verified'}
                  </Label>
                </div>
              </TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>First Name</TableCell>
              <TableCell>{individualInfo.firstName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Last Name</TableCell>
              <TableCell>{individualInfo.lastName}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Job Title</TableCell>
              <TableCell>{individualInfo.jobTitle}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Phone Number</TableCell>
              <TableCell>{individualInfo.telephone}</TableCell>
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
      <IndividualEditModal
        individual={individualInfo}
        onClose={values => handleEditClose(values)}
        open={openEdit}
      />
    </Card>
  );
}

IndividualInfo.propTypes = {
  className: PropTypes.string,
  individual: PropTypes.object.isRequired
};

export default IndividualInfo;
