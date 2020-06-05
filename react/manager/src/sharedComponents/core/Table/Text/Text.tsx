import * as React from "react";
import { createUseStyles } from "react-jss";
import moment from "moment";

import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  textRoot: {
    display: "flex",
    alignItems: "center",
    fontSize: theme.fontSizes.small,
    fontWeight: 300,
  },
}));

type Props = {
  text: string | Date;
  color?: string;
  formatDate?: boolean; // If given formats the text as a date to "DD-MM-YY" format using moment
};

function Text({ text, color, formatDate }: Props) {
  const classes = useStyles();

  if (formatDate) {
    text = moment(text).format("DD-MM-YY");
  }

  const extraStyles: any = {};
  if (color) {
    extraStyles.color = color;
  }

  return (
    <div className={classes.textRoot} style={extraStyles}>
      {text}
    </div>
  );
}

export default Text;
