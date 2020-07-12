import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { ResultItem } from './UserSearch';
import ThemeObject, { Theme } from 'helpers/theme';
import { ScaleLoader } from 'react-spinners';
import UserResult from './UserResult';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {},
  results: {
    padding: theme.spacing(1)
  },
  loader: {
    paddingTop: theme.spacing(1),
    display: 'flex',
    justifyContent: 'center'
  }
}));

type Props = {
  results: ResultItem[];
  loading: boolean;
  onSelect: (result: ResultItem) => void;
};

function UserResults({ results, loading, onSelect }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.results}>
      {!loading && (
        <div>
          {results.map((result) => (
            <UserResult result={result} onClick={onSelect} key={result.uuid} />
          ))}
        </div>
      )}
      {loading && (
        <div className={classes.loader}>
          <ScaleLoader
            color={ThemeObject.colors.borderGrey}
            height={20}
            loading={loading}
          />
        </div>
      )}
    </div>
  );
}

export default UserResults;
