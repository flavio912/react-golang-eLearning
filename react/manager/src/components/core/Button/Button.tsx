import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Icon, { IconNames } from "../Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  button: {
    borderRadius: theme.buttonBorderRadius,
    padding: [0, theme.spacing(2)],
    border: "1px solid",
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    cursor: "pointer",
    height: 40,
    fontSize: theme.fontSizes.default,
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
    fontWeight: "bold !important",
  },
  small: {
    padding: [0, theme.spacing(1)],
  },
  default: {
    color: theme.colors.primaryBlack,
    borderColor: theme.colors.borderGrey,
    backgroundColor: "white",
    fontWeight: 200,
    borderRadius: 4,
  },
  grey: {
    borderWidth: 1,
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
}));

export type Archetypes = "default" | "grey" | "submit";

interface Props {
  archetype?: Archetypes;
  iconLeft?: IconNames;
  iconRight?: IconNames;
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
      {iconLeft && (
        <Icon
          style={{ marginRight: small ? 5 : 10 }}
          name={iconLeft}
          size={small ? 12 : 15}
        />
      )}
      {children}
      {iconRight && (
        <Icon
          style={{ marginLeft: small ? 5 : 10 }}
          name={iconRight}
          size={small ? 12 : 15}
        />
      )}
    </button>
  );
}

export default Button;
