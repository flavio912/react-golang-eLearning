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
import theme, { Theme } from "helpers/theme";
import ProfileCard from "components/Overview/ProfileCard";
import PageHeader from "components/PageHeader";
import Table from "components/core/Table";
import UserLabel from "components/core/Table/UserLabel";
import Text from "components/core/Table/Text/Text";
import Icon from "components/core/Icon";
import CourseCompletion from "components/core/Table/CourseCompletion";
import Dropdown from "components/core/Dropdown";
import Spacer from "components/core/Spacers/Spacer";
import Paginator from "components/Paginator";

type Props = {};

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    flexGrow: 1,
    maxWidth: 1275,
  },
  searchAndFilter: {
    display: "flex",
    justifyContent: "space-between",
    height: 40,
  },
  tableOptions: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
  },
  dropdown: {
    background: "white",
  },
  divider: {
    width: theme.spacing(1),
  },
  search: {
    flex: 0.4,
  },
}));

const delegateRow = (
  userUUID: string,
  name: string,
  profileUrl: string | undefined,
  email: string,
  coursesCompleted: number,
  totalCourses: number,
  lastActiveTimestamp: string,
  nextExpiryTimestamp: string
): any => ({
  key: userUUID,
  cells: [
    { component: () => <UserLabel name={name} profileUrl={profileUrl} /> },
    { component: () => <Text text={email} color={theme.colors.textBlue} /> },
    {
      component: () => (
        <CourseCompletion total={totalCourses} complete={coursesCompleted} />
      ),
    },
    { component: () => <Text text={lastActiveTimestamp} formatDate /> },
    { component: () => <Text text={nextExpiryTimestamp} formatDate /> },
    { component: () => <Icon name={"Card_SecondaryActon_Dots"} /> },
  ],
});

const DelegatesPage = () => {
  const theme = useTheme();
  const classes = useStyles({ theme });
  return (
    <div className={classes.root}>
      <PageHeader
        showCreateButtons
        title="Fedex"
        subTitle="Organisation Overview"
        sideText="127 members total"
      />
      <div className={classes.searchAndFilter}>
        <div className={classes.search}>
          <UserSearch
            companyName={"Fedex"}
            searchFunction={async (query: string) => {
              return [];
            }}
          />
        </div>
        <div className={classes.tableOptions}>
          <Button
            archetype={"default"}
            icon={{ right: "FilterAdjust" }}
            children={"Filters"}
          />
          <Spacer spacing={1} horizontal />
          <Dropdown
            placeholder={"Sort By"}
            options={[]}
            setSelected={(selected) => <div />}
            className={classes.dropdown}
          />
          <Spacer spacing={1} horizontal />
          <Button
            archetype={"default"}
            icon={{ right: "DownloadCSV" }}
            noIconPadding
          />
        </div>
      </div>
      <Spacer spacing={3} vertical />
      <Table
        header={[
          "Name",
          "Email",
          "Courses Completed",
          "Last Active",
          "Next Expiry",
          "Actions",
        ]}
        rows={[
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            3,
            6,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Tim Smith",
            "",
            "email@email.com",
            1,
            9,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Sam Smith",
            "",
            "email@email.com",
            7,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Bob Dylan",
            "",
            "email@email.com",
            10,
            20,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
          delegateRow(
            "asda",
            "Jim Smith",
            "",
            "email@email.com",
            2,
            7,
            "2013-04-20T20:00:00+0800",
            "2013-04-20T20:00:00+0800"
          ),
        ]}
      />
      <Spacer vertical spacing={3} />
      <Paginator
        currentPage={1}
        updatePage={() => {}}
        numPages={10}
        itemsPerPage={10}
      />
    </div>
  );
};

export default DelegatesPage;
