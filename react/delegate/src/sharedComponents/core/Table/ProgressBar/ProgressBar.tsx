import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import "react-step-progress-bar/styles.css";
//@ts-ignore
import { ProgressBar as Progress } from "react-step-progress-bar";

import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  progressBarRoot: {
    display: "flex",
    alignItems: "center",
  },
  progressSlider: {},
}));

type Props = {
  percent: number;
  width?: number;
};

function ProgressBar({ percent, width = 200 }: Props) {
  const theme: any = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.progressBarRoot}>
      <Progress
        height={7}
        width={width}
        percent={percent}
        filledBackground={theme.primaryGradient}
        unfilledBackground={theme.colors.progressGrey}
      />
    </div>
  );
}

export default ProgressBar;
