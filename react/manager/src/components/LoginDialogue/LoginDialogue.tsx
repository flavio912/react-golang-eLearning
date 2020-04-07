import * as React from "react";
import Card from "../core/Card";
import FancyInput from "./FancyInput";
import FancyButton from "./FancyButton";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    width: 340,
  },
  heading: {
    fontWeight: 800,
    color: theme.colors.primaryBlack,
  },
  subheading: {
    color: theme.colors.secondaryGrey,
    fontWeight: 400,
    fontSize: 15,
    marginTop: 0,
    marginBottom: theme.spacing(2),
  },
  link: {
    margin: [15, 0, 0, 0],
    textAlign: "center",
    color: theme.colors.textBlue,
    fontSize: theme.fontSizes.small,
  },
}));

type Props = {};

function LoginDialogue(props: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <Card padding="medium" className={classes.root}>
      <h1 className={classes.heading}>Login to TTC Hub</h1>
      <p className={classes.subheading}>
        Glad to have you back, please enter your login details to proceed
      </p>
      <FancyInput label="Email" labelColor={"#5CC301"} type={"email"} />
      <FancyInput label="Password" labelColor={"#5CC301"} type={"password"} />
      <FancyButton text="Login to TTC" />
      <a className={classes.link} href="">
        I don't have a TTC Hub account
      </a>
    </Card>
  );
}

export default LoginDialogue;
