import React, { useState } from 'react';
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
  Radio,
  RadioGroup,
  FormControlLabel,
  Button
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import TagsInput from 'src/components/TagsInput';
import QuestionType from './QuestionType';
import AnswerInput from './AnswerInput';

const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1
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
  }
}));

function Overview({ state, setState }) {
  const classes = useStyles();

  const [showAnswerInput, setShowAnswerInput] = useState(false);
  const [answer, setAnswer] = useState({
    answerType: 'TEXT',
    text: ''
  });

  const onSaveAnswer = () => {
    setState('answers', [...state.answers, answer]);
    setShowAnswerInput(false);
    setAnswer({});
  };

  const getCorrectAnswer = () => {
    for (let i = 0; i < state.answers.length; i++) {
      if (state.answers[i].isCorrect) {
        return i;
      }
    }
    return 0;
  };

  return (
    <div className={classes.root}>
      <Grid container spacing={2}>
        <Grid container spacing={4} direction={'column'}>
          <Grid item>
            <Card>
              <CardHeader title={'Question Information'} />
              <Divider />
              <CardContent>
                <TagsInput />
              </CardContent>
            </Card>
          </Grid>
          <Grid item>
            <QuestionType
              name={state.name}
              onChange={val => setState('name', val)}
            />
          </Grid>
          <Grid item>
            <Typography variant="h5">Answers and options</Typography>
          </Grid>
          <Grid item>
            <Card>
              <CardHeader title={'Available Answers'} />
              <Divider />
              <CardContent>
                <RadioGroup
                  aria-label="is correct"
                  name="course type"
                  value={getCorrectAnswer()}
                  onChange={(evt, value) => {
                    const newAns = [...state.answers].map(ans => {
                      ans.isCorrect = false;
                      return ans;
                    });

                    newAns[value].isCorrect = true;
                    setState('answers', newAns);
                  }}
                >
                  <Grid container spacing={3} direction={'column'}>
                    {state.answers.map((answer, index) => (
                      <Grid item key={index}>
                        <div className={classes.answerItem}>
                          {answer.imageURL && (
                            <img
                              src={answer.imageURL}
                              className={classes.previewImage}
                            />
                          )}
                          <Typography variant="h6">{answer.text}</Typography>
                          <FormControlLabel value={index} control={<Radio />} />
                        </div>
                      </Grid>
                    ))}

                    <Grid item>
                      <Button
                        variant="contained"
                        color="primary"
                        onClick={() => setShowAnswerInput(true)}
                      >
                        Add Answer
                      </Button>
                    </Grid>
                  </Grid>
                </RadioGroup>
              </CardContent>
            </Card>
          </Grid>
          {showAnswerInput && (
            <Grid item>
              <AnswerInput
                answer={answer}
                onSave={onSaveAnswer}
                onChange={ans => {
                  console.log('ans', ans);
                  setAnswer(ans);
                }}
              />
            </Grid>
          )}
        </Grid>
      </Grid>
    </div>
  );
}

export default Overview;
