import * as React from 'react';
import { createUseStyles } from 'react-jss';
import moment from 'moment';

import { Theme } from 'helpers/theme';
import { format } from 'path';

const useStyles = createUseStyles((theme: Theme) => ({
  textRoot: {
    display: 'flex',
    alignItems: 'center',
    fontSize: theme.fontSizes.small,
    fontWeight: 300
  }
}));

type Props = {
  text: string | Date;
  color?: string;
  formatDate?: boolean; // If given formats the text as a date to "DD-MM-YY" format using moment
  invalidDateText?: string; // If given date is invalid or empty show this text instead
};

function Text({ text, color, formatDate, invalidDateText }: Props) {
  const classes = useStyles();

  if (formatDate) {
    console.log('text', text);
    text = moment(text).format('DD-MM-YY');
    if (!moment(text).isValid()) {
      text = invalidDateText ?? '-';
    }
  }

  const extraStyles: any = {};
  if (color) {
    extraStyles.color = color;
  }

  return (
    <div className={classes.textRoot} style={extraStyles}>
      {text}
    </div>
  );
}

export default Text;
