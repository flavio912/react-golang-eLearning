import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";
import QuickInfo from "../QuickInfo";
import Card from "sharedComponents/core/Card";
import { IconNames } from "sharedComponents/core/Icon/Icon";
import Spacer from "components/core/Spacers/Spacer";
const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    color: theme.colors.textBlue,
    padding: "20px 33px",
  },
  infoItems: {
    display: "flex",
    justifyContent: "space-between",
    marginTop: theme.spacing(2),
  },
  progressTitle: {
    fontWeight: 300,
  },
}));

type Props = {
  coursesDone: number;
  timeTracked: { h: number; m: number } | string;
  title: string;
  coursesPercent?: number;
  timePercent?: number;
  courseTimeTrackedIcon?: IconNames;
  courseNewCourseIcon?: IconNames;
  courseTitle?: string;
};
function TrainingProgressCard({
  coursesDone,
  timeTracked,
  title,
  coursesPercent,
  timePercent,
  courseTitle = "Courses done",
  courseTimeTrackedIcon = "CourseTimeTracked",
  courseNewCourseIcon = "CourseNewCourseGreen",
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <Card padding="medium" className={classes.root}>
      <div className={classes.progressTitle}>{title}</div>
      <div className={classes.infoItems}>
        <QuickInfo
          icon={courseNewCourseIcon as IconNames}
          text={courseTitle}
          value={coursesDone}
          percentValue={coursesPercent}
        />
        <Spacer horizontal spacing={3} />
        <QuickInfo
          icon={courseTimeTrackedIcon as IconNames}
          text="Time Tracked"
          value={timeTracked}
          percentValue={timePercent}
        />
      </div>
    </Card>
  );
}

export default TrainingProgressCard;
