import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

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
  sideComponent: {
    marginLeft: 10,
  },
  titlesHolder: {},
  divider: {
    height: 40,
    background: theme.colors.borderGrey,
    width: 1,
    margin: [0, theme.spacing(3)],
  },
  backTitle: {
    cursor: "pointer",
    marginBottom: 8,
    display: "flex",
    alignItems: "center",
  },
  backText: {
    color: theme.colors.textBlue,
    fontSize: theme.fontSizes.xSmall,
    marginLeft: 7,
  },
}));

type Props = {
  title: string;
  subTitle: string;
  sideText?: string;
  sideComponent?: React.ReactNode;
  backProps?: {
    text: string;
    onClick: Function;
  };
};

function PageTitle({
  title,
  subTitle,
  sideText,
  sideComponent,
  backProps,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.pageTitleRoot}>
      <div className={classes.titlesHolder}>
        {backProps && (
          <div
            className={classes.backTitle}
            onClick={() => backProps.onClick()}
          >
            <Icon name="ArrowLeft" size={12} />
            <span className={classes.backText}>{backProps.text}</span>
          </div>
        )}
        <h1 className={classes.titleText}>{title}</h1>
        <p className={classes.subtitleText}>{subTitle}</p>
      </div>
      {sideText && (
        <div className={classes.sideInfo}>
          <div className={classes.divider} />
          <p className={classes.sideText}>{sideText}</p>
        </div>
      )}
      {sideComponent && (
        <span className={classes.sideComponent}>{sideComponent}</span>
      )}
    </div>
  );
}

export default PageTitle;
