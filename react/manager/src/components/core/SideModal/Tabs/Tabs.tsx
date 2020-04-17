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
    display: "flex",
    flexDirection: "column",
  },
  pagebody: {
    flexGrow: 1,
    display: "flex",
    flexDirection: "column",
    padding: [30, 40],
  },
  footer: {
    height: 120,
    borderTop: `2px solid ${theme.colors.borderGrey}`,
    backgroundColor: theme.colors.backgroundGrey,
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    padding: [0, 40],
  },
  smallHeading: {
    fontSize: theme.fontSizes.smallHeading,
    fontWeight: 900,
  },
  largeText: {
    fontSize: 15,
    maxWidth: 500,
  },
  text: {
    fontSize: theme.fontSizes.default,
  },
  termsBox: {
    width: "100%",
    display: "flex",
    flexDirection: "column",
    alignItems: "stretch",
    border: `1px solid ${theme.colors.borderGrey}`,
    borderRadius: [0, 0, 3, 3],
    margin: [theme.spacing(1), 0],
  },
  termsTitle: {
    backgroundColor: theme.colors.primaryBlue,
    color: "white",
    fontSize: theme.fontSizes.default,
    fontWeight: 900,
    padding: [10, 15],
  },
  termsContent: {
    maxHeight: 180,
    padding: [0, theme.spacing(1)],
    overflowY: "auto",
  },
  course: {
    height: 40,
    padding: [0, 15],
    border: `1px solid ${theme.colors.borderGrey}`,
    backgroundColor: "white",
    fontSize: theme.fontSizes.default,
    borderRadius: 4,
    display: "flex",
    alignItems: "center",
    margin: [theme.spacing(1), 0, theme.spacing(2)],
  },
  currentTotal: {
    fontSize: 15,
    fontWeight: "bold",
  },
  money: {
    color: "#1081AA",
    marginLeft: theme.spacing(1),
  },
}));

type ComponentProps = {
  state: any;
  setState: (newState: any) => void;
  setTab: (tab: string) => void;
  closeModal: () => void;
};

export type TabContent = {
  key: string;
  component: ({
    state,
    setState,
    setTab,
    closeModal,
  }: ComponentProps) => JSX.Element;
};

type Props = {
  content: TabContent[];
  initialState: any;
  closeModal: () => void;
};

function Tabs({ content, closeModal, initialState }: Props) {
  const classes = useStyles();
  const [state, setState] = React.useState<any>(initialState);
  const [active, setActive] = React.useState<string>(content[0].key || "");
  const [ref, { width }] = useDimensions();
  const [{ right }, set] = useSpring(() => ({
    right: 0,
    config: config.default,
  }));

  const throwErr = () => {
    throw new Error("Invalid Tab error");
  };

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
        {content.map(({ key, component: Page }: TabContent) => (
          <div key={key} className={classes.content} style={{ width }}>
            <Page
              state={state}
              setState={setState}
              setTab={(tab: string) =>
                content.map((c) => c.key).includes(tab)
                  ? setActive(tab)
                  : throwErr()
              }
              closeModal={closeModal}
            />
          </div>
        ))}
      </animated.div>
    </div>
  );
}

type SimpleProps = {
  children: React.ReactNode;
};

export function Body({ children }: SimpleProps) {
  const classes = useStyles();
  return <div className={classes.pagebody}>{children}</div>;
}

export function Footer({ children }: SimpleProps) {
  const classes = useStyles();
  return <div className={classes.footer}>{children}</div>;
}

export function Heading({ children }: SimpleProps) {
  const classes = useStyles();
  return <h1 className={classes.smallHeading}>{children}</h1>;
}

export function LargeText({ children }: SimpleProps) {
  const classes = useStyles();
  return <p className={classes.largeText}>{children}</p>;
}

export function Text({ children }: SimpleProps) {
  const classes = useStyles();
  return <p className={classes.text}>{children}</p>;
}

export function TermsBox({
  children,
  title,
}: {
  children: React.ReactNode;
  title: string;
}) {
  const classes = useStyles();
  return (
    <div className={classes.termsBox}>
      <div className={classes.termsTitle}>{title}</div>
      <div className={classes.termsContent}>{children}</div>
    </div>
  );
}

export function CourseHighlight({ children }: SimpleProps) {
  const classes = useStyles();
  return <p className={classes.course}>{children}</p>;
}

export function CurrentTotal({ total }: { total?: number }) {
  const classes = useStyles();
  return total ? (
    <p className={classes.currentTotal}>
      Current total:
      <span className={classes.money}>£{total.toFixed(2)}</span>
    </p>
  ) : (
    <div />
  );
}

export default Tabs;
