import React, { useState } from 'react';
import {
  Card,
  CardHeader,
  Typography,
  TextField,
  CardContent,
  Container,
  Button,
  Grid,
  Chip,
  ButtonGroup,
  GridList,
  Radio,
  RadioGroup,
  InputLabel,
  FormControlLabel,
  FormLabel,
  Divider,
  Switch,
  IconButton
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import Clear from '@material-ui/icons/Clear';
import { Autocomplete } from '@material-ui/lab';

const useStyles = makeStyles(theme => ({}));

function Learn() {
  const classes = useStyles();

  const [bulletPoints, setBulletPoints] = useState([
    { text: '', key: new Date().getTime() }
  ]);

  const addNewRepeater = () => {
    const newBulletPoints = [...bulletPoints];
    newBulletPoints.push({ text: '', key: new Date().getTime() });
    setBulletPoints(newBulletPoints);
  };

  const removeRepeater = index => {
    if (index === 0 && bulletPoints.length == 1) return;
    const newBulletPoints = [...bulletPoints];
    newBulletPoints.splice(index, 1);
    setBulletPoints(newBulletPoints);
  };

  return (
    <Card>
      <CardHeader title={"What they'll learn"} />
      <Divider />
      <CardContent>
        <Grid container spacing={2} direction={'column'}>
          {bulletPoints.map((bullet, index) => (
            <Grid item key={bullet.key}>
              <Grid container spacing={1}>
                <Grid item xs={11}>
                  <TextField
                    fullWidth
                    multiline
                    label={`Bullet point ${index + 1}`} // Normal people don't do 0 indexing. ;)
                    variant={'outlined'}
                  />
                </Grid>
                <Grid item xs={1}>
                  <IconButton
                    onClick={() => {
                      removeRepeater(index);
                    }}
                  >
                    <Clear />
                  </IconButton>
                </Grid>
              </Grid>
            </Grid>
          ))}
          <Grid item>
            <Button
              onClick={addNewRepeater}
              variant={'outlined'}
              color={'primary'}
            >
              Add bullet point
            </Button>
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default Learn;
