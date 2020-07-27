import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import CircleBorder, { User } from "sharedComponents/core/CircleBorder";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        width: '100%',
        backgroundPosition: 'center',
        backgroundSize: 'cover',
        backgroundRepeat: 'no-repeat',
        padding: '75px 0'
    },
    centerer: {
        display: 'flex',
        justifyContent: 'center'
    },
    centered: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'flex-start',
        width: '1000px',
        '@media (max-width: 800px)': {
            alignItems: 'center',
        }
    },
    padding: {
        padding: '50px 0',
    },
    bar: {
        width: (author: boolean) => author ? 0 : '50px',
        margin:(author: boolean) => author ? 0 : '20px 0',
        borderBottom: ['2px', 'solid', theme.colors.primaryWhite],
        '@media (min-width: 800px)': {
            display: 'none',
        }
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
        maxWidth: '615px',
        '@media (max-width: 800px)': {
            textAlign: 'center',
        }
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
        margin: '2px 20px 0 0'
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
    const classes = useStyles({ author, theme });

  return (
      <div
        className={classNames(classes.root, author && classes.padding, className)}
        style={{ backgroundImage: `url(${image})` }}
    >
        <div className={classes.centerer}>
            <div className={classes.centered}>
                {featured && (
                    <div className={classNames(classes.detail, classes.marginBottom)}>{featured}  •  {date}</div>
                )}
                <div className={classes.bar} />
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
      </div>
  );
}

export default ArticleHeader;