import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import classNames from "classnames";
import QuickInfo from "../QuickInfo";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    justifyContent: "space-between",
  },
  spacer: {
    width: theme.spacing(3),
  },
  quickOverviewRoot: {
    flex: 1,
    display: "flex",
    alignItems: "center",
  },
}));

type Props = {
  purchasedCourses: number;
  numDelegates: number;
  numValidCertificates: number;
  numCertificatesExpiringSoon: number;
};

function QuickOverview({
  purchasedCourses,
  numDelegates,
  numValidCertificates,
  numCertificatesExpiringSoon,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.quickOverviewRoot}>
      <div className={classes.root}>
        <QuickInfo
          icon="CourseNewCourseGreen"
          text="Courses"
          value={purchasedCourses}
        />
        <div className={classes.spacer} />
        <QuickInfo
          icon="Icon_Delegates"
          text="Delegates"
          value={numDelegates}
        />
        <div className={classes.spacer} />
        <QuickInfo
          icon="CourseCertificates"
          text="Certificates"
          value={numValidCertificates}
        />
        <div className={classes.spacer} />
        <QuickInfo
          icon="CourseExpiringSoon"
          text="Expire Soon"
          value={numCertificatesExpiringSoon}
        />
      </div>
    </div>
  );
}

export default QuickOverview;
