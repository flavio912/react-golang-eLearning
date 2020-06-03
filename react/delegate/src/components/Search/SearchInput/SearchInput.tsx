import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import CoreInput from 'sharedComponents/core/Input/CoreInput';

const useStyles = createUseStyles((theme: Theme) => ({
  searchInputRoot: {
    backgroundColor: theme.colors.primaryWhite
  },
  searchText: {
    '& input': {
      lineHeight: `25px`,
      letterSpacing: -0.5,
      color: theme.colors.primaryBlack,
      paddingBottom: 30,
      fontSize: theme.fontSizes.tinyHeading,
      fontWeight: 'bold',
      width: '100%'
    },
    borderBottom: `1px solid ${theme.colors.borderGrey}`
  }
}));
type Props = {
  value?: string;
  onChange: (text: string) => string | Promise<void> | void;
  placeholder?: string;
};

function SearchInput({ onChange, value, placeholder }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.searchInputRoot}>
      <div className={classes.searchText}>
        <CoreInput
          type="text"
          value={value}
          placeholder={placeholder}
          onChange={onChange}
        />
      </div>
    </div>
  );
}

export default SearchInput;
