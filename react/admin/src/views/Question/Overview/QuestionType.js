import React from 'react';
import {
  Grid,
  TextField,
  Card,
  CardHeader,
  CardContent,
  Divider,
  FormControl,
  MenuItem,
  InputLabel,
  Select,
  Typography,
  Button
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  formControl: {
    width: '100%'
  }
}));

function QuestionType({ name, onChange }) {
  const classes = useStyles();

  return (
    <Card>
      <CardHeader title={'Write your question here'} />
      <Divider />
      <CardContent>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <TextField
              fullWidth
              label="Question"
              name="question"
              onChange={inp => {
                onChange(inp.target.value);
              }}
              placeholder="Question text"
              value={name}
              variant="outlined"
            />
          </Grid>
          <Grid item>
            <Typography variant="h6">Question Type</Typography>
          </Grid>
          <Grid item>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel id="demo-simple-select-outlined-label">
                Question Type
              </InputLabel>
              <Select
                labelId="demo-simple-select-outlined-label"
                id="demo-simple-select-outlined"
                value={1}
                onChange={() => {}}
                label="Question Type"
              >
                <MenuItem value={1}>Single Choice</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item>
            <Typography variant="h6">Randomise</Typography>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default QuestionType;
