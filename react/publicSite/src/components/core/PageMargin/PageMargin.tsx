import * as React from 'react';
import classNames from 'classnames';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  centerer: {
    display: 'flex',
    justifyContent: 'center',
    padding: '0 20px',
  },
  centered: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    width: theme.centerColumnWidth,
    padding: 'auto',
    position: 'relative',
  },
}));

type Props = {
  centererStyle?: string;
  centeredStyle?: string;
  children?: React.ReactNode;
};

function PageMargin({ centererStyle, centeredStyle, children }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classNames(classes.centerer, centererStyle)}>
      <div className={classNames(classes.centered, centeredStyle)}>
        {children}
      </div>
    </div>
  );
}

export default PageMargin;
