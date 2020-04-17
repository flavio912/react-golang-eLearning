import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import Icon from "../Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: "flex",
    flexDirection: "column",
  },
  checkbox: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    cursor: "pointer",
  },
  label: {
    margin: theme.spacing(1),
    fontSize: theme.fontSizes.small,
  },
}));

type Checkbox = {
  label: string;
  checked: boolean;
};

type Props = {
  children: Checkbox[];
  setBoxes: (values: Checkbox[]) => void;
};

function Checkbox({ children, setBoxes }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      {children.map(({ label, checked }: Checkbox, index: number) => (
        <div
          className={classes.checkbox}
          key={label}
          onClick={() =>
            setBoxes(
              children.reduce(
                (arr: Checkbox[], curr: Checkbox, i: number) => [
                  ...arr,
                  i === index ? { label, checked: !curr.checked } : curr,
                ],
                []
              )
            )
          }
        >
          <Icon
            name={checked ? "FormCheckbox_Checked" : "FormCheckbox_Unchecked"}
            size={15}
            pointer
          />
          <p className={classes.label}>{label}</p>
        </div>
      ))}
    </div>
  );
}

export default Checkbox;
