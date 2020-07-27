import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    background: theme.primaryGradient,
    display: 'inline-block',
    borderRadius: theme.buttonBorderRadius
  },
  inner: {
    background: '#f3fcf4',
    margin: 1,
    borderRadius: theme.buttonBorderRadius - 1,
    display: 'flex',
    padding: 4
  },
  ttcTag: {
    background: '#f3fcf4',
    fontWeight: 600,
    fontSize: 9,
    padding: '2px 5px',
    borderRadius: 3,
    justifyContent: 'center',
    alignItems: 'center',
    display: 'flex',
    boxShadow: '0px 1px 4px #0000001f',
    border: [1, 'solid', '#81BE86']
  },
  identifier: {
    padding: [0, 9],
    justifyContent: 'center',
    alignItems: 'center',
    display: 'flex',
    fontSize: 11,
    fontWeight: 400
  }
}));

type Props = {
  ident: string;
  label?: string;
};

function IdentTag({ ident, label = 'TTC-ID' }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.root}>
      <div className={classes.inner}>
        <div className={classes.ttcTag}>{label}</div>
        <div className={classes.identifier}>{ident}</div>
      </div>
    </div>
  );
}

export default IdentTag;
