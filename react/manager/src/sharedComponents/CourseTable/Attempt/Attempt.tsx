import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Text from 'sharedComponents/core/Table/Text/Text';
import themeConfig, { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    alignItems: 'center',
    '& div': {
      fontSize: theme.fontSizes.large
    }
  },
  stText: {
    fontSize: theme.fontSizes.tiny,
    marginBottom: 5
  }
}));

type Props = {
  attempt: string;
};

function Attempt({ attempt }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.root}>
      <Text text={attempt} color={themeConfig.colors.secondaryBlack} />
      <span className={classes.stText}>st</span>
    </div>
  );
}

export default Attempt;
