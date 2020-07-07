import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

type Props = {
  children: React.ReactNode
};

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'flex',
    flexDirection: 'column',
    flex: 1,
    width: '100%',
    height: '100%',
    padding: '20px',
    boxSizing: 'border-box',
    border: '10px solid',
    borderImageSlice: 1,
    borderWidth: 10,
    borderImageSource: theme.verticalGradient
  },
}));

function FancyBorder({ children }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      {children}
    </div>
  );
}

export default FancyBorder;
