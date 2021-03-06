import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";
import QuickInfo from "../QuickInfo";
import Card from "sharedComponents/core/Cards/Card";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    color: theme.colors.textBlue,
    minWidth: 278,
  },
  infoItems: {
    display: "flex",
    justifyContent: "space-between",
    marginTop: theme.spacing(2),
  },
}));

type Props = {
  coursesDone: number;
  timeTracked: { h: number; m: number };
  title: string;
};

function TrainingProgressCard({ coursesDone, timeTracked, title }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <Card padding="medium" className={classes.root}>
      <div>{title}</div>
      <div className={classes.infoItems}>
        <QuickInfo
          icon="CourseNewCourseGreen"
          text="Courses done"
          value={coursesDone}
        />
        <QuickInfo
          icon="CourseTimeTracked"
          text="Time Tracked"
          value={timeTracked}
        />
      </div>
    </Card>
  );
}

export default TrainingProgressCard;
