import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  card: {
    boxShadow: theme.shadows.primary,
    borderRadius: theme.primaryBorderRadius,
    background: 'white'
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
  children: React.ReactNode;
  padding?: PaddingOptions;
  className?: string;
};

function Card({
  children,
  padding = 'none',
  className,
  ...props
}: Props & React.PropsWithoutRef<JSX.IntrinsicElements['div']>) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const paddingLink = {
    none: '',
    small: classes.smallPadding,
    medium: classes.mediumPadding,
    large: classes.largePadding
  };

  return (
    <div
      className={classNames(classes.card, paddingLink[padding], className)}
      {...props}
    >
      {children}
    </div>
  );
}

export default Card;
