import * as React from 'react';
import { createUseStyles } from 'react-jss';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  demoText: {
    fontSize: theme.fontSizes.default,
    color: theme.colors.textGrey
  },
  blueText: {
    color: theme.colors.navyBlue
  }
}));

type Props = {
  children?: React.ReactNode;
  onBook: () => void;
};

function Footer({ children, onBook }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.demoText} onClick={onBook}>
      Or <span className={classes.blueText}>Book A Demo</span> today for more
      information
    </div>
  );
}

export default Footer;
