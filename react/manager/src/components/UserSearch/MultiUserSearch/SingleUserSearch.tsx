import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import UserResults from '../UserResults';
import CoreInput from 'components/core/Input/CoreInput';
import { ResultItem } from '../UserSearch';
import IdentTag from 'sharedComponents/IdentTag';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    border: [1, 'solid', theme.colors.borderGrey],
    borderRadius: theme.primaryBorderRadius,
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    background: 'white'
  },
  rootFocused: {
    extend: 'root',
    boxShadow: theme.shadows.primary
  },
  search: {
    display: 'flex',
    alignItems: 'center',
    flex: 1
  },
  searchInput: {
    fontSize: theme.fontSizes.default,
    fontWeight: 200
  },
  openDelegate: {
    background: theme.colors.primaryBlue,
    color: 'white',
    alignItems: 'center',
    padding: [0, 38],
    display: 'flex',
    fontWeight: 600,
    fontSize: theme.fontSizes.default,
    borderRadius: [0, theme.primaryBorderRadius, theme.primaryBorderRadius, 0],
    cursor: 'pointer'
  },
  buttonCont: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: theme.spacing(1)
  }
}));

type Props = {
  searchFunction: (query: string) => Promise<ResultItem[]>;
  debounceTime?: number;
  selection: ResultItem | undefined;
  setSelection: (user: ResultItem | undefined) => void;
};

function SingleUserSearch({
  searchFunction,
  selection,
  setSelection,
  debounceTime = 600
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [loading, setLoading]: [boolean, any] = React.useState(false);
  const [focus, setFocus]: [boolean, any] = React.useState(false);
  const [results, setResults]: [ResultItem[], any] = React.useState([]);
  const [input, setInput]: [string, any] = React.useState('');
  const [debouncer, setDebouncer]: [number | undefined, any] = React.useState();

  const onFocus = () => {
    setFocus(true);
    onChange('');

    // Reset the input when focusing again
    if (selection !== undefined) {
      setInput('');
      setSelection(undefined);
    }
  };
  const onBlur = () => {
    setFocus(false);
  };
  const onChange = (text: string) => {
    setLoading(true);
    clearTimeout(debouncer);
    const timeout = setTimeout(async () => {
      if (text.length === 0 && results.length > 0) {
        setLoading(false);
        return;
      }
      const res = await searchFunction(text);
      setLoading(false);
      setResults(res);
    }, debounceTime);
    setDebouncer(timeout);
  };
  const onSelect = (result: ResultItem) => {
    setFocus(false);
    setSelection(result);
    setInput(result.key);
  };

  return (
    <div className={focus ? classes.rootFocused : classes.root}>
      <div className={classes.buttonCont}>
        <div className={classes.search}>
          <CoreInput
            type="search"
            placeholder="e.g. Paul Smith"
            onChange={onChange}
            className={classes.searchInput}
            onFocus={onFocus}
            onBlur={onBlur}
            value={input}
            setValue={setInput}
          />
        </div>
        {selection && <IdentTag ident={selection.value} />}
      </div>
      {focus && (
        <UserResults loading={loading} results={results} onSelect={onSelect} />
      )}
    </div>
  );
}
export default SingleUserSearch;
