import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { useRouter } from 'found';
import { Theme } from "helpers/theme";
import CircleBorder from 'sharedComponents/core/CircleBorder';
import Card, { PaddingOptions } from "sharedComponents/core/Cards/Card";

const useStyles = createUseStyles((theme: Theme) => ({
    mainRoot: {
        display: 'flex',
        flexWrap: 'wrap'
    },
    root: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between',
        alignItems: 'flex-start'
    },
    text: {
        fontSize: theme.fontSizes.default,
        fontWeight: '400',
        color: theme.colors.textGrey,
        margin: '20px 0 15px 0'
    },
    quote: {
        fontSize: theme.fontSizes.tinyHeading,
        fontStyle: 'italic',
        color: "#34373A",
        fontWeight: '600',
        maxWidth: '270px',
        margin: '15px 0'
    },
    quotationMarks: {
        height: '43px',
        fontSize: '75px',
        fontWeight: '800',
        color: theme.colors.navyBlue
    },
    name: {
        fontSize: theme.fontSizes.xSmall,
        fontWeight: '600',
        marginBottom: '3px'
    },
    title: {
        fontSize: theme.fontSizes.xSmall,
        fontWeight: '600',
        color: theme.colors.textGrey
    },
    profilePicture: {
        marginRight: '10.5px'
    },
    link: {
        cursor: 'pointer',
        width: '111px',
        fontSize: theme.fontSizes.default,
        fontWeight: 'bold',
        color: theme.colors.navyBlue,
        paddingBottom: '1px',
        borderBottom: ['1.5px', 'solid', theme.colors.borderGrey]
    },
    logoContainer: {
        display: 'flex',
        flex: 1
    },
    logo: {
        maxWidth: '200px',
        maxHeight: '60px',
        alignSelf: 'flex-start',
        '@media (max-width: 500px)': {
            width: '100%',
        }
    },
    quoteLogo: {
        maxWidth: '200px',
        maxHeight: '35px',
        alignSelf: 'flex-start',
        '@media (max-width: 500px)': {
            width: '100%',
        },
        margin: '25px 0 30px 0'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center'
    }
}));

export type Author = {
    name: string;
    title: string;
    quote: string;
}

type Props = {
    logoURL: string;
    link: string;
    text?: string;
    author?: Author;
    padding?: PaddingOptions;
    className?: string;
};

function BrandCard({ logoURL, link, text, author, padding = "large", className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const { router } = useRouter();
    const onClick = (link?: string) => {
        link && router.push(link);
    }

  return (
      <Card padding={padding} className={classNames(classes.mainRoot, className)}>
          {author ? (
              <div className={classes.root}>
                <div className={classes.quotationMarks}>“</div>
                <div className={classes.quote}>{author.quote}</div>
                <div className={classes.row}>
                    <CircleBorder
                        borderType="plain"
                        user={{ name: author.name }}
                        className={classes.profilePicture}
                    />
                    <div>
                        <div className={classes.name}>{author.name}</div>
                        <div className={classes.title}>{author.title}</div>
                    </div>
                </div>
                <img src={logoURL} className={classes.quoteLogo} />
                <div
                    className={classes.link}
                    onClick={() => onClick(link)}
                >
                    Read their story
                </div>
              </div>
          ) : (
            <div className={classes.root}>
                <div className={classes.logoContainer}>
                    <img src={logoURL} className={classes.logo} />
                </div>
                <div className={classes.text}>{text}</div>
                <div
                    className={classes.link}
                    onClick={() => onClick(link)}
                >
                    Read their story
                </div>
            </div>
          )
        }
      </Card>
  );
}

export default BrandCard;