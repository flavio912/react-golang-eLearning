import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classnames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  dashboardRoot: {
    position: "relative",
    display: "grid",
    gridTemplateColumns: "repeat(6, 1fr)",
    gridTemplateRows: "repeat(4, auto)",
    gridGap: theme.spacing(2),
  },
}));

type Props = {};

function Dashboard({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.dashboardRoot}>
      <h1>Dashboard</h1>
    </div>
  );
}

export default Dashboard;
