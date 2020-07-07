import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import SignatureCanvas from "react-signature-canvas";

type Props = {
  width: number,
  height: number
};

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'block'
  },
  sigCanvas: {
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: 5
  }
}));

function Signature({ width, height }: Props) {
  const classes = useStyles();

  return (
    <div className={classes.container}>
      <SignatureCanvas
        canvasProps={{width: width, height: height, className: classes.sigCanvas}} />
    </div>
  );
}

export default Signature;
