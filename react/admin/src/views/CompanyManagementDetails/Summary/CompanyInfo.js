import React, { useState } from 'react';
import PropTypes from 'prop-types';
import Imgix from 'react-imgix';
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
  TableCell
} from '@material-ui/core';
import EditIcon from '@material-ui/icons/Edit';
import CompanyEditModal from './CompanyEditModal';
import Label from 'src/components/Label';

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

function CustomerInfo({ company, onUpdate, className, ...rest }) {
  const classes = useStyles();
  const [openEdit, setOpenEdit] = useState(false);

  const handleEditOpen = () => {
    setOpenEdit(true);
  };

  const handleEditClose = () => {
    onUpdate();
    setOpenEdit(false);
  };

  return (
    <Card {...rest} className={clsx(classes.root, className)}>
      <CardHeader title="Company info" />
      <Divider />
      <CardContent className={classes.content}>
        <Table>
          <TableBody>
            <TableRow>
              <TableCell>Contact Email</TableCell>
              <TableCell>{company.contactEmail}</TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>Address 1</TableCell>
              <TableCell>{company.address.addressLine1}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Address 2</TableCell>
              <TableCell>{company.address.addressLine2}</TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>County</TableCell>
              <TableCell>{company.address.county}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Post Code</TableCell>
              <TableCell>{company.address.postCode}</TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>Country</TableCell>
              <TableCell>{company.address.country}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Company Type</TableCell>
              <TableCell>
                <Label color={'rgb(119, 125, 156)'}>
                  {company.isContract ? 'Contract' : 'Pay as you go'}
                </Label>
              </TableCell>
            </TableRow>
            <TableRow selected>
              <TableCell>Logo</TableCell>
              <TableCell>
                {company.logoURL && <Imgix src={company.logoURL} height={50} />}
              </TableCell>
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
      <CompanyEditModal
        company={company}
        onClose={handleEditClose}
        open={openEdit}
      />
    </Card>
  );
}

CustomerInfo.propTypes = {
  className: PropTypes.string,
  customer: PropTypes.object.isRequired
};

export default CustomerInfo;
