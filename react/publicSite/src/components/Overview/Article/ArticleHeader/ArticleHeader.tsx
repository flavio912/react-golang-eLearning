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
        padding: '10% 0 10% 0',
        alignItems: 'center',
        '@media (max-width: 700px)': {
            padding: '5%',
        }
    },
    padding: {
        paddingBottom: '5%',
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
        marginTop: '20px'
    },
    title: {
        fontSize: 40,
        fontWeight: '800',
        color: theme.colors.primaryWhite,
        maxWidth: '615px'
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
    },
    marginBottom: {
        marginBottom: '6px'
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
        className={classNames(classes.root, author && classes.padding, className)}
        style={{ backgroundImage: `url(${image})` }}
    >
        <div>
            {featured && (
                <div className={classNames(classes.detail, classes.marginBottom)}>{featured}  •  {date}</div>
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
      </div>
  );
}

export default ArticleHeader;