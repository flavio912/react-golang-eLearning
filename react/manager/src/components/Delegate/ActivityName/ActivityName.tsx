import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Text from "components/core/Table/Text/Text";
import themeConfig, { Theme } from "helpers/theme";
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
  },
  userName: {
    color: theme.colors.textNavyBlue,
    fontSize: theme.fontSizes.tiny,
    marginRight: 16,
    fontWeight: "bold",
    width: 29,
    height: 29,
    background: theme.colors.textSolitude,
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    borderRadius: 30,
    letterSpacing: -0.28,
  },
}));

type Props = {
  title: string;
  userName: string;
};
function ActiveName({ userName, title }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const nameArr = userName.split(" ");
  const shortName = nameArr.map((item: string) => item.charAt(0));
  return (
    <div className={classes.root}>
      <span className={classes.userName}>{shortName}</span>
      <Text text={title} color={themeConfig.colors.secondaryBlack}></Text>
    </div>
  );
}

export default ActiveName;
