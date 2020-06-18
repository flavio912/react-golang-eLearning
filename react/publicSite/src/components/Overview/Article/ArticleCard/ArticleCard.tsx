import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        height: '370px',
        width: '402px',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        borderRadius: '3px'
    },
    image: {
        height: '241px',
        width: '402px',
        borderRadius: '3px 3px 0 0'
    },
    type: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        width: '160px',
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
        margin: '25px 0 20px 0'
    },
    description: {
        textAlign: 'center',
        fontSize: theme.fontSizes.extraLarge,
        margin: '0 25px 20px 25px'
    }
}));

type Props = {
    type: string;
    imageURL: string;
    date: string;
    description: string;
    className?: string;
};

function ArticleCard({ type, imageURL, date, description, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
      <div className={classNames(classes.root, className)}>
          <div className={classes.image} style={{ backgroundImage: `url(${imageURL})` }}>
              <div className={classes.type}>{type.toUpperCase()}</div>
          </div>
          <div className={classes.date}>{date}</div>
          <div className={classes.description}>{description}</div>
      </div>
  );
}

export default ArticleCard;