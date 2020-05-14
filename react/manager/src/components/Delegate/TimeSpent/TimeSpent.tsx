import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Text from "components/core/Table/Text/Text";
import themeConfig, { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    "& div": {
      marginLeft: 7,
    },
  },
}));

type Props = {
  h: string | number;
  m: string | number;
};

function TimeSpent({ h, m }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <span className={classes.root}>
      <Icon name={"CourseStatus_Completed"} size={25} />
      <Text
        text={`${h}hr ${m}mins`}
        color={themeConfig.colors.secondaryBlack}
      />
    </span>
  );
}

export default TimeSpent;
