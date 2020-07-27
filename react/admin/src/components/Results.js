import React from 'react';
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
  TableRow
} from '@material-ui/core';

const useStyles = makeStyles(theme => ({
  actions: {
    padding: theme.spacing(1),
    justifyContent: 'flex-end'
  },
}));

function Results({ results, handleChangePage, handleChangeRowsPerPage, headers, cells, className, ...rest }) {
  const classes = useStyles();

  if (!results) return <div>Loading...</div>

  return (
    <Card {...rest} className={className}>
      <Table>
        <TableHead>
          <TableRow>
            {headers && headers.map((header, i) => (
              <TableCell key={i}>{header}</TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {results.edges && results.edges.map(result => (
            <TableRow key={result.uuid}>
              {cells && cells.map((cell, i) => (
                <TableCell key={i}>
                  {cell.component
                    ? cell.component(result)
                    : result[cell.field]
                  }
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <CardActions className={classes.actions}>
        <TablePagination
          component="div"
          count={results.pageInfo && results.pageInfo.total}
          onChangePage={handleChangePage}
          onChangeRowsPerPage={handleChangeRowsPerPage}
          page={results.pageInfo && results.pageInfo.offset}
          rowsPerPage={results.pageInfo && results.pageInfo.limit}
          rowsPerPageOptions={[5, 10, 25]}
        />
      </CardActions>
    </Card>
  );
}

Results.propTypes = {
  results: PropTypes.exact({
    edges: PropTypes.array,
    pageInfo: PropTypes.exact({
      total: PropTypes.number,
      offset: PropTypes.number,
      limit: PropTypes.number,
      given: PropTypes.number,
      __typename: PropTypes.string,
    }),
    __typename: PropTypes.string,
  }),
  className: PropTypes.string,
};

export default Results;
