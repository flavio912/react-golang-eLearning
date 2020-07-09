import React, { useState } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import getInitials from 'src/utils/getInitials';
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
  TableRow,
  Chip,
  Button
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

function Results({ className, questions, ...rest }) {
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
            <TableCell>Tests Linked</TableCell>
            {/* <TableCell>Lessons</TableCell> */}
            <TableCell>Possible Answers</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {questions.map(question => (
            <TableRow key={question.uuid}>
              <TableCell className={classes.questionName}>
                {question.text}
              </TableCell>
              <TableCell>{question.numTestsUsedIn}</TableCell>
              {/* <TableCell>{question.numLessons}</TableCell> */}
              {/* <TableCell>
                {question.tags.map(tag => (
                  <Chip color={tag.color} label={tag.name} />
                ))}
              </TableCell> */}
              <TableCell>{question.numAnswers}</TableCell>
              <TableCell>
                <Button
                  color="primary"
                  component={RouterLink}
                  size="small"
                  to={`/question/${question.uuid}/overview`}
                  variant="outlined"
                >
                  Edit
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={questions.length}
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
  questions: PropTypes.object
};

export default Results;
