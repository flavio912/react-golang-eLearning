import * as React from "react";
import Card from "../../sharedComponents/core/Card";
import FancyInput from "./FancyInput";
import FancyButton from "./FancyButton";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import { ReactComponent as Logo } from "../../assets/logo/ttc-logo.svg";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    width: 378,
    background: "white",
  },
  logoContainer: {
    padding: [30, 0, 20],
  },
  logo: {
    height: 70,
  },
  heading: {
    fontWeight: 800,
    color: theme.colors.primaryBlack,
  },
  subheading: {
    color: theme.colors.textGrey,
    fontWeight: 300,
    fontSize: 15,
    marginTop: 0,
    marginBottom: theme.spacing(2),
  },
  link: {
    margin: [15, 0, 30, 0],
    textAlign: "center",
    color: theme.colors.textBlue,
    fontSize: theme.fontSizes.small,
  },
  errMessage: {
    color: "#43454a",
    fontWeight: 200,
    fontSize: 15,
    textAlign: "center",
    margin: 3,
  },
}));

type Props = {
  onSubmit: (
    email: string,
    password: string,
    errorCallback: (err: string) => void
  ) => void;
};

function LoginDialogue({ onSubmit }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [error, setError] = React.useState("");
  const onLogin = () => {
    onSubmit(email, password, (err) => {
      setError(err);
    });
  };
  return (
    <Card padding="medium" className={classes.root}>
      <div className={classes.logoContainer}>
        <Logo className={classes.logo} />
      </div>
      <h1 className={classes.heading}>Login to TTC Hub</h1>
      <p className={classes.subheading}>
        Glad to have you back, please enter your login details to proceed
      </p>
      <p className={classes.errMessage}>{error}</p>
      <FancyInput
        label="Email"
        labelColor={"#5CC301"}
        type={"email"}
        onChange={setEmail}
      />
      <FancyInput
        label="Password"
        labelColor={"#5CC301"}
        type={"password"}
        onChange={setPassword}
      />
      <FancyButton text="Login to TTC" onClick={onLogin} />
      <a className={classes.link} href="https://example.com">
        I don't have a TTC Hub account
      </a>
    </Card>
  );
}

export default LoginDialogue;
