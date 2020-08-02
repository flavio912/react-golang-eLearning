import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { ResultItem } from './UserSearch';
import IdentTag from 'sharedComponents/IdentTag';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    padding: [theme.spacing(0), theme.spacing(1)],
    borderRadius: theme.primaryBorderRadius,
    '&:hover': {
      background: theme.colors.searchHoverGrey
    },
    fontSize: theme.fontSizes.default,
    fontWeight: 300,
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    cursor: 'pointer'
  }
}));

type Props = {
  result: ResultItem;
  onClick: (result: ResultItem) => void;
};

function UserResult({ result, onClick }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div
      className={classes.root}
      onMouseDown={() => {
        onClick(result);
      }}
    >
      {result.key}
      <IdentTag ident={result.value} />
    </div>
  );
}

export default UserResult;
