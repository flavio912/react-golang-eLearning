import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        //height: '534px',
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: theme.colors.footerBlue,
        padding: '60px 100px 30px 100px'
    },
    row: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'flex-start',
    },
    border: {
        paddingBottom: '100px',
        marginBottom: '20px',
        borderBottom: ['1px', 'solid', theme.colors.footerGrey]
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'flex-start',
        marginRight: '50px'
    },
    header: {
        fontSize: theme.fontSizes.extraLarge,
        color: theme.colors.primaryWhite,
        fontWeight: 'bold',
        marginBottom: '10px',
    },
    link: {
        cursor: 'pointer',
        fontSize: theme.fontSizes.large,
        color: theme.colors.footerGrey,
        margin: '5px 0'
    },
    alert: {
        height: '17px',
        marginLeft: '8px',
        borderRadius: '3px',
        fontSize: theme.fontSizes.xTiny,
        fontWeight: '800',
        color: theme.colors.primaryWhite,
    },
    triangle: {
        width: 0,
        height: 0,
        margin: '2px 4px 0 0',
        borderLeft: '3px solid transparent',
        borderRight: '3px solid transparent',
        borderBottom: ['6px', 'solid', theme.colors.primaryWhite],
    },
    new: {
        width: '33px',
        backgroundColor: theme.colors.navyBlue,
    },
    increase: {
        width: '49px',
        backgroundColor: theme.colors.primaryGreen,
    },
    center: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    message: {
        fontSize: theme.fontSizes.smallHeading,
        color: theme.colors.primaryWhite,
        maxWidth: '275px',
        margin: '10px 0 50px 0'
    },
    copyright: {
        fontSize: theme.fontSizes.xSmall,
        color: theme.colors.footerGrey
    }
}));

export type Column = {
    header: string;
    links: Link[];
}

export type Link = {
    name: string;
    link: string;
    alert?: Alert;
}

export type Alert = {
    type: "new" | "increase";
    value: string | number;
}

type Props = {
    columns: Column[];
    className?: string;
};

function Footer({ columns, className }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
      <div className={classNames(classes.root, className)}>
          <div className={classNames(classes.row, classes.border)}>
              <div className={classes.column}>
                  <Icon name="TTC_Logo_Icon" size={46} />
                  <div className={classes.message}>Redefining the <strong>future</strong> of <strong>compliance</strong> and training</div>
                  <Icon name="SampleImage_ClassroomCoursesDetail_Feat" style={{height: '74px', width: '152px'}} />
              </div>
              {columns && columns.map((column: Column) => (
                  <div className={classes.column}>
                      <div className={classes.header}>{column.header}</div>
                      {column.links && column.links.map((link: Link) => (
                          <div className={classNames(classes.row, classes.center)}>
                              <div className={classes.link}>{link.name}</div>
                              {link.alert && (
                                <div
                                    className={classNames(classes.alert, classes.center, classes[link.alert.type])}
                                >
                                    {link.alert.type === "increase" && (<div className={classes.triangle} />)} {link.alert.value}
                                </div>
                              )}
                          </div>
                      ))}
                  </div>
              ))}
          </div>
          <div className={classes.row}>
              <div>Socials</div>
              <div className={classes.copyright}>© 2020 TTC Hub. All Rights Reserved. Registered in England | Company registration number 10849230.</div>
          </div>
      </div>
  );
}

export default Footer;