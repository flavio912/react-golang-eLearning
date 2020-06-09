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
        alignSelf: 'center',
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        textAlign: 'center',
        marginBottom: '10px'
    },
    defaultDesc: {
        fontSize: '40px',
        fontWeight: '800',
        marginTop: '10px',
        maxWidth: '950px',
        textAlign: 'center',
    },
    buttonsTitle: {
        fontSize: '40px',
    },
    buttonsDesc: {
        fontSize: theme.fontSizes.heading,
        fontWeight: '500',
    },
    courseTitle: {
        alignSelf: 'flex-start',
        fontSize: theme.fontSizes.heading,
        textAlign: 'left',
    },
    courseDesc: {
        fontSize: theme.fontSizes.tinyHeading,
        fontWeight: '500',
        textAlign: 'left',
    },
    bar: {
        alignSelf: 'center',
        width: '55px',
        height: '3px',
        backgroundColor: theme.colors.navyBlue
    },
    jumpText: {
        marginRight: '20px'
    },
    updatedText: {
        marginLeft: '25px'
    },
    button: {
        fontWeight: '800',
        marginRight: '20px',
        height: '53px',
        width: '211px'
    },
    buttons: {
        marginTop: '18px'
    },
    times: {
        marginTop: '31px'
    },
    history: {
        alignSelf: 'flex-start',
        marginLeft: '90px',
        marginBottom: '59px'
    },
    course: {
        alignSelf: 'flex-start',
        marginLeft: '90px'
    },
    component: {
        position: 'absolute',
        top: '15%',
        right: '5%'
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
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'flex-start',
        alignItems: 'flex-start'
    },
    spacer: {
        minWidth: '450px'
    }
}));

export type Archetypes = "default" | "buttons" | "course";

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
    estimatedTime?: string;
    lastUpdated?: string;
    sideComponent?: React.Component;
    className?: string;
};

function PageHeader({ title, description, archetype, history, buttons, estimatedTime, lastUpdated, sideComponent, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

    const titleStyle = classes[archetype + "Title"];
    const descStyle = classes[archetype + "Desc"];

  return (
      <div className={classNames(classes.root, className)}>

          {history &&
            <div className={classNames(classes.row, classes.history)}>
                {history.map((page: string, index: number) => (
                    index !== history.length - 1
                    ? <div className={classes.extraLarge}>{page} <Icon name="Right_Arrow" size={12}/></div>
                    : <div className={classNames(classes.extraLarge, classes.bold)}>{page}</div>
                ))}
          </div>}

          <div className={classNames(classes.row, archetype && archetype === "course" && classes.course)}>
              <div className={classes.column}>
                <div className={classNames(classes.defaultTitle, titleStyle)}>{title}</div>
                {archetype && archetype === "default" && <div className={classes.bar}/>}
                <div className={classNames(classes.defaultDesc, descStyle)}>{description}</div>

                {archetype && archetype === "course" &&
                    <div className={classNames(classes.row, classes.times)}>
                        <div className={classes.extraLarge}>{<strong>Estimated Time:</strong>} {estimatedTime}</div>
                        <div className={classNames(classes.updatedText, classes.extraLarge)}>{<strong>Last Updated:</strong>} {lastUpdated}</div>
                    </div>
                } 

                {archetype && archetype === "buttons" &&
                    <div className={classNames(classes.row, classes.buttons)}>
                        <div className={classNames(classes.jumpText, classes.extraLarge)}>Jump to:</div>
                        {buttons && buttons.map((buttonLink: ButtonLink) => (
                            <Button
                                className={classNames(classes.button, classes.extraLarge)}
                                onClick={() => console.log(buttonLink.link)}
                            >
                                {buttonLink.title}
                            </Button>
                        ))}
                    </div>
                }
              </div>
              {archetype && archetype === "course" &&
                <div className={classes.spacer} />
              }
          </div>
          {archetype && archetype === "course" &&
                <div className={classes.component}>
                    {sideComponent}
                </div>
              }
      </div>
  );
}

export default PageHeader;