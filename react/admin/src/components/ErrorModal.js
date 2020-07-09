import React, { useState, useEffect } from 'react';
import {
  Grid,
  TextField,
  Card,
  CardHeader,
  CardContent,
  Typography
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';

const useStyles = makeStyles(theme => ({
  errorRoot: {
    position: 'fixed',
    bottom: 50,
    right: 50
  }
}));

function ErrorModal({ error }) {
  const [show, setShow] = useState(false);

  const classes = useStyles();

  useEffect(() => {
    setShow(!!error);
    setTimeout(() => {
      setShow(false);
    }, 7000);
  }, [error]);

  if (!show) return false;
  if (!error) return false;
  var message = error.message;
  if (error.graphQLErrors && error.graphQLErrors.length == 1) {
    message = error.graphQLErrors[0]?.extensions?.message;
  }

  return (
    <Card className={classes.errorRoot}>
      <CardContent>
        <Typography>{message}</Typography>
      </CardContent>
    </Card>
  );
}

export default ErrorModal;
