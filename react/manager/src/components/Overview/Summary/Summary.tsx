import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import QuickInfo from "../QuickInfo";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    justifyContent: "space-between",
  },
  spacer: {
    width: theme.spacing(3),
  },
  summaryRoot: {
    flex: 1,
    display: "flex",
    alignItems: "center",
  },
}));

type Props = {
  numActiveCourses: number;
  numLastActive: number;
  numCertificates: number;
  numExpiringSoon: number;
};

function Summary({
  numActiveCourses,
  numLastActive,
  numCertificates,
  numExpiringSoon,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.summaryRoot}>
      <div className={classes.root}>
        <QuickInfo
          icon="CourseNewCourseGreen"
          text="Active Courses"
          value={numActiveCourses}
        />
        <div className={classes.spacer} />
        <QuickInfo
          icon="CourseCertificates"
          text="Certificates"
          value={numCertificates}
        />
        <div className={classes.spacer} />
        <QuickInfo
          icon="CourseExpiringSoon"
          text="Expire Soon"
          value={numExpiringSoon}
        />
        <div className={classes.spacer} />
        <QuickInfo
          icon="Course_Calendar"
          text="Last active"
          value={`${numLastActive}d`}
        />
      </div>
    </div>
  );
}

export default Summary;
