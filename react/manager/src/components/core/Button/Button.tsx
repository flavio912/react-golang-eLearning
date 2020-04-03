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
  },
  bold: {
    fontWeight: "bold",
  },
  smallPadding: {
    padding: theme.spacing(1),
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
  },
  iconLeft: {
    // replace with real icon
    margin: [0, theme.spacing(0), 0, 0],
  },
  iconRight: {
    // replace with real icon
    margin: [0, 0, 0, theme.spacing(0)],
  },
}));

export type Archetypes = "default" | "grey" | "submit";

interface Props {
  archetype?: Archetypes;
  iconLeft?: boolean;
  iconRight?: boolean;
  bold?: boolean;
}

function Button({
  archetype,
  iconLeft,
  iconRight,
  bold,
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
        (iconLeft || iconRight) && classes.smallPadding
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
