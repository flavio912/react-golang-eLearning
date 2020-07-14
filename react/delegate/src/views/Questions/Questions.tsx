import React, { useState } from 'react';
import { createUseStyles, useTheme } from 'react-jss';
// import { Grid } from '@material-ui/core';
import { Theme } from 'helpers/theme';
import Heading from 'components/core/Heading';
import Question from 'components/Misc/Question';
import Button from 'components/core/Input/Button';

const useStyles = createUseStyles((theme: Theme) => ({
  questionsRoot: {
    display: 'flex',
    flexGrow: 1,
    maxWidth: 1275
  },
  viewPanel: {
    flex: 2
  },
  infoPanel: {
    flex: 1,
    marginLeft: 30,
    borderTop: '1px solid',
    paddingTop: 30
  },
  mainHeading: {
    gridArea: 'headin'
  },
  subHeading: {
    gridArea: 'subhea',
    maxWidth: 466
  },
  nextQuestionWrap: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: 150,
    background: 'white',
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: 6,
  },
  infoHeading: {
    color: `${theme.colors.primaryGreen}`
  },
  questionStep: {
    margin: 0,
    letterSpacing: 0
  }
}));

type Props = {
  className?: string;  
};

const QUESTIONS = [
  {
    id: 0,
    type: "text",
    head: "Module 1 - General Philosophy Test Question 1",
    title: "What is the authority with responsibility for the transport of dangerous goods by air in the UK?",
    options: [
      {
        id: 1,
        title: "The Department for Transport (DfT)",
        index: "A"
      },
      {
        id: 2,
        title: "The Ministry of Justice (MoJ)",
        index: "B"
      },
    ],
  },
  {
    id: 1,
    type: "image",
    head: "Module 2 - Dangerous Goods Test Question 2",
    title: "Which of the following items is NOT classified as suspcious or dangerous cargo",
    options: [
      {
        id: 1,
        image: "https://media.istockphoto.com/photos/chef-knife-picture-id874095794"
      },
      {
        id: 2,
        image: "https://media.istockphoto.com/photos/traditional-chefs-knife-isolated-on-a-white-background-picture-id832724072"
      }
    ]
  }
];

function Questions({ className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [curQuestion, setCurQuestion] = useState(QUESTIONS[0]);

  const handleNextQuestion = () => {
    if (curQuestion.id >= QUESTIONS.length - 1) return;
    setCurQuestion(QUESTIONS[curQuestion.id + 1]);
  }
  
  return (
    <div className={classes.questionsRoot}>
      <div className={classes.viewPanel}>
        <Heading
          text={curQuestion.head}
          size={'large'}
          className={classes.mainHeading}
        />
        {curQuestion.type === "text"?
            <Question question={curQuestion} type="text" onSelected={() => {}} />
            :
            <Question question={curQuestion} type="image" onSelected={() => {}} />
        }
        <div className={classes.nextQuestionWrap}>      
          <Button
            title={"Next Question"}
            padding="large"
            onClick={handleNextQuestion}
          />
        </div>
      </div>
      <div className={classes.infoPanel}>
        <span className={classes.infoHeading}>Question</span>
        <h1 className={classes.questionStep}>{curQuestion.id+1} / {QUESTIONS.length}</h1>
      </div>
    </div>
  )
}

export default Questions;