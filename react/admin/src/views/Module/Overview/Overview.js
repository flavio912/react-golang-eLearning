import React, { useState } from 'react';
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

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
  },
  shortDescription: {
    width: '100%'
  }
}));

function Overview({}) {
  const classes = useStyles();

  const onChange=()=>{};
  const [name, setName] = React.useState();
  const [description, setDescription] = React.useState();

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
                            onChange(inp.target.value);
                        }}
                        placeholder="e.g. Fire Safety Module 1"
                        value={name}
                        variant="outlined"
                        />
                    </Grid>
                    <Grid item>
                    <TagsInput />
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
                    value={description}
                    onChange={inp => {
                        setDescription(inp.target.value);
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
                    <TextField
                    label=""
                    multiline
                    className={classes.shortDescription}
                    rows={5}
                    value={description}
                    onChange={inp => {
                        setDescription(inp.target.value);
                    }}
                    placeholder="Description"
                    variant="outlined"
                    />
                </CardContent>
                </Card>
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
