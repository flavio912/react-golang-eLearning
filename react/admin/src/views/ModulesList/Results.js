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
  Chip,
  Button,
} from '@material-ui/core';
import CreateOutlinedIcon from '@material-ui/icons/CreateOutlined';

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

function Results({ className, modules, ...rest }) {
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
            <TableCell>Courses Linked</TableCell>
            <TableCell>Lessons</TableCell>
            <TableCell>Tags</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {modules && modules.map(module => (
            <TableRow key={module.uuid}>
              <TableCell className={classes.moduleName}>
                {module.name}
              </TableCell>
              <TableCell>{module.numCoursesUsedIn}</TableCell>
              <TableCell>{module.syllabus.length}</TableCell>
              <TableCell>
                {module.tags && module.tags.map(tag => (
                  <Chip style={{ backgroundColor: tag.color }} label={tag.name} />
                ))}
              </TableCell>
              <TableCell>
                <Button
                  color="default"
                  component={RouterLink}
                  size="small"
                  to={`/modules/${module.uuid}/overview`}
                >
                  <CreateOutlinedIcon />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={modules.length}
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
  modules: PropTypes.array
};

export default Results;
