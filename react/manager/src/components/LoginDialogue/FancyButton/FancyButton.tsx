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
    background: theme.primaryGradient,
    transition: "margin 0.1s ease",
    "&::-moz-focus-inner": {
      border: 0,
    },
    cursor: "pointer",
  },
}));

type Props = {
  text: string;
};

function FancyButton({
  text,
  ...props
}: Props & React.PropsWithoutRef<JSX.IntrinsicElements["input"]>) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <input {...props} type="submit" value={text} className={classes.submit} />
    </div>
  );
}

export default FancyButton;
