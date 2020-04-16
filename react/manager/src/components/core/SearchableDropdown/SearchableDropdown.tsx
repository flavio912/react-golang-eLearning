import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    width: '100%',
    height: 40,
    position: 'relative'
  },
  clickableBox: {
    width: '100%',
    height: 40,
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: [0, theme.spacing(3)],
  },
  dropdown: {
    position: 'absolute',
    top: 50,
    width: '100%',
    height: 450, // don't know how this should be calculated
  }
}));

type Props = {
  placeholder: string;
};

function SearchableDropdown({ placeholder }: Props) {
  const classes = useStyles();
  const [selected, setSelected] = React.useState<string | null>(null);
  const [isOpen, setOpen] = React.useState<boolean>(false);

  return (
    <>
      <div className={classes.clickableBox}>
        <p>{selected || placeholder}</p>
        {/* Icon goes here */}
      </div>
      {isOpen && (
        <div className={classes.dropdown}>

        </div>
      )}
    </>
  );
}

export default SearchableDropdown;
