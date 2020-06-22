import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { useRouter } from 'found';
import { Theme } from "helpers/theme";
import Icon, { IconNames } from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        //height: '606px',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        flexWrap: 'wrap'
    },
    container: {
        display: 'flex',
        flex: 1,
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'flex-start'
    },
    title: {
        fontSize: theme.fontSizes.heading,
        fontWeight: '500',
        color: theme.colors.textGrey,
        marginBottom: '15px'
    },
    subtitle: {
        fontSize: theme.fontSizes.extraLargeHeading,
        fontWeight: '800',
        marginBottom: '20px'
    },
    description: {
        fontSize: theme.fontSizes.xSmallHeading,
        color: theme.colors.textGrey,
        marginBottom: '25px'
    },
    link: {
        cursor: 'pointer',
        fontSize: theme.fontSizes.xSmallHeading,
        fontWeight: 'bold',
        color: theme.colors.navyBlue,
        marginBottom: '25px'
    },
    image: {
        height: '457px',
        maxWidth: '650px'
    },
    text: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '500',
        color: theme.colors.textGrey,
        marginTop: '10px'
    },
    margin: {
        marginBottom: '70px'
    },
    center: {
        alignItems: 'center'
    },
    maxWidth: {
        minWidth: '400px',
        maxWidth: '400px',
        marginRight: '100px'
    },
    rowReverse: {
        flexDirection: 'row-reverse',
    },
    marginLeft: {
        margin: '0 0 0 100px'
    }
}));

export type Link = {
    title: string;
    link: string;
}

export type Row = {
    iconName: IconNames;
    text: string;
    link: Link;
}

type Props = {
    title?: string;
    subtitle?: string;
    description?: string;
    link?: Link;
    image?: string;
    stack?: Row[];
    textRight?: boolean;
    className?: string;
};

function ImageWithText({ title, subtitle, description, link, image, stack, textRight, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const { router } = useRouter();
    const onClick = (link?: string) => {
        link && router.push(link);
    }
    
  return (
      <div className={classNames(classes.root, textRight && classes.rowReverse, className)}>
        {stack ? (
            <div className={classNames(classes.container, classes.maxWidth, textRight && classes.marginLeft)}>
                {stack.map((row: Row, index: number) => (
                    <div className={classNames(classes.root, index !== stack.length - 1 && classes.margin)}>
                        <Icon name={row.iconName} size={68} style={{ marginRight: '40px' }} />
                        <div onClick={() => onClick(row.link?.link)}>
                            <div className={classes.link}>{row.link?.title}</div>
                            <div className={classes.text}>{row.text}</div>
                        </div>
                    </div>
                ))}
            </div>
        ) : (
            <div className={classNames(classes.container, classes.maxWidth, textRight && classes.marginLeft)}>
                <div className={classes.title}>{title}</div>
                <div className={classes.subtitle}>{subtitle}</div>
                <div className={classes.description}>{description}</div>
                <div className={classes.link} onClick={() => onClick(link?.link)}>
                    {link?.title}
                    {link && <Icon name="ArrowRightNavyBlue" size={10} style={{ margin: '0 0 1px 5px' }} />}
                </div>
            </div>
          )
        }
          <div className={classNames(classes.container, classes.center)}>
            <img className={classes.image} src={image} />
          </div>
      </div>
  );
}

export default ImageWithText;