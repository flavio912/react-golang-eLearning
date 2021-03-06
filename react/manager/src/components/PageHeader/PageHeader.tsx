import * as React from 'react';
import PageTitle from 'components/PageTitle';
import Button from 'sharedComponents/core/Input/Button';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import DelegateSlideIn, {
  DelegateInfo
} from 'components/Delegate/DelegateSlideIn';
import MultiUser from 'components/MultiUser';
import SingleUser from 'components/SingleUser';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex'
  },
  infoHeader: {
    display: 'flex',
    justifyContent: 'space-between',
    marginBottom: theme.spacing(2)
  },
  mainButtons: {
    display: 'inline-grid',
    gridTemplateColumns: '1fr 1fr',
    gridGap: theme.spacing(2)
  }
}));

type Props = {
  title: string;
  subTitle: string;
  sideText?: string;
  sideComponent?: React.ReactNode;
  showCreateButtons: boolean;
  backProps?: {
    text: string;
    onClick: Function;
  };
};
const PageHeader = ({
  title,
  subTitle,
  showCreateButtons,
  sideText,
  sideComponent,
  backProps
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [openDelegateSlideIn, setOpenDelegateSlideIn] = React.useState(false);
  const [isMultiUserOpen, setMultiUserOpen] = React.useState(false);
  const [isSingleUserOpen, setSingleUserOpen] = React.useState(false);
  return (
    <div className={classes.infoHeader}>
      <PageTitle
        title={title}
        subTitle={subTitle}
        sideText={sideText}
        sideComponent={sideComponent}
        backProps={backProps}
      />
      {showCreateButtons && (
        <div className={classes.mainButtons}>
          <Button
            bold
            archetype="submit"
            onClick={() => setMultiUserOpen(true)}
          >
            Quick Booking
          </Button>
          <Button
            bold
            archetype="submit"
            onClick={() => {
              setOpenDelegateSlideIn(true);
            }}
          >
            Add Delegates
          </Button>
        </div>
      )}
      <DelegateSlideIn
        isOpen={openDelegateSlideIn}
        onClose={() => setOpenDelegateSlideIn(false)}
      />
      <MultiUser
        isOpen={isMultiUserOpen}
        onClose={() => setMultiUserOpen(false)}
      />
      <SingleUser
        user={{firstName: 'John', uuid: '0000'}}
        isOpen={isSingleUserOpen}
        onClose={() => setSingleUserOpen(false)}
      />
    </div>
  );
};

export default PageHeader;
