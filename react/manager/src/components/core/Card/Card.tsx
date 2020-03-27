import * as React from 'react';
import PropTypes from 'prop-types';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';

const useStyles = createUseStyles((theme: any) => ({
  card: {
    boxShadow: '0px 2px 10px #0000001f',
    borderRadius: theme.borderRadius
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

function Card({ children, padding = 'none', className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const paddingLink = {
    none: '',
    small: classes.smallPadding,
    medium: classes.mediumPadding,
    large: classes.largePadding
  };

  return (
    <div className={classNames(classes.card, paddingLink[padding], className)}>
      {children}
    </div>
  );
}

export default Card;
