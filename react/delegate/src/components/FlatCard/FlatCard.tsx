import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    backgroundColor: theme.colors.primaryWhite,
    borderRadius: 6,
    border: `1px solid #EDEDED`,
    height: 170,
    width: 682,
    display: "flex",
    boxSizing: "border-box",
    alignItems: "center",
    justifyContent: "center",
  },
}));

type Props = {
  className?: string;
  children: React.ReactNode;
  style?: any;
  backgroundColor?: string;
};

function FlatCard({ className, style, children, backgroundColor }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  style = {
    ...style,
    backgroundColor,
  };
  return (
    <div className={classNames(classes.root, className)} style={style}>
      {children}
    </div>
  );
}

export default FlatCard;
