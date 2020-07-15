import React, { useState } from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import TagsInput from 'src/components/TagsInput';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  lessonName: {
    marginBottom: 20
  },
  buttonText: {
    color: '#4a4a4a',
    fontSize: 11,
    fontWeight: 'weight: 700'
  },
  shortDescription: {
    width: '100%'
  },
  answerItem: {
    border: '1px solid gainsboro',
    borderRadius: 3,
    padding: '6px 21px',
    justifyContent: 'space-between',
    alignItems: 'center',
    display: 'flex'
  },
  formControl: {
    width: '100%'
  },
  previewImage: {
    width: 200,
    maxHeight: 200
  },
  editItems: {
    alignItems: 'center'
  }
}));

function Overview({ state, setState }) {
  const classes = useStyles();

  const handleTags = () => {};

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Card>
              <CardHeader title={'Lesson Information'} />
              <Divider />
              <CardContent>
                <Grid item className={classes.lessonName}>
                  <TextField
                    fullWidth
                    label="Name"
                    name="lesson"
                    value={state.name}
                    onChange={inp => {
                      setState('name', inp.target.value);
                    }}
                    placeholder="Lesson Name"
                    variant="outlined"
                  />
                </Grid>
                <TagsInput onChange={handleTags} />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title={'Lesson Description'} />
              <Divider />
              <CardContent>
                <Grid item>
                  <TextField
                    fullWidth
                    multiline
                    rows={8}
                    label="Description"
                    name="description"
                    value={state.text}
                    onChange={inp => {
                      setState('description', inp.target.value);
                    }}
                    placeholder="Lesson Description"
                    variant="outlined"
                  />
                </Grid>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
