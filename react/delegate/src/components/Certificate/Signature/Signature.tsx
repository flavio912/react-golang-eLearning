import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

type Props = {
  width: number,
  height: number,
  imgSrc: string,
};

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'block',
  }
}));

function Signature({ width, height, imgSrc }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <img src={imgSrc} width={width} height={height} />
    </div>
  );
}

export default Signature;
