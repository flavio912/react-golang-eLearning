import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import Text from 'sharedComponents/core/Table/Text/Text';
import themeConfig, { Theme } from 'helpers/theme';
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex'
  },
  userName: {
    marginRight: 16,
    width: 33,
    height: 33,
    background: theme.colors.textSolitude,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 30,
    backgroundPosition: 'center',
    backgroundRepeat: 'no-repeat'
  }
}));

type Props = {
  title: string;
  avatar: string;
};
function ActiveName({ avatar, title }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.root}>
      <span
        className={classes.userName}
        style={{ backgroundImage: `url(${avatar})` }}
      ></span>
      <Text text={title} color={themeConfig.colors.secondaryBlack}></Text>
    </div>
  );
}

export default ActiveName;
