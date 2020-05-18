import * as React from "react";
import { createUseStyles } from "react-jss";
import Icon, { IconNames } from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";
import classnames from "classnames";
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
    fontSize: 18,
    margin: 0,
    "& span": {
      fontWeight: 400,
    },
  },
  iconContainer: {
    display: "flex",
    alignItems: "center",
  },
  percentValue: {
    color: theme.colors.secondaryGreen,
    lineHeight: 11,
    fontWeight: 700,
    fontSize: 9,
    display: "flex",
    alignItems: "center",
    borderRadius: 10,
    width: 43,
    height: 16,
    justifyContent: "center",
    marginLeft: 6,
    letterSpacing: -0.23,
  },
  percentSuccess: {
    color: theme.colors.secondaryGreen,
    backgroundColor: theme.colors.hoverGreen,
  },
  percentError: {
    color: theme.colors.secondaryDanger,
    backgroundColor: theme.colors.hoverDanger,
  },
}));

type Props = {
  icon: IconNames;
  text: string;
  value: number | string | { h: number; m: number };
  percentValue?: number;
};

function QuickInfo({ icon, text, value, percentValue }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <div className={classes.iconContainer}>
        <Icon name={icon} size={25} />
        {percentValue && (
          <span
            className={classnames(classes.percentValue, {
              [classes.percentSuccess]: percentValue >= 0,
              [classes.percentError]: percentValue < 0,
            })}
          >
            {percentValue >= 0 ? "+" : ""}
            {percentValue}%
          </span>
        )}
      </div>
      <p className={classes.text}>{text}</p>
      {typeof value === "number" || typeof value === "string" ? (
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
