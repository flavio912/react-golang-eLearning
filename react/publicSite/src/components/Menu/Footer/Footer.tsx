import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { useRouter } from 'found';
import { Theme } from 'helpers/theme';
import Icon from 'sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        //height: '534px',
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: theme.colors.footerBlue,
        padding: '35px 75px 30px 50px'
    },
    row: {
        display: 'flex',
        flexWrap: 'wrap',
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'flex-start',
        '@media (max-width: 650px)': {
            flexDirection: 'column',
            alignItems: 'center',
        }
    },
    border: {
        paddingBottom: '100px',
        marginBottom: '20px',
        borderBottom: ['1px', 'solid', theme.colors.footerGrey],
        //alignItems: 'flex-start',
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'flex-start',
        margin: '25px 25px',
        //marginRight: '50px',
        '@media (max-width: 650px)': {
            alignItems: 'center'
        }
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
        flexDirection: 'row',
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
        color: theme.colors.footerGrey,
        '@media (max-width: 650px)': {
            marginTop: '25px',
            textAlign: 'center'
        }
    },
    mobileBorder: {
        '@media (max-width: 650px)': {
            borderBottom: ['1px', 'solid', theme.colors.footerGrey],
            paddingBottom: '25px'
        }
    }
}));

export type Column = {
  id: number;
  header: string;
  links: Link[];
};

export type Link = {
  id: number;
  name: string;
  link: string;
  alert?: Alert;
};

export type Alert = {
  type: 'new' | 'increase';
  value: string | number;
};

type Props = {
  columns: Column[];
  className?: string;
};

function Footer({ columns, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const { router } = useRouter();
  const onClick = (link?: string) => {
    link && router.push(link);
  };

  return (
      <div className={classNames(classes.root, className)}>
          <div className={classNames(classes.row, classes.border)}>
              <div className={classNames(classes.column, classes.mobileBorder)}>
                  <Icon name="Blue_TTC_Logo_Icon" size={46} />
                  <div className={classes.message}>Redefining the <strong>future</strong> of <strong>compliance</strong> and training</div>
                  <Icon name="AviationSecurityCert" style={{height: '74px', width: '152px' }} />
              </div>
              {columns && columns.map((column: Column) => (
                  <div key={column.id} className={classes.column}>
                      <div className={classes.header}>{column.header}</div>
                      {column.links && column.links.map((link: Link) => (
                          <div key={link.id} className={classNames(classes.row, classes.center)}>
                              <div
                                onClick={() => onClick(link.link)}
                                className={classes.link}
                              >
                                  {link.name}
                              </div>
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
          <div className={classNames(classes.row, classes.center)}>
            <Icon
              name="Twitter_Logo"
              size={25}
              style={{ marginRight: '10px' }}
            />
            <Icon
              name="Facebook_Logo"
              size={25}
              style={{ marginRight: '10px' }}
            />
            <Icon name="LinkedIn_Logo" size={25} />
          </div>
          <div className={classes.copyright}>
            © 2020 TTC Hub. All Rights Reserved. Registered in England | Company
            registration number 10849230.
          </div>
        </div>

      
    </div>
  );
}

export default Footer;
