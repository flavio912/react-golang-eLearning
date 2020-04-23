import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {},
  title: {},
  titleText: {
    margin: 0,
    color: theme.colors.primaryBlack,
    fontSize: 33,
    fontWeight: 800,
  },
  subtitleText: {
    margin: 0,
    color: theme.colors.textGrey,
    fontSize: theme.fontSizes.default,
    fontWeight: 300,
    marginTop: 5,
  },
}));

type Props = {
  title: string;
  subTitle: string;
};

function PageTitle({ title, subTitle }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.title}>
      <h1 className={classes.titleText}>{title}</h1>
      <p className={classes.subtitleText}>{subTitle}</p>
    </div>
  );
}

export default PageTitle;
