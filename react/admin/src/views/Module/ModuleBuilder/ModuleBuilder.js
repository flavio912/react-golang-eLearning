import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Typography,
  TextField,
  InputAdornment,
  Table,
  TableBody,
  TableRow,
  TableCell,
  Chip,
  Button,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import Autocomplete from '@material-ui/lab/Autocomplete';
import SearchIcon from '@material-ui/icons/Search';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  heading: {
    margin: theme.spacing(2)
  },
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

const lessons = [
  {
    uuid: '1231231231',
    name: 'Firesafety Module 1 - Lesson 1',
    numCoursesUsedIn: 3,
    type: 'Lesson',
    tags: [
      {
        name: 'FIRE SAFETY',
        color: '#123'
      },
      {
        name: 'HEALTH & SAFETY',
        color: '#123'
      }
    ]
  },
  {
    uuid: '1231231231',
    name: 'Firesafety Module 1 - Lesson 2',
    numCoursesUsedIn: 3,
    type: 'Lesson',
    tags: [
      {
        name: 'FIRE SAFETY',
        color: '#123'
      }
    ]
  },
  {
    uuid: '1231231231',
    name: 'Firesafety Module 1 - Lesson 3',
    numCoursesUsedIn: 3,
    type: 'Lesson',
    tags: [
      {
        name: 'FIRE SAFETY',
        color: '#123'
      }
    ]
  }
];

function ModuleBuilder({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item className={classes.heading}>
            <Typography
              variant="h3"
              color="textPrimary"
              className={classes.header}
            >
              Build this Module
            </Typography>
            <Typography
              variant="body1"
              color="textPrimary"
            >
              All modules are comprised of <strong>Lessons</strong> and <strong>Tests</strong><br />
              Add your first lesson to get started
            </Typography>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader
                title="Search Lessons and Tests"
                className={classes.noPadding}
              />
              <CardContent>
              <Autocomplete
                freeSolo
                options={[]}
                renderInput={(params) => (
                  <TextField
                    {...params}
                    placeholder="Search Lessons or Tests"
                    InputProps={{
                      startAdornment: (
                      <InputAdornment position="start">
                        <SearchIcon />
                      </InputAdornment>
                      )
                    }}
                  /> 
                )}
              />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader
                title="Suggested Lessons based on Tags"
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
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Module Structure" />
              <Divider />
              <CardContent>
                <div>TBA</div>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default ModuleBuilder;
