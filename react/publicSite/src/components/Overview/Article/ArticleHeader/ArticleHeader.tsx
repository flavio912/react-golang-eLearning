import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import CircleBorder, { User } from "sharedComponents/core/CircleBorder";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        backgroundPosition: 'center',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        padding: '110px 0 60px 180px'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
    },
    padding: {
        padding: '120px 0 80px 180px'
    },
    title: {
        fontSize: 40,
        fontWeight: '800',
        color: theme.colors.primaryWhite,
        maxWidth: '615px',
        margin: '20px 0'
    },
    detail: {
        fontSize: theme.fontSizes.large,
        color: theme.colors.primaryWhite
    },
    name: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        color: theme.colors.primaryWhite,
        marginBottom: '5px'
    },
    profile: {
        marginRight: '20px'
    },
    italic: {
        fontStyle: 'italic',
        fontSize: theme.fontSizes.default,
    },
    underline: {
        textDecorationLine: 'underline'
    }
}));

type Props = {
    title: string;
    date: string;
    image: string;
    featured?: string;
    genre?: string;
    author?: User;
    className?: string;
};

function ArticleHeader({ title, date, image, featured, genre, author, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
      <div
        className={classNames(classes.root, featured && classes.padding, className)}
        style={{ backgroundImage: `url(${image})` }}
    >
          {featured && (
            <div className={classes.detail}>{featured}  •  {date}</div>
          )}

          <div className={classNames(classes.title, featured && classes.underline)}>{title}</div>

          {author && (
            <div className={classes.row}>
                <CircleBorder user={author} size={50} className={classes.profile}/>
                <div>
                    <div className={classes.name}>{author.name}</div>
                    <div className={classNames(classes.detail, classes.italic)}>{date} • {genre}</div>
                </div>
            </div>
          )}
      </div>
  );
}

export default ArticleHeader;