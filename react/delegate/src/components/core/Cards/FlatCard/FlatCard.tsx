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
    width: '100%',
    display: 'flex',
    boxShadow: ({ shadow }: { shadow: boolean }) =>
      shadow ? '2px 2px 10px 0 rgba(0,0,0,0.07)' : ''
  },
  smallPadding: {
    padding: '15px 20px'
  },
  mediumPadding: {
    padding: '20px 25px'
  },
  largePadding: {
    padding: '30px 35px'
  }
}));

export type PaddingOptions = 'none' | 'small' | 'medium' | 'large';

type Props = {
  className?: string;
  children: React.ReactNode;
  shadow: boolean;
  style?: any;
  padding?: PaddingOptions;
  backgroundColor?: string;
};

function FlatCard({
  className,
  style,
  children,
  backgroundColor,
  padding = 'none',
  shadow
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ shadow, theme });

  const paddingLink = {
    none: '',
    small: classes.smallPadding,
    medium: classes.mediumPadding,
    large: classes.largePadding
  };

  style = {
    ...style,
    backgroundColor
  };
  return (
    <div
      className={classNames(
        classes.flatCardRoot,
        paddingLink[padding],
        className
      )}
      style={style}
    >
      {children}
    </div>
  );
}

export default FlatCard;
