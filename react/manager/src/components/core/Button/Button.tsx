import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  button: {
    borderRadius: theme.primaryBorderRadius,
    padding: [theme.spacing(1), theme.spacing(2)],
    border: "1px solid",
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    cursor: "pointer",
    height: 40,
    transition: "0.1s ease",
    transitionProperty: "border-colour, background-color",
    outline: "none",
    // it would be nice to have a light blue hover state
    "&:focus": {
      borderColor: theme.colors.primaryBlue,
    },
    "&::-moz-focus-inner": {
      border: 0,
    },
  },
  bold: {
    fontWeight: "bold",
  },
  small: {
    padding: theme.spacing(1),
    "& $iconLeft": {
      // replace with real icon
      margin: [0, theme.spacing(0), 0, 0],
      colour: "red",
    },
    "& $iconRight": {
      // replace with real icon
      margin: [0, 0, 0, theme.spacing(0)],
    },
  },
  default: {
    color: theme.colors.primaryBlack,
    borderColor: theme.colors.borderGrey,
    backgroundColor: "white",
  },
  grey: {
    borderWidth: 2,
    color: theme.colors.primaryBlack,
    borderColor: theme.colors.borderGrey,
    backgroundColor: theme.colors.backgroundGrey,
  },
  submit: {
    color: "white",
    borderColor: theme.colors.primaryBlue,
    backgroundColor: theme.colors.primaryBlue,
    "&:focus": {
      borderColor: "#0044db",
      backgroundColor: "#0044db",
    },
  },
  iconLeft: {
    // replace with real icon
    margin: [0, theme.spacing(2), 0, 0],
  },
  iconRight: {
    // replace with real icon
    margin: [0, 0, 0, theme.spacing(2)],
  },
}));

export type Archetypes = "default" | "grey" | "submit";

interface Props {
  archetype?: Archetypes;
  iconLeft?: boolean;
  iconRight?: boolean;
  bold?: boolean;
  small?: boolean;
}

function Button({
  archetype,
  iconLeft,
  iconRight,
  bold,
  small,
  children,
  ...props
}: Props & React.ButtonHTMLAttributes<HTMLButtonElement>) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <button
      className={classNames(
        classes.button,
        classes[archetype || "default"],
        bold && classes.bold,
        small && classes.small
      )}
      {...props}
    >
      {/* replace with actual icon */}
      {/* prop should also be a string (icon name) */}
      {iconLeft && <p className={classNames(classes.iconLeft)}>+</p>}
      {children}
      {iconRight && <p className={classNames(classes.iconRight)}>></p>}
    </button>
  );
}

export default Button;
