import React from 'react';

import { Theme } from 'helpers/theme';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';

const useStyles = createUseStyles((theme: Theme) => ({
  pageRoot: {
    width: '100%',
    display: 'flex',
    position: 'relative',
    justifyContent: 'center'
  },
  pagePadding: {
    padding: '42px 60px'
  }
}));

type Props = {
  children?: React.ReactNode;
  noPadding?: boolean;
};

function Page({ children, noPadding }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div
      className={classNames(
        classes.pageRoot,
        noPadding ? '' : classes.pagePadding
      )}
    >
      {children}
    </div>
  );
}

export default Page;
