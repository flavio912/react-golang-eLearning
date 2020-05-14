import * as React from "react";
import { createUseStyles } from "react-jss";
import Icon from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  actionsRoot: {
    borderLeft: "1px solid #ededed",
    paddingLeft: 23,
    height: 38,
    display: "flex",
    alignItems: "center",
  },
}));

type Props = {};

function Action({}: Props) {
  const classes = useStyles();

  return (
    <div className={classes.actionsRoot}>
      <Icon name={"Card_SecondaryActon_Dots"} />
    </div>
  );
}

export default Action;
