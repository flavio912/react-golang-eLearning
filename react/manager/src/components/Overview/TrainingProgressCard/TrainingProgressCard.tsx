import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";
import QuickInfo from "../QuickInfo";
import Card from "sharedComponents/core/Card";
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
  timeTracked: { h: number; m: number };
  title: string;
};

function TrainingProgressCard({ coursesDone, timeTracked, title }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <Card padding="medium" className={classes.root}>
      <div className={classes.progressTitle}>{title}</div>
      <div className={classes.infoItems}>
        <QuickInfo
          icon="CourseNewCourseGreen"
          text="Courses done"
          value={coursesDone}
        />
        <Spacer horizontal spacing={3} />
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
