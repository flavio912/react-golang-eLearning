import * as React from "react";
import { createUseStyles } from "react-jss";
import { ReactComponent as SpinnerSVG } from "../../../assets/Loading_Screen_Donut.svg";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  "@keyframes spin": {
    from: { transform: "rotate(0deg)" },
    to: { transform: "rotate(360deg)" },
  },
  spinner: {
    width: 100,
    height: 100,
    animation: "$spin 2s infinite linear",
  },
}));

type Props = {};

function Spinner(props: Props) {
  const classes = useStyles();
  return <SpinnerSVG className={classes.spinner} />;
}

export default Spinner;
