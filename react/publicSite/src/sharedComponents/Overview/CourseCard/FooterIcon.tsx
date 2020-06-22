import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import moment from 'moment';
import Icon, { IconNames } from "../../core/Icon";
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center'
  },
  icon: {
    display: 'flex',
    margin: '10px'
  },
  iconValue: {
    display: 'flex',
    fontSize:  theme.fontSizes.large,
    fontWeight: 900
  },
}));

type Props = {
  name: IconNames;
  size: number;
  text?: string;
  value?: number;
  date?: Date;
  className?: string;
};

function FooterIcon({ name, size, text, value, date, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
        <Icon className={classNames(classes.icon)} name={name} size={size} />
        <div className={classNames(classes.iconValue)}>
            {text && text}
            {value && value}
            {date && moment(date).calendar()}
        </div>
    </div>
    
  );
}

export default FooterIcon;
