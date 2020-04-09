import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import useDimensions from "react-use-dimensions";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    width: "100%",
    height: "100%",
    overflow: "hidden",
    display: "flex",
    flexDirection: "column",
  },
  heading: {
    width: "100%",
    height: 70,
    borderBottom: `2px solid ${theme.colors.borderGrey}`,
    display: "flex",
    flexDirection: "row",
    justifyContent: "flex-start",
    alignItems: "flex-end",
  },
  body: {
    flexGrow: 1,
    display: "flex",
    flexDirection: "row",
  },
  content: {
    flexShrink: 0,
  },
}));

export type TabContent = [
  {
    key: string;
    component: typeof React.Component;
  }
];

type Props = {
  children: TabContent;
};

function Tabs({ children }: Props) {
  const classes = useStyles();
  const [ref, { width }] = useDimensions();

  return (
    <div className={classes.container}>
      <div className={classes.heading}>
        {children.map(({ key }) => (
          <p key={key}>{key}</p>
        ))}
      </div>
      <div className={classes.body} ref={ref}>
        {children.map(({ key, component: Content }) => (
          <div key={key} className={classes.content} style={{ width }}>
            <Content />
          </div>
        ))}
      </div>
    </div>
  );
}

export default Tabs;
