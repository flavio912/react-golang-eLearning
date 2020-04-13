import * as React from "react";
import Card, { PaddingOptions } from "../../core/Card";
import Button from "../../core/Button";
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
    width: '778px',
    fontSize: theme.fontSizes.default,
  },
  root: {
    display: "flex",
    flexDirection: "column",
  },
  heading: {
    fontWeight: '300',
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
  name: string;
  value: number;
}

type Props = {
  leftHeading: string;
  rightHeading: string;
  options: Array<string>;
  updates: Array<Update>;
  data: Array<Statistic>;
  padding?: PaddingOptions;
  className?: string;
};

function ActvityCard({
  leftHeading,
  rightHeading,
  options,
  updates,
  data,
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

        <div className={classNames(classes.graph)}>
          <Graph
            heading="Active Delegates"
            outerValue={data[0].value}
            innerValue={data[1].value}
          />
        </div>

        <div className={classNames(classes.row)}>
          <StatCircle
            heading={data[0].name}
            value={data[0].value}
            color="green"
            border="right"
          />
          <StatCircle
            heading={data[1].name}
            value={data[1].value}
            color="red"
            border="right"
          />
          <StatCircle
            heading="Total"
            value={data[0].value + data[1].value}
            color="grey"
          />
        </div>

        <div className={classNames(classes.button)}>
          <Button archetype="submit">View all Delegates</Button>
        </div>
      </div>
      <div className={classes.dividerHolder}>
        <div className={classes.divider} />
      </div>
      <div className={classes.root}>
        <div className={classNames(classes.heading, classes.marginBottom)}>
          {rightHeading}
        </div>

        {updates.slice(0,7).map((update) => (
          <UserUpdate
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
