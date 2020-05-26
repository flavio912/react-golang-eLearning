import * as React from "react";
import PageTitle from "components/PageTitle";
import Button from "sharedComponents/core/Button";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";
import DelegateSlideIn from "components/DelegateSlideIn";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
  },
  infoHeader: {
    display: "flex",
    justifyContent: "space-between",
    marginBottom: theme.spacing(2),
  },
  mainButtons: {
    display: "inline-grid",
    gridTemplateColumns: "1fr 1fr",
    gridGap: theme.spacing(2),
  },
}));

const addDelegate = (
  profileUrl: string,
  firstName: string,
  lastName: string,
  jobTitle: string,
  email: string,
  phone: string,
  ttcId: string,
  errorCallback: (err : string) => void
) => {
  console.log("added a new delegate");
}

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
  backProps,
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const [openDelegateSlideIn, setOpenDelegateSlideIn] = React.useState(false);

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
          <Button bold archetype="submit">
            Quick Booking
          </Button>
          <Button bold archetype="submit" onClick={(_) => {
            setOpenDelegateSlideIn(true);
          }}>
            Add Delegates
          </Button>
        </div>
      )}
      <DelegateSlideIn isOpen={openDelegateSlideIn} onClose={() => setOpenDelegateSlideIn(false)} submitDelegate={addDelegate}/>
    </div>
  );
};

export default PageHeader;
