import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";

const useStyles = createUseStyles((theme: Theme) => ({
  titleWrapperRoot: {
    display: "flex",
    flexDirection: "column",
  },
  titleWrapperTitle: {
    marginBottom: theme.spacing(2),
    fontWeight: 700,
    fontSize: theme.fontSizes.default,
    color: theme.colors.primaryBlack,
  },
}));

type Props = {
  title: string;
  children: React.ReactNode;
  className?: string;
};

function Card({ title, children, className = "" }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.titleWrapperRoot, className)}>
      <div className={classes.titleWrapperTitle}>{title}</div>
      {children}
    </div>
  );
}

export default Card;
