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
    width: '282px',
    padding: '2px 2px',
    border: "1px solid",
    borderColor: theme.colors.borderGrey,
    borderRadius: theme.buttonBorderRadius,
    backgroundColor: 'rgba(255,255,255,0.5)',
  },
  option: {
    cursor: "pointer",
    width: '50%',
    padding: '8px 0',
    color: theme.colors.textGrey,
    backgroundColor: 'rgba(255,255,255,0.5)',
    opacity: 0.7,
    fontWeight: 600,
    border: "0px solid",
    borderRadius: theme.buttonBorderRadius,
    fontSize: theme.fontSizes.small,
    outline: 'none'
  },
  selected: {
    backgroundColor: theme.colors.primaryWhite,
    border: "1px solid",
    borderColor: theme.colors.borderGrey,
    boxShadow: theme.shadows.primary,
    opacity: 1
  },
}));

interface Props {
  selected: string;
  options: Array<string>;
  onClick: Function;
  className?: string;
}

function SelectButton({ selected, options, onClick, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
      {options && options.map(option => (
        <button
          key={option}
          onClick={() => onClick(option)}
          className={classNames(
            classes.option,
            selected === option && classes.selected
          )}
        >
          {option}
        </button>
      ))}
    </div>
    
  );
}

export default SelectButton;
