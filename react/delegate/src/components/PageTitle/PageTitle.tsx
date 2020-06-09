import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';
import classnames from 'classnames';
const useStyles = createUseStyles((theme: Theme) => ({
  pageTitleRoot: {
    display: 'flex'
  },
  titleText: {
    margin: 0,
    color: theme.colors.primaryBlack,
    fontSize: 41,
    fontWeight: 900,
    letterSpacing: -1.02,
    lineHeight: `43px`
  },
  titlesHolder: {},
  backTitle: {
    cursor: 'pointer',
    marginBottom: 34,
    display: 'flex',
    alignItems: 'center'
  },
  backTextBold: {
    marginLeft: 10,
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.default,
    fontWeight: 700,
    lineHeight: `18px`,
    letterSpacing: -0.35,
    marginRight: 5
  },
  backText: {
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.default,
    fontWeight: 400,
    lineHeight: `18px`,
    letterSpacing: -0.35
  }
}));

type Props = {
  title: string;
  rootClassName?: string;
  backProps?: {
    text: string;
    onClick: Function;
    className?: string;
  };
};

function PageTitle({ title, rootClassName, backProps }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classnames(classes.pageTitleRoot, rootClassName)}>
      <div className={classes.titlesHolder}>
        {backProps && (
          <div
            className={classnames(classes.backTitle, backProps.className)}
            onClick={() => backProps.onClick()}
          >
            <Icon name="ArrowLeft" size={11} />
            <span className={classes.backTextBold}>Back to</span>
            <span className={classes.backText}>{backProps.text}</span>
          </div>
        )}
        <h1 className={classes.titleText}>{title}</h1>
      </div>
    </div>
  );
}

export default PageTitle;
