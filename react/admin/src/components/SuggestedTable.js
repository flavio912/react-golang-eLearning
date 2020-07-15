import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
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

function SuggestedTable({ title, lessons }) {
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
            {lessons.map(lesson => (
                <TableRow key={lesson.uuid}>
                <TableCell>
                <Typography
                    className={classes.bold}
                    variant="subtitle2"
                    color="textPrimary"
                >
                    {lesson.name}
                </Typography>
                <Typography
                    className={classes.bold}
                    variant="body2"
                    color="textSecondary"
                >
                    Used in {lesson.numCoursesUsedIn} other Courses
                </Typography>
                </TableCell>
                <TableCell>
                    <Typography
                    className={classes.heavy}
                    variant="subtitle2"
                    color="textSecondary"
                    >
                    Type: 
                    </Typography>
                </TableCell>
                <TableCell>
                    <Typography
                    className={classes.bold}
                    variant="body2"
                    color="textPrimary"
                    >
                    {lesson.type}
                    </Typography>
                </TableCell>
                <TableCell>
                    <Typography
                    className={classes.heavy}
                    variant="subtitle2"
                    color="textSecondary"
                    >
                    Tags: 
                    </Typography>
                </TableCell>
                <TableCell>
                    {lesson.tags.map(tag => (
                    <Chip color={tag.color} label={tag.name} />
                    ))}
                </TableCell>
                <TableCell>
                    <Button
                    color="default"
                    component={RouterLink}
                    size="small"
                    variant="contained"
                    to={`/module/${lesson.uuid}/overview`}
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
