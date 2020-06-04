import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Button from "sharedComponents/core/Input/Button";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: theme.colors.lightBlue,
        padding: '57px 0'
    },
    defaultTitle: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        textAlign: 'center',
        margin: '10px'
    },
    defaultDesc: {
        fontSize: '40px',
        fontWeight: '800',
        marginTop: '10px',
        maxWidth: '950px',
        textAlign: 'center'
    },
    buttonsTitle: {
        fontSize: '40px',
        fontWeight: '800',
        textAlign: 'center',
        margin: '10px'
    },
    buttonsDesc: {
        fontSize: theme.fontSizes.heading,
        fontWeight: '500',
        marginTop: '10px',
        maxWidth: '950px',
        textAlign: 'center'
    },
    bar: {
        width: '55px',
        height: '3px',
        backgroundColor: theme.colors.navyBlue
    },
    jumpText: {
        fontSize: theme.fontSizes.extraLarge,
        marginRight: '20px'
    },
    button: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        marginRight: '20px',
        height: '53px',
        width: '211px'
    },
    buttons: {
        marginTop: '18px'
    },
    history: {
        alignSelf: 'flex-start',
        marginLeft: '90px',
        marginBottom: '59px'
    },
    extraLarge: {
        fontSize: theme.fontSizes.extraLarge,
    },
    bold: {
        fontWeight: 'bold',
        marginLeft: '3px',
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center'
    }
}));

export type Archetypes = "default" | "buttons";

export type ButtonLink = {
    title: string;
    link: string;
}

type Props = {
    title: string;
    description: string;
    archetype?: Archetypes;
    history?: string[];
    buttons?: ButtonLink[];
};

function PageHeader({ title, description, archetype, history, buttons }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const titleStyle = classes[archetype + "Title"];
    const descStyle = classes[archetype + "Desc"];

  return (
      <div className={classes.root}>
          {history &&
            <div className={classNames(classes.row, classes.history)}>
                {history.map((page: string, index: number) => (
                    index !== history.length - 1
                    ? <div className={classes.extraLarge}>{page} <Icon name="Right_Arrow" size={12}/></div>
                    : <div className={classNames(classes.extraLarge, classes.bold)}>{page}</div>
                ))}
          </div>}
          <div className={titleStyle}>{title}</div>
          {archetype && archetype === "default" && <div className={classes.bar}/>}
          <div className={descStyle}>{description}</div>
          {archetype && archetype === "buttons" &&
            <div className={classNames(classes.row, classes.buttons)}>
                <div className={classes.jumpText}>Jump to:</div>
                {buttons && buttons.map((buttonLink: ButtonLink) => (
                    <Button className={classes.button} onClick={() => console.log(buttonLink.link)}>
                        {buttonLink.title}
                    </Button>
                ))}
            </div>
          }
      </div>
  );
}

export default PageHeader;