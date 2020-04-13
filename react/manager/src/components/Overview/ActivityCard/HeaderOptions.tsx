import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
  },
  option: {
    fontSize: theme.fontSizes.default,
    fontWeight: 300,
    cursor: "pointer",
    color: theme.colors.primaryBlack,
    paddingLeft: "15px",
  },
  selected: {
    color: theme.colors.textBlue,
  },
  borderRight: {
    borderRight: `1px solid ${theme.colors.borderGrey}`,
    paddingRight: "15px",
  },
}));

type Props = {
  selected: string;
  options: Array<string>;
  onClick: Function;
  className?: string;
};

function HeaderOptions({ selected, options, onClick, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <span className={classNames(classes.root, className)}>
      {options.map((option, index) => (
        <span
          onClick={() => onClick(option)}
          className={classNames(
            classes.option,
            selected === option && classes.selected,
            index !== options.length - 1 && classes.borderRight
          )}
        >
          {option}
        </span>
      ))}
    </span>
  );
}

export default HeaderOptions;
