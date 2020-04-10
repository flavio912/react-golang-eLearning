import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
// @ts-ignore
import useDimensions from "react-use-dimensions";
import { useSpring, animated, config } from "react-spring";
import classNames from "classnames";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    width: "100%",
    height: "100%",
    overflow: "hidden",
    display: "flex",
    flexDirection: "column",
  },
  heading: {
    width: "calc(100% - 30px)",
    height: 70,
    borderBottom: `2px solid ${theme.colors.borderGrey}`,
    display: "flex",
    flexDirection: "row",
    justifyContent: "flex-start",
    alignItems: "flex-end",
    paddingLeft: 30,
  },
  tabName: {
    margin: [0, 0, 0, 10],
    padding: [5, 10],
    position: "relative",
    bottom: -2,
    cursor: "pointer",
    borderBottom: "2px solid transparent",
    color: theme.colors.textGrey,
    fontSize: theme.fontSizes.large,
  },
  active: {
    color: "black",
    borderBottomColor: theme.colors.primaryGreen,
  },
  body: {
    flexGrow: 1,
    display: "flex",
    flexDirection: "row",
    position: "relative",
  },
  content: {
    flexShrink: 0,
    height: "100%",
  },
}));

type ComponentProps = {
  state: any;
  setState: (newState: any) => void;
  setTab: (tab: string) => void;
  closeModal?: () => void;
};

export type TabContent = {
  key: string;
  component: ({ state, setState, setTab }: ComponentProps) => JSX.Element;
};

type Props = {
  content: TabContent[];
  closeModal?: () => void;
};

function Tabs({ content, closeModal }: Props) {
  const classes = useStyles();
  const [state, setState] = React.useState<any>();
  const [active, setActive] = React.useState<string>(content[0].key || "");
  const [{ right }, set] = useSpring(() => ({
    right: 0,
    config: config.default,
  }));
  const [ref, { width }] = useDimensions();

  React.useEffect(() => {
    content.forEach(({ key }, index) => {
      if (active === key) {
        set({ right: index * width });
      }
    });
  }, [content, active, width, set]);

  return (
    <div className={classes.container}>
      <div className={classes.heading}>
        {content.map(({ key }) => (
          <p
            key={key}
            className={classNames(
              classes.tabName,
              key === active && classes.active
            )}
            onClick={() => setActive(key)}
          >
            {key}
          </p>
        ))}
      </div>
      <animated.div className={classes.body} ref={ref} style={{ right }}>
        {content.map(({ key, component: Page }) => (
          <div key={key} className={classes.content} style={{ width }}>
            <Page
              state={state}
              setState={setState}
              setTab={(tab: string) => setActive(tab)}
              closeModal={closeModal}
            />
          </div>
        ))}
      </animated.div>
    </div>
  );
}

export default Tabs;
