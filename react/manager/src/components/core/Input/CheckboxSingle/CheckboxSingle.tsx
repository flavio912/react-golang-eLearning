import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import Icon from "../../../../sharedComponents/core/Icon";

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
    maxWidth: 500,
  },
}));

type Checkbox = {
  label: string;
  checked: boolean;
};

type Props = {
  box: Checkbox;
  setBox: (value: Checkbox) => void;
};

function CheckboxSingle({ box, setBox }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <div
        className={classes.checkbox}
        key={box.label}
        onClick={() =>
          setBox({
            ...box,
            checked: !box.checked,
          })
        }
      >
        <Icon
          name={box.checked ? "FormCheckbox_Checked" : "FormCheckbox_Unchecked"}
          size={15}
          pointer
        />
        <p className={classes.label}>{box.label}</p>
      </div>
    </div>
  );
}

export default CheckboxSingle;
