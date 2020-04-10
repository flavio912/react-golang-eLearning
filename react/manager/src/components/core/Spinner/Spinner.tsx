import * as React from "react";
import { createUseStyles } from "react-jss";
import classNames from "classnames";
import { ReactComponent as SpinnerSVG } from "../../../assets/Loading_Screen_Donut.svg";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  "@keyframes spin": {
    from: { transform: "rotate(0deg)" },
    to: { transform: "rotate(360deg)" },
  },
  spinner: {
    animation: "$spin 2s infinite linear",
    width: 100,
    height: 100,
  },
}));

type Props = {
  size?: number;
};

function Spinner({ size }: Props) {
  const { spinner } = useStyles();
  return (
    <SpinnerSVG
      className={spinner}
      style={size ? { width: size, height: size } : undefined}
    />
  );
}

export default Spinner;
