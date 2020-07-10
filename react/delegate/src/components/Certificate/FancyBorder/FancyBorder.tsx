import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

type Props = {
  children: React.ReactNode,
  paperSize: string
};

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'flex',
    flexDirection: 'column',
    flex: 1,
    width: `calc(${theme.paperSizes.A4.width} - 10px)`,
    height: theme.paperSizes.A4.height,
    padding: 20,
    boxSizing: 'border-box',
    border: '15px solid',
    borderImageSlice: 1,
    borderWidth: 15,
    borderImageSource: theme.verticalGradient
  },
}));

function FancyBorder({ children, paperSize }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      {children}
    </div>
  );
}

export default FancyBorder;
