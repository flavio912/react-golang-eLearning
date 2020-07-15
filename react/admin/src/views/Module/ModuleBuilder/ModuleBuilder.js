import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  Typography,
  TextField,
  InputAdornment,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import Autocomplete from '@material-ui/lab/Autocomplete';
import SearchIcon from '@material-ui/icons/Search';
import ReoderableListItem from 'src/components/ReorderableList/ReorderableListItem';
import ReoderableDropdown from 'src/components/ReorderableList/ReorderableDropdown';
import SuggestedTable from 'src/components/SuggestedTable';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  heading: {
    margin: theme.spacing(2)
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

  const onDelete = () => {};

  const [ items, setItems ] = React.useState([
    {id: 0, component: <ReoderableListItem text="text" onDelete={onDelete} />},
    {id: 1, component: <ReoderableListItem text="text" onDelete={onDelete} />},
    {id: 2, component: <ReoderableListItem text="text" onDelete={onDelete} />},
  ]);

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
            <SuggestedTable
              title="Suggested Lessons based on Tags"
              lessons={lessons}
            />
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title="Module Structure" />
              <Divider />
              <CardContent>
                <ReoderableDropdown title="Module 1" items={items} setItems={setItems} />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default ModuleBuilder;
