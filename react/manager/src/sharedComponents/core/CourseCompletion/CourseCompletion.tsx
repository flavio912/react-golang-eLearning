import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';

import { Theme } from 'helpers/theme';
import ProgressBar from '../ProgressBar';

const useStyles = createUseStyles((theme: Theme) => ({
  CourseCompletionRoot: {
    display: 'flex',
    alignItems: 'center'
  },
  percentText: {
    fontSize: theme.fontSizes.tiny,
    fontWeight: 300,
    color: theme.colors.textBlue,
    marginRight: theme.spacing(1)
  },
  fractionHolder: {
    marginLeft: theme.spacing(1),
    fontSize: theme.fontSizes.tiny,
    color: theme.colors.textBlue,
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center'
  },
  numerator: {},
  divider: {
    height: 1,
    minWidth: 20,
    background: theme.colors.primaryGreen,
    width: '100%'
  },
  denominator: {}
}));

type Props = {
  complete: number;
  total: number;
  width?: number;
  fraction?: boolean;
};

function CourseCompletion({
  complete,
  total,
  width = 200,
  fraction = true
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  var percent = 0;
  if (total > 0) {
    percent = Math.round((complete / total) * 100);
  }

  return (
    <div className={classes.CourseCompletionRoot}>
      <div className={classes.percentText}>{percent}%</div>
      <ProgressBar percent={percent} width={width} />
      {fraction && (
        <div className={classes.fractionHolder}>
          <div className={classes.numerator}>{complete}</div>
          <div className={classes.divider} />
          <div className={classes.denominator}>{total}</div>
        </div>
      )}
    </div>
  );
}

export default CourseCompletion;
