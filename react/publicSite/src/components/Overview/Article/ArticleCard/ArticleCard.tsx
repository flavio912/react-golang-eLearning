import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    cursor: 'pointer',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    borderRadius: '3px',
    backgroundColor: theme.colors.backgroundGrey,
    '&:hover': {
      boxShadow: '4px 9px 39px rgba(0,0,0,0.19)',
      backgroundColor: theme.colors.primaryWhite,
    }
  },
  image: {
    display: 'flex',
    alignItems: 'flex-start',
    justifyContent: 'flex-start',
    height: '241px',
    width: '100%',
    borderRadius: '3px 3px 0 0'
  },
  type: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '0 20px',
    height: '33px',
    backgroundColor: theme.colors.navyBlue,
    fontSize: theme.fontSizes.xSmall,
    color: theme.colors.primaryWhite,
    fontWeight: '800',
    borderRadius: '3px 0 10px 0'
  },
  date: {
    fontSize: theme.fontSizes.xSmall,
    fontWeight: '800',
    color: theme.colors.textGrey,
    margin: '15px 0'
  },
  description: {
    textAlign: 'center',
    fontSize: theme.fontSizes.large,
    margin: '0 25px 20px 25px'
  }
}));

export type ArticleDetails = {
  type: string;
  imageURL: string;
  date: string;
  description: string;
}

type Props = {
  article: ArticleDetails;
  onClick: () => void;
  className?: string;
};

function ArticleCard({ article, onClick, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div
      className={classNames(classes.root, className)}
      onClick={onClick}
    >
      <div
        className={classes.image}
        style={{ backgroundImage: `url(${article.imageURL})` }}
      >
        <div className={classes.type}>{article.type.toUpperCase()}</div>
      </div>
      <div className={classes.date}>{article.date}</div>
      <div className={classes.description}>{article.description}</div>
    </div>
  );
}

export default ArticleCard;
