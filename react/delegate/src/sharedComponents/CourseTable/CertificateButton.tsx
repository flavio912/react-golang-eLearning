import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  button: {
    borderRadius: 3,
    fontWeight: 600,
    padding: '12px 18px'
  },
  active: {
    background: theme.colors.primaryBlue,
    border: 'none',
    color: 'white'
  },
  inactive: {
    background: theme.colors.searchHoverGrey,
    border: `1px solid #E9EBEB`,
    color: '#34373ab5'
  }
}));

type Props = {
  url?: string;
};

function CertificateButton({ url }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <button
      className={classnames(
        classes.button,
        url ? classes.active : classes.inactive
      )}
      onClick={() => {
        if (url) window.open(url);
      }}
    >
      {url ? 'Certificate' : 'Not Available'}
    </button>
  );
}

export default CertificateButton;
