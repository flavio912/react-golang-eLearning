import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Icon, { IconNames } from "sharedComponents/core/Icon";
import CircleBorder from "sharedComponents/core/CircleBorder";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        alignItems: 'center'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
    },
    panelRow: {
        '@media (min-width: 1000px)': {
            marginTop: ({ noBorders }: { noBorders: boolean }) => noBorders ? '50px' : 0,
        },
        '@media (max-width: 600px)': {
            flexDirection: 'column',
            justifyContent: 'center'
        }
    },
    imageRow: {
        '@media (max-width: 1000px)': {
            flexDirection: 'column',
            alignItems: 'flex-start'
        }
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
    },
    stepColumn: {
        '@media (max-width: 600px)': {
            width: '287px',
        }
    },
    iconColumn: {
        '@media (max-width: 600px)': {
            width: '400px',
            alignItems: 'center'
        }
    },
    image: {
        height: '176px',
        width: '287px',
        borderRadius: '3px',
        boxShadow: '4px 2px 10px -2px rgba(0,0,0,0.19)'
    },
    imageTitle: {
        fontSize: theme.fontSizes.tinyHeading,
        '@media (max-width: 1000px)': {
            margin: '10px 0'
        }
    },
    iconTitle: {
        fontSize: theme.fontSizes.heading,
    },
    bold: {
        fontWeight: 'bold',
        marginBottom: '15px'
    },
    description: {
        fontSize: 15,
        color: theme.colors.textGrey,
        lineHeight: '25px'
    },
    imageDescription: {
        '@media (max-width: 1000px)': {
            maxWidth: '287px'
        }
    },
    iconDescription: {
        '@media (max-width: 600px)': {
            textAlign: 'center',
            margin: '0 15px'
        }
    },
    stepNumber: {
        alignSelf: 'flex-start',
        position: 'relative',
        right: '28px',
        top: '10px',
        minWidth: '54px',
        '@media (max-width: 1000px)': {
            order: -1,
            left: '10px',
            top: '28px',
        }
    },
    blueLine: {
        width: '111px',
        borderBottom: ['4px', 'solid', theme.colors.navyBlue],
        marginBottom: '15px'
    },
    panel: {
        paddingTop: '25px',
        '@media (min-width: 600px)': {
            padding: '0 0 50px 50px'
        }
    },
    borderTop: {
        '@media (min-width: 600px)': {
            borderTop: ['0.5px', 'solid', theme.colors.borderGrey],
            padding: '50px 0 0 50px'
        }
    },
    borderRight: {
        '@media (min-width: 600px)': {
            borderRight: ['0.5px', 'solid', theme.colors.borderGrey],
            padding: '0 50px 50px 0'
        }
    },
    borderTopRight: {
        '@media (min-width: 600px)': {
            borderTop: ['0.5px', 'solid', theme.colors.borderGrey],
            borderRight: ['0.5px', 'solid', theme.colors.borderGrey],
            padding: '50px 50px 0 0'
        }
    }
}));

export type Panel = {
    title: string;
    desciption: string;
    iconName?: IconNames;
    imageURL?: string;
}

type Props = {
    panels: Panel[];
    noBorders?: boolean;
    className?: string;
};

function FourPanel({ panels, noBorders, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ noBorders, theme });

    const panel = (index: number) => (
        panels[index].imageURL ? (
            <div className={classNames(classes.imageRow, classes.row)}>
                <img className={classes.image} src={panels[index].imageURL} />
                <CircleBorder
                    className={classes.stepNumber}
                    text={index + 1}
                    fontSize={30}
                    borderSize={10}
                />
                <div className={classNames(classes.stepColumn, classes.column)}>
                    <div className={classNames(classes.bold, classes.imageTitle)}>{panels[index].title}</div>
                    <div className={classes.blueLine}/>
                    <div className={classNames(classes.description, classes.imageDescription)}>{panels[index].desciption}</div>
                </div>
            </div>
        ) : (
            <div className={classNames(classes.iconColumn, classes.column)}>
                <Icon
                    //@ts-ignore}
                    name={panels[index].iconName ? panels[index].iconName : 'TTC_Logo_Icon'}
                    size={56}
                    style={{ marginBottom: '15px' }}
                />
                <div className={classNames(classes.bold, classes.iconTitle)}>{panels[index].title}</div>
                <div className={classNames(classes.description, classes.iconDescription)}>{panels[index].desciption}</div>
            </div>
        )
    )

    if (panels.length < 4) {
        console.error("Must use array of length 4")
        return null;
    }

    const noBorder = { border: 'none', padding: '0 25px' };

  return (
      <div className={classNames(classes.root, classes.column, className)}>
          <div className={classNames(classes.panelRow, classes.row)}>
            <div
                className={classNames(classes.panel, classes.column, classes.borderRight)}
                style={noBorders ? noBorder : {}}
            >
                {panel(0)}
            </div>
            <div
                className={classNames(classes.panel, classes.column)}
                style={noBorders ? noBorder : {}}
            >
                {panel(1)}
            </div>
          </div>
          <div className={classNames(classes.panelRow, classes.row)}>
            <div
                className={classNames(classes.panel, classes.column, classes.borderTopRight)}
                style={noBorders ? noBorder : {}}
            >
                {panel(2)}
            </div>
            <div
                className={classNames(classes.panel, classes.column, classes.borderTop)}
                style={noBorders ? noBorder : {}}
            >
                {panel(3)}
            </div>
          </div>
      </div>
  );
}

export default FourPanel;