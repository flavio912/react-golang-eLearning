import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
// @ts-ignore
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

export type ComponentProps = {
  state: any;
  setState: (newState: any) => void;
  setTab: (tab: string) => void;
};

export type TabContent = {
  key: string;
  component: ({ state, setState, setTab }: ComponentProps) => JSX.Element;
};

type Props = {
  content: TabContent[];
};

function Tabs({ content }: Props) {
  const classes = useStyles();
  const [state, setState] = React.useState<any>();
  const [active, setActive] = React.useState<string>(content[0].key || "");
  const [ref, { width }] = useDimensions();

  return (
    <div className={classes.container}>
      <div className={classes.heading}>
        {content.map(({ key }) => (
          <p key={key} style={key === active ? { color: "green" } : undefined}>
            {key}
          </p>
        ))}
      </div>
      <div className={classes.body} ref={ref}>
        {content.map(({ key, component: Page }) => (
          <div key={key} className={classes.content} style={{ width }}>
            <Page
              state={state}
              setState={setState}
              setTab={(tab: string) => setActive(tab)}
            />
          </div>
        ))}
      </div>
    </div>
  );
}

export default Tabs;
