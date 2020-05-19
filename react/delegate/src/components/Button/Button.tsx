import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import ButtonBase from "sharedComponents/core/Button";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    backgroundColor: theme.colors.navyBlue,
    borderRadius: 6,
    borderColor: "transparent",
    height: 59,
    width: 230,
    cursor: 'pointer'
  },
  buttonText: {
    fontSize: 19.5,
    fontWeight: "bold",
    color: theme.colors.primaryWhite,
    letterSpacing: -0.49,
  },
}));

type Props = {
  title: string;
  onClick: Function;
  className?: string;
};

function Button({ title, onClick, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <ButtonBase
      onClick={() => onClick()}
      className={classNames(classes.root, className)}
    >
      <span className={classes.buttonText}>{title}</span>
    </ButtonBase>
  );
}

export default Button;
