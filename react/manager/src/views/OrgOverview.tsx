import * as React from "react";
import HeaderMenu from "components/Menu/HeaderMenu";
import SideMenu from "components/Menu/SideMenu";
import { Tab } from "components/Menu/SideMenu/SideMenu";
import ActvityCard from "components/Overview/ActivityCard";
import { createUseStyles, useTheme } from "react-jss";
import classes from "*.module.css";
import classNames from "classnames";
import Button from "components/core/Button";
import UserSearch from "components/UserSearch";
import PageTitle from "components/PageTitle";
import TitleWrapper from "components/Overview/TitleWrapper";
import QuickInfo from "components/Overview/QuickInfo";
import QuickOverview from "components/Overview/QuickOverview";
import TrainingProgressCard from "components/Overview/TrainingProgressCard";
import { Theme } from "helpers/theme";
import ProfileCard from "components/Overview/ProfileCard";

type Props = {};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    flexGrow: 1,
    maxWidth: 1200,
  },
  activity: {},
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
  statsRow: {
    display: "flex",
    justifyContent: "space-between",
    marginBottom: theme.spacing(2),
  },
  cardFlex: {
    display: "flex",
  },
  breaker: {
    width: theme.spacing(2),
  },
  searchRow: {
    marginBottom: theme.spacing(2),
  },
  bottomRow: {
    display: "flex",
    justifyContent: "space-between",
    marginBottom: theme.spacing(2),
  },
}));

export const OrgOverview = () => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.root}>
      <div className={classes.infoHeader}>
        <PageTitle title="Fedex" subTitle="Organisation Overview" />
        <div className={classes.mainButtons}>
          <Button bold archetype="submit">
            Quick Booking
          </Button>
          <Button bold archetype="submit">
            Add Delegates
          </Button>
        </div>
      </div>
      <div className={classes.searchRow}>
        <UserSearch
          companyName="TESTcompany"
          searchFunction={async (query: string) => {
            return [];
          }}
        />
      </div>
      <div className={classes.statsRow}>
        <TitleWrapper title="Quick Overview">
          <QuickOverview
            purchasedCourses={20}
            numDelegates={130}
            numValidCertificates={10}
            numCertificatesExpiringSoon={15}
          />
        </TitleWrapper>
        <TitleWrapper title="Training Progress">
          <div className={classes.cardFlex}>
            <TrainingProgressCard
              coursesDone={20}
              timeTracked={{ h: 20, m: 15 }}
              title="Weekly"
            />
            <div className={classes.breaker} />
            <TrainingProgressCard
              coursesDone={20}
              timeTracked={{ h: 20, m: 15 }}
              title="Weekly"
            />
          </div>
        </TitleWrapper>
      </div>
      <div className={classes.bottomRow}>
        <TitleWrapper title="Your information">
          <ProfileCard
            heading="Profile"
            fields={[
              { fieldName: "Name", value: "Fred Eccleston" },
              { fieldName: "Role", value: "Group Leader" },
              { fieldName: "Email", value: "Group Leader" },
              { fieldName: "Tel Contact", value: "Group Leader" },
              { fieldName: "Active since", value: "Group Leader" },
            ]}
            padding="medium"
          />
        </TitleWrapper>
        <TitleWrapper title={"Activity"}>
          <ActvityCard
            className={classes.activity}
            padding={"medium"}
            leftHeading={"Delegates activity"}
            rightHeading={"Recent Updates"}
            options={["This month", "All Time"]}
            updates={[]}
            data={[
              { name: "name", value: 10 },
              { name: "time", value: 15 },
            ]}
          />
        </TitleWrapper>
      </div>
    </div>
  );
};
