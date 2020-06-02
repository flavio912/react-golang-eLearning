import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import Text from "components/core/Table/Text/Text";
import Button from "sharedComponents/core/Input/Button";
import themeConfig, { Theme } from "helpers/theme";
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    "& button": {
      marginLeft: 51,
    },
    "& div": {
      fontSize: themeConfig.fontSizes.default,
    },
  },
}));

type Props = { title: string };
function ActiveCoursesEmpty({ title }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.root}>
      <Text text={title} color={themeConfig.colors.secondaryBlack} />
      <Button bold archetype="submit">
        Book now
      </Button>
    </div>
  );
}

export default ActiveCoursesEmpty;
