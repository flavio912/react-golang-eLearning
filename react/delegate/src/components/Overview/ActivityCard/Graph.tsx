import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import {
  CircularProgressbarWithChildren,
  buildStyles,
} from "react-circular-progressbar";
import "react-circular-progressbar/dist/styles.css";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    width: "180px",
  },
  innerContainer: {
    width: "80%",
  },
  container: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
  },
  heading: {
    fontSize: theme.fontSizes.small,
    fontWeight: '300',
    color: theme.colors.textGrey,
  },
  statValue: {
    fontSize: theme.fontSizes.large,
    fontWeight: '700',
    color: theme.colors.primaryBlack,
  },
}));

const outerGraph = (theme: any) =>
  buildStyles({
    pathTransitionDuration: 1,
    pathColor: theme.colors.primaryGreen,
    trailColor: theme.colors.borderGrey,
  });

const innerGraph = (theme: any) =>
  buildStyles({
    rotation: 0.5,
    pathTransitionDuration: 1,
    pathColor: theme.colors.primaryRed,
    trailColor: theme.colors.borderGrey,
  });

type Props = {
  heading: string;
  innerValue: number;
  outerValue: number;
  className?: string;
};

function Graph({ heading, innerValue, outerValue, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [percentage, setPercentage] = React.useState(0);
  const [inner, setInner] = React.useState(0);
  const [outer, setOuter] = React.useState(0);

  setTimeout(() => {
    setPercentage((outerValue / (innerValue + outerValue)) * 100);
    setInner(innerValue);
    setOuter(outerValue);
  }, 100);

  return (
    <div className={classNames(classes.root, className)}>
      <CircularProgressbarWithChildren
        value={outer}
        strokeWidth={5}
        maxValue={innerValue + outerValue}
        styles={outerGraph(theme)}
      >
        <div className={classNames(classes.innerContainer)}>
          <CircularProgressbarWithChildren
            value={inner}
            strokeWidth={5}
            maxValue={innerValue + outerValue}
            styles={innerGraph(theme)}
          >
            <div className={classNames(classes.container)}>
              <div className={classNames(classes.statValue)}>
                {percentage.toFixed(0)}%
              </div>
              <div className={classNames(classes.heading)}>{heading}</div>
            </div>
          </CircularProgressbarWithChildren>
        </div>
      </CircularProgressbarWithChildren>
    </div>
  );
}

export default Graph;
