import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  pageTitleRoot: {
    display: "flex",
  },
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
  sideInfo: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  sideText: {
    fontSize: theme.fontSizes.default,
    fontWeight: 700,
  },
  titlesHolder: {},
  divider: {
    height: 40,
    background: theme.colors.borderGrey,
    width: 1,
    margin: [0, theme.spacing(3)],
  },
}));

type Props = {
  title: string;
  subTitle: string;
  sideText?: string;
};

function PageTitle({ title, subTitle, sideText }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classes.pageTitleRoot}>
      <div className={classes.titlesHolder}>
        <h1 className={classes.titleText}>{title}</h1>
        <p className={classes.subtitleText}>{subTitle}</p>
      </div>
      {sideText && (
        <div className={classes.sideInfo}>
          <div className={classes.divider} />
          <p className={classes.sideText}>{sideText}</p>
        </div>
      )}
    </div>
  );
}

export default PageTitle;
