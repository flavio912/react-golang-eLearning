import * as React from "react";
import { createUseStyles } from "react-jss";
import Icon, { IconNames } from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: "flex",
    flexDirection: "column",
  },
  text: {
    color: theme.colors.textBlue,
    fontSize: theme.fontSizes.small,
    margin: [6, 0],
    fontWeight: 300,
  },
  value: {
    fontWeight: 800,
    color: "black",
    fontSize: 20,
    margin: 0,
    "& span": {
      fontWeight: 400,
    },
  },
}));

type Props = {
  icon: IconNames;
  text: string;
  value: number | { h: number; m: number };
};

function QuickInfo({ icon, text, value }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <Icon name={icon} size={25} />
      <p className={classes.text}>{text}</p>
      {typeof value === "number" ? (
        <p className={classes.value}>{value}</p>
      ) : (
        <p className={classes.value}>
          {value.h}
          <span>h </span>
          {value.m}
          <span>m</span>
        </p>
      )}
    </div>
  );
}

export default QuickInfo;
