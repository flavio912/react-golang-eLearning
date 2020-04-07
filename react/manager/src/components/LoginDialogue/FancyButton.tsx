import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    backgroundColor: "black",
    borderRadius: theme.buttonBorderRadius,
    margin: [20, 0, 0, 0],
    padding: [0, 0, 3, 0],
    overflow: "hidden",
    transition: "0.1s ease",
    transitionProperty: "margin padding",
    "&:hover": {
      margin: [23, 0, 0, 0],
      padding: 0,
    },
  },
  submit: {
    border: "none",
    borderRadius: [0, 0, theme.buttonBorderRadius, theme.buttonBorderRadius],
    width: "100%",
    height: 50,
    color: "white",
    fontSize: 20,
    fontWeight: 800,
    background: `linear-gradient(to right, ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`,
    transition: "margin 0.1s ease",
    "&::-moz-focus-inner": {
      border: 0,
    },
  },
}));

type Props = {
  text: string;
};

function FancyButton({ text }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <input type="submit" value={text} className={classes.submit} />
    </div>
  );
}

export default FancyButton;
