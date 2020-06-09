import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  statusRoot: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
  },
  textContent: {
    display: "flex",
    flexDirection: "column",
    marginLeft: 12,
  },
  statusText: {
    fontSize: theme.fontSizes.tiny,
    color: theme.colors.secondaryBlack,
  },
  expiresText: {
    fontSize: theme.fontSizes.small,
    color: theme.colors.textBlue,
  },
}));

type Props = {
  isComplete: boolean;
  expires?: string;
};

function Status({ isComplete, expires }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.statusRoot}>
      {isComplete ? (
        <Icon size={17} name="CourseStatus_Completed" />
      ) : (
        <Icon size={17} name="CourseStatus_Incomplete" />
      )}
      <div className={classes.textContent}>
        <span className={classes.statusText}>
          {isComplete ? "Completed" : "Incomplete"}
        </span>
        {expires && (
          <span className={classes.expiresText}>[Expires {expires}]</span>
        )}
      </div>
    </div>
  );
}

export default Status;
