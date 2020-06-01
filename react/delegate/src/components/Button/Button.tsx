import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import ButtonBase from 'sharedComponents/core/Button';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    backgroundColor: theme.colors.navyBlue,
    borderRadius: 6,
    borderColor: 'transparent',
    cursor: 'pointer',
    height: 'auto',
    whiteSpace: 'nowrap'
  },
  buttonText: {
    fontSize: 19.5,
    fontWeight: 'bold',
    color: theme.colors.primaryWhite,
    letterSpacing: -0.49
  },
  smallPadding: {
    padding: '5px 20px'
  },
  mediumPadding: {
    padding: '10px 30px'
  },
  largePadding: {
    padding: '15px 35px'
  },
  massivePadding: {
    padding: '17px 39px'
  }
}));
export type PaddingOptions = 'none' | 'small' | 'medium' | 'large' | 'massive';

export type ButtonProps = {
  title: string;
  onClick: Function;
  className?: string;
  padding?: PaddingOptions;
};

function Button({ title, onClick, className, padding = 'none' }: ButtonProps) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const paddingLink = {
    none: '',
    small: classes.smallPadding,
    medium: classes.mediumPadding,
    large: classes.largePadding,
    massive: classes.massivePadding
  };
  return (
    <ButtonBase
      onClick={() => onClick()}
      className={classNames(classes.root, paddingLink[padding], className)}
    >
      <span className={classes.buttonText}>{title}</span>
    </ButtonBase>
  );
}

export default Button;
