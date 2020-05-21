import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import CoreInput, { InputTypes } from "sharedComponents/core/CoreInput";
import classnames from "classnames";
import { Theme } from "helpers/theme";
import Heading from "components/Heading";

const useStyles = createUseStyles((theme: Theme) => ({
  dashboardRoot: {
    position: "relative",
    display: "grid",
    gridTemplateColumns: "repeat(6, 1fr)",
    gridTemplateRows: "repeat(4, auto)",
    gridGap: theme.spacing(2),
  },
  trainingHeader: {
    gridArea: "1 / 1 / 2 / 4",
  },
}));

type Props = {};

function Dashboard({}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const userName = "James";
  return (
    <div className={classes.dashboardRoot}>
      <div className={classes.trainingHeader}>
        <Heading text="Training Zone" size={"large"} />
        <Heading
          text={`You're doing great ${userName}, keep up the good work so you don't loose your momentum`}
          size={"medium"}
        />
      </div>
    </div>
  );
}

export default Dashboard;
