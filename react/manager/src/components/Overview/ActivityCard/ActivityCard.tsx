import * as React from "react";
import Card, { PaddingOptions } from "../../../sharedComponents/core/Cards/Card";
import Button from "../../../sharedComponents/core/Input/Button";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import StatCircle from "./StatCircle";
import HeaderOptions from "./HeaderOptions";
import Graph from "./Graph";
import UserUpdate, { Update } from "./UserUpdate";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: "grid",
    gridTemplateColumns: "1fr 50px 1fr",
    background: "white",
  },
  root: {
    display: "flex",
    flexDirection: "column",
  },
  heading: {
    fontSize: theme.fontSizes.small,
    color: theme.colors.primaryBlack,
  },
  row: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
  },
  graph: {
    display: "flex",
    alignSelf: "center",
    justifyContent: "center",
    margin: "30px 0 45px 0",
  },
  button: {
    alignSelf: "center",
    marginTop: "50px",
  },
  marginBottom: {
    margin: "0 0 15px 0",
  },
  divider: {
    width: 1,
    background: theme.colors.borderGrey,
  },
  dividerHolder: {
    display: "flex",
    justifyContent: "center",
  },
}));

export interface Statistic {
  innerRing: {
    name: string;
    value: number;
  };
  outerRing: {
    name: string;
    value: number;
  };
}

type Props = {
  leftHeading: string;
  rightHeading: string;
  options: string[];
  updates: Update[];
  data: Statistic;
  onClick?: Function;
  padding?: PaddingOptions;
  className?: string;
};

function ActvityCard({
  leftHeading,
  rightHeading,
  options,
  updates,
  data,
  onClick,
  padding = "none",
  className,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [currentOption, setCurrentOption] = React.useState(options[0]);
  const optionPressed = (option: string) => {
    setCurrentOption(option);
  };

  return (
    <Card
      className={classNames(classes.container, className)}
      padding={padding}
    >
      <div className={classes.root}>
        <div className={classNames(classes.row)}>
          <div className={classNames(classes.heading)}>{leftHeading}</div>
          <HeaderOptions
            selected={currentOption}
            options={options}
            onClick={optionPressed}
          />
        </div>

        <Graph
          className={classNames(classes.graph)}
          heading="Active Delegates"
          outerValue={data.outerRing.value}
          innerValue={data.innerRing.value}
        />

        <div className={classNames(classes.row)}>
          <StatCircle
            heading={data.outerRing.name}
            value={data.outerRing.value}
            color="green"
            border="right"
          />
          <StatCircle
            heading={data.innerRing.name}
            value={data.innerRing.value}
            color="red"
            border="right"
          />
          <StatCircle
            heading="Total"
            value={data.innerRing.value + data.outerRing.value}
            color="grey"
          />
        </div>

        <Button
          style={{ alignSelf: "center", marginTop: "50px" }}
          archetype="submit"
          onClick={() => onClick && onClick()}
        >
          View all Delegates
        </Button>
      </div>
      <div className={classes.dividerHolder}>
        <div className={classes.divider} />
      </div>
      <div className={classes.root}>
        <div className={classNames(classes.heading, classes.marginBottom)}>
          {rightHeading}
        </div>

        {updates &&
          updates
            .slice(0, 7)
            .map((update, i) => (
              <UserUpdate
                key={i}
                name={update.name}
                course={update.course}
                time={update.time}
              />
            ))}
      </div>
    </Card>
  );
}

export default ActvityCard;
