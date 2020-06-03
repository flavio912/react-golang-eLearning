import * as React from 'react';
import { createUseStyles } from 'react-jss';
import classNames from 'classnames';

import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    height: 40,
    position: 'relative',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'strech',
    background: theme.colors.primaryWhite
  },
  clickableBox: {
    height: 40,
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: [0, 15],
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: theme.buttonBorderRadius,
    cursor: 'pointer'
  },
  dropdown: {
    position: 'absolute',
    top: 50,
    width: '100%',
    maxHeight: 450, // don't know how this should be calculated
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: theme.primaryBorderRadius,
    display: 'flex',
    flexDirection: 'column',
    backgroundColor: theme.colors.primaryWhite
  },
  searchResults: {
    overflowY: 'auto'
  },
  title: {
    fontSize: theme.fontSizes.small,
    fontWeight: 300,
    marginRight: 19
  },
  option: {
    height: 40,
    cursor: 'pointer',
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: [0, 15],
    fontSize: theme.fontSizes.default,
    borderBottom: `1px solid ${theme.colors.borderGrey}`,
    '&:last-child': {
      borderBottom: 'none'
    }
  },
  triangleDown: {
    width: 0,
    height: 0,
    borderLeft: '5px solid transparent',
    borderRight: '5px solid transparent',
    borderTop: ` ${5 * 2 * 0.866}px solid ${theme.colors.primaryBlue}`
  },
  triangleUp: {
    width: 0,
    height: 0,
    borderLeft: '5px solid transparent',
    borderRight: '5px solid transparent',
    borderBottom: `${5 * 2 * 0.866}px solid ${theme.colors.primaryBlue}`
  }
}));

export type DropdownOption = {
  id: number;
  title: string;
  component?: React.ReactNode;
};

type Props = {
  placeholder: string;
  options: DropdownOption[];
  selected?: DropdownOption;
  setSelected: (
    selected: ((current: DropdownOption) => DropdownOption) | DropdownOption
  ) => void;
  className?: string;
};

function Dropdown({
  placeholder,
  options,
  selected,
  setSelected,
  className
}: Props) {
  const classes = useStyles();
  const [isOpen, setOpen] = React.useState<boolean>(false);

  return (
    <>
      <div className={classNames(classes.container, className)}>
        <div
          className={classes.clickableBox}
          onClick={() => setOpen((o) => !o)}
        >
          <p className={classes.title}>
            {selected ? selected.title : placeholder}
          </p>
          <div className={isOpen ? classes.triangleUp : classes.triangleDown} />
        </div>
        {isOpen && (
          <div className={classes.dropdown}>
            <div className={classes.searchResults}>
              {options.map(({ id, title, component }: DropdownOption) => (
                <div
                  key={id}
                  className={classes.option}
                  onClick={() => {
                    setSelected({ id, title, component });
                    setOpen(false);
                  }}
                >
                  <span className={classes.title}>
                    {component ? component : title}
                  </span>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </>
  );
}

export default Dropdown;
