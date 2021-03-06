import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'flex-start'
  },
  fieldName: {
    flex: 1,
    fontSize: theme.fontSizes.small,
    color: theme.colors.textBlue
  },
  value: {
    flex: 2,
    fontSize: theme.fontSizes.large,
    color: theme.colors.textGrey,
  },
  smallPadding: {
    padding: '15px 0px'
  },
  mediumPadding: {
    padding: '26px 0px'
  },
  largePadding: {
    padding: '30px 0px'
  },
  border: {
    borderBottom: `1px solid ${theme.colors.borderGrey}`
  }
}));

export type PaddingOptions = 'none' | 'small' | 'medium' | 'large';
export interface Field {
  fieldName: string;
  value: string;
  padding?: PaddingOptions;
};

type Props = {
  fieldName: string;
  value: string;
  padding?: PaddingOptions;
  border?: boolean;
  className?: string;
};

function InfoField({ fieldName, value, padding = 'medium', border = true, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const paddingLink = {
    none: '',
    small: classes.smallPadding,
    medium: classes.mediumPadding,
    large: classes.largePadding
  };

  return (
    <div className={classNames(classes.container, paddingLink[padding], border && classes.border, className)}>
      <span className={classNames(classes.fieldName)}>{fieldName}</span>
      <span className={classNames(classes.value)}>{value}</span>
    </div>
  );
}

export default InfoField;
