import React from 'react';
import {
  Grid,
  Card,
  CardHeader,
  CardContent,
  Divider,
  TextField,
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import TagsInput from 'src/components/TagsInput';
import FilesDropzone from 'src/components/FilesDropzone';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  shortDescription: {
    width: '100%'
  }
}));

function Overview({ state, setState }) {
  const classes = useStyles();

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid item={6} xs={8}>
            <Grid container spacing={2} direction={'column'}>
            <Grid item>
                <Card>
                <CardHeader title="About this Module" />
                <Divider />
                <CardContent>
                  <Grid container spacing={4} direction={'column'}>
                    <Grid item>
                      <TextField
                        fullWidth
                        label="Module Name"
                        name="modulename"
                        onChange={inp => {
                            setState('name', inp.target.value);
                        }}
                        placeholder="e.g. Fire Safety Module 1"
                        value={state.name}
                        variant="outlined"
                      />
                    </Grid>
                    <Grid item>
                      <TagsInput onChange={(tags) => setState('tags', tags)} />
                    </Grid>
                  </Grid>
                </CardContent>
                </Card>
            </Grid>
            <Grid item>
                <Card>
                <CardHeader title="Module Description" />
                <Divider />
                <CardContent>
                    <TextField
                    label=""
                    multiline
                    className={classes.shortDescription}
                    rows={5}
                    value={state.description}
                    onChange={inp => {
                      setState('description', inp.target.value);
                    }}
                    placeholder="Description"
                    variant="outlined"
                    />
                </CardContent>
                </Card>
            </Grid>
            </Grid>
        </Grid>
        <Grid item xs={4}>
          <Card>
            <CardHeader title="Module Banner Image" />
            <Divider />
            <CardContent>
              <FilesDropzone onUpload={(files) => console.log(files)}/>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
