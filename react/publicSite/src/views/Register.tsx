import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classnames from 'classnames';
import { Theme } from 'helpers/theme';
import FloatingVideo from 'components/core/VideoPlayer/FloatingVideo';
import ImageWithText from 'components/core/ImageWithText';
import Spacer from 'sharedComponents/core/Spacers/Spacer';
import DetailsCard from 'components/Overview/Registration/DetailsCard';

const useStyles = createUseStyles((theme: Theme) => ({
  registerRoot: {
    width: '100%'
  },
  centerer: {
    display: 'flex',
    justifyContent: 'center'
  },
  centered: {
    width: theme.centerColumnWidth
  },
  whiteSpacer: {
    background: theme.colors.primaryWhite,
    padding: '60px 0px 100px 0px'
  },
  heading: {
    fontSize: 32,
    color: theme.colors.primaryBlack,
    fontWeight: 800,
    padding: '60px 0px',
    textAlign: 'center'
  },
  explore: {
    paddingTop: 70
  }
}));

type Props = {};

function Register({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.registerRoot}>
      <DetailsCard />
    </div>
  );
}

export default Register;
