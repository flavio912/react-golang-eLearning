import React from 'react';
import {
  Card,
  CardHeader,
  CardContent,
  Typography,
  Table,
  TableBody,
  TableRow,
  TableCell,
  Chip,
  Button,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
const useStyles = makeStyles(theme => ({
  header: {
    marginBottom: theme.spacing(1)
  },
  bold: {
    fontWeight: 'bold'
  },
  heavy: {
    fontWeight: 400
  },
  noPadding: {
    paddingBottom: 0
  },
}));

function SuggestedTable({ title, suggestions, onAdd }) {
  const classes = useStyles();

  return (
    <Card>
      <CardHeader
        title={title}
        className={classes.noPadding}
      />
      <CardContent>
        <Table>
          <TableBody>
            {suggestions && suggestions.map(suggestion => (
              <TableRow key={suggestion.uuid}>
                <TableCell>
                <Typography
                  className={classes.bold}
                  variant="subtitle2"
                  color="textPrimary"
                >
                  {suggestion.name}
                </Typography>
                <Typography
                  className={classes.bold}
                  variant="body2"
                  color="textSecondary"
                >
                  Used in {suggestion.numCoursesUsedIn} other Courses
                </Typography>
                </TableCell>
                <TableCell padding="none">
                  <Typography
                    className={classes.heavy}
                    variant="subtitle2"
                    color="textSecondary"
                    >
                    Type: 
                  </Typography>
                </TableCell>
                <TableCell padding="none">
                  <Typography
                    className={classes.bold}
                    variant="body1"
                    color="textPrimary"
                  >
                    {suggestion.type && suggestion.type.replace(/^\w/, (c) => c.toUpperCase())}
                  </Typography>
                </TableCell>
                <TableCell padding="none">
                  <Typography
                    className={classes.heavy}
                    variant="subtitle2"
                    color="textSecondary"
                  >
                    Tags: 
                  </Typography>
                </TableCell>
                <TableCell padding="none">
                  {suggestion.tags && suggestion.tags.map(tag => (
                    <Chip
                      key={tag.name}
                      style={{ backgroundColor: tag.color }}
                      label={tag.name}
                    />
                  ))}
                </TableCell>
                <TableCell align="right">
                  <Button
                    color="default"
                    onClick={() => onAdd(suggestion)}
                    >
                    + ADD
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
       </CardContent>
    </Card>
  );
}

export default SuggestedTable;
