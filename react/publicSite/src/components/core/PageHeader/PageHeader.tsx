import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        height: '454px',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: theme.colors.lightBlue
    },
    title: {
        fontSize: theme.fontSizes.extraLarge,
        fontWeight: '800',
        textAlign: 'center',
        margin: '15px'
    },
    desc: {
        fontSize: '40px',
        fontWeight: '800',
        marginTop: '10px',
        maxWidth: '950px',
        textAlign: 'center'
    },
    bar: {
        width: '55px',
        height: '3px',
        backgroundColor: theme.colors.navyBlue
    }
}));

type Props = {
    title: string;
    description: string;
};

function PageHeader({ title, description }: Props) {
    const theme = useTheme();
    const classes = useStyles({ theme });

  return (
      <div className={classes.root}>
          <div className={classes.title}>{title.toUpperCase()}</div>
          <div className={classes.bar}/>
          <div className={classes.desc}>{description}</div>
      </div>
  );
}

export default PageHeader;