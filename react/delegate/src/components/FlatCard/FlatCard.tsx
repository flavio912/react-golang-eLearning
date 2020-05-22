import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import { boolean } from '@storybook/addon-knobs';

const useStyles = createUseStyles((theme: Theme) => ({
  flatCardRoot: {
    backgroundColor: theme.colors.primaryWhite,
    borderRadius: 6,
    border: `1px solid ${theme.colors.borderGrey}`,
    height: 170,
    width: '100%',
    display: 'flex',
    boxSizing: 'border-box',
    alignItems: 'center',
    justifyContent: 'center',
    boxShadow: ({ shadow }: { shadow: boolean }) =>
      shadow ? '2px 2px 10px 0 rgba(0,0,0,0.07)' : ''
  }
}));

type Props = {
  className?: string;
  children: React.ReactNode;
  shadow: boolean;
  style?: any;
  backgroundColor?: string;
};

function FlatCard({
  className,
  style,
  children,
  backgroundColor,
  shadow
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ shadow, theme });
  style = {
    ...style,
    backgroundColor
  };
  return (
    <div className={classNames(classes.flatCardRoot, className)} style={style}>
      {children}
    </div>
  );
}

export default FlatCard;
