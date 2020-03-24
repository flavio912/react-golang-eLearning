import React from 'react';
import clsx from 'clsx';
import moment from 'moment';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/styles';
import {
  Grid,
  Hidden,
  Typography,
  ButtonGroup,
  Button
} from '@material-ui/core';

const useStyles = makeStyles(() => ({
  root: {}
}));
function Toolbar({
  date,
  view,
  onDatePrev,
  onDateNext,
  onEventAdd,
  onDateToday,
  className,
  ...rest
}) {
  const classes = useStyles();

  return (
    <div {...rest} className={clsx(classes.root, className)}>
      <Grid alignItems="flex-end" container justify="space-between" spacing={3}>
        <Grid item>
          <Typography component="h2" gutterBottom variant="overline">
            Dashboard
          </Typography>
          <Typography component="h1" variant="h3">
            Classroom courses
          </Typography>
        </Grid>
      </Grid>
      <Grid alignItems="center" container justify="space-between" spacing={3}>
        <Grid item>
          <ButtonGroup>
            <Button onClick={onDatePrev}>Prev</Button>
            <Button onClick={onDateToday}>Today</Button>
            <Button onClick={onDateNext}>Next</Button>
          </ButtonGroup>
        </Grid>
        <Hidden smDown>
          <Grid item>
            <Typography variant="h3">
              {moment(date).format('MMMM YYYY')}
            </Typography>
          </Grid>
          <Grid item />
        </Hidden>
      </Grid>
    </div>
  );
}

Toolbar.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
  date: PropTypes.any.isRequired,
  onDateNext: PropTypes.func,
  onDatePrev: PropTypes.func,
  onDateToday: PropTypes.func,
  onEventAdd: PropTypes.func,
  view: PropTypes.string.isRequired
};

export default Toolbar;
