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
  Button,
  CircularProgress
} from '@material-ui/core';
import { makeStyles } from '@material-ui/styles';
import { gql } from 'apollo-boost';
import { useMutation, useQuery } from '@apollo/react-hooks';

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
  input: {
    display: 'none'
  }
}));
