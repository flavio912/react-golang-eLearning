import * as React from "react";
import PageTitle from "components/PageTitle";
import Button from "sharedComponents/core/Button";
import { createUseStyles, useTheme } from "react-jss";
import { Theme } from "helpers/theme";

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

type Props = {
  title: string;
  subTitle: string;
  sideText?: string;
  sideComponent?: React.ReactNode;
  showCreateButtons: boolean;
};
const PageHeader = ({
  title,
  subTitle,
  showCreateButtons,
  sideText,
  sideComponent
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.infoHeader}>
      <PageTitle title={title} subTitle={subTitle} sideText={sideText} sideComponent={sideComponent} />
      {showCreateButtons && (
        <div className={classes.mainButtons}>
          <Button bold archetype="submit">
            Quick Booking
          </Button>
          <Button bold archetype="submit">
            Add Delegates
          </Button>
        </div>
      )}
    </div>
  );
};

export default PageHeader;
