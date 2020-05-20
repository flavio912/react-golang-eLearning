import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  spacerRoot: {},
}));

type Props = {
  spacing: number;
  vertical?: boolean;
  horizontal?: boolean;
};

function Spacer({ spacing, vertical, horizontal }: Props) {
  //@ts-ignore
  const theme: Theme = useTheme();
  const classes = useStyles({ theme });

  const addedStyle: any = {};

  if (vertical && horizontal) {
    throw new Error("Cannot have both vertical and horizontal spacing");
  }

  if (theme && theme.spacing) {
    if (vertical) {
      addedStyle.height = theme.spacing(spacing);
    }
    if (horizontal) {
      addedStyle.width = theme.spacing(spacing);
    }
  }
  return <div className={classes.spacerRoot} style={addedStyle}></div>;
}

export default Spacer;
