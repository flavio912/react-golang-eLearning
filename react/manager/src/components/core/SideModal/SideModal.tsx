import * as React from "react";
import { createUseStyles } from "react-jss";
import { animated, useSpring, config } from "react-spring";
import { Theme } from "helpers/theme";
import Icon from "../../../sharedComponents/core/Icon";

const modalWidth = 750;

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    width: "100%",
    height: "100vh",
    overflowX: "hidden",
    position: "fixed",
    zIndex: 45,
    top: 0,
  },
  background: {
    zIndex: 50,
    width: "100%",
    height: "100%",
    backgroundColor: "rgba(7,67,121,0.75)",
  },
  modal: {
    zIndex: 55,
    width: "100%",
    maxWidth: modalWidth,
    height: "100%",
    position: "absolute",
    top: 0,
    display: "flex",
    flexDirection: "column",
    alignItems: "stretch",
  },
  header: {
    height: 76,
    background: theme.primaryGradient,
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    padding: [0, 40],
  },
  title: {
    color: "white",
    fontWeight: "bold",
    fontSize: theme.fontSizes.smallHeading,
  },
  body: {
    backgroundColor: "white",
    flexGrow: 1,
  },
}));

type Props = {
  title: string;
  isOpen: boolean;
  closeModal(): void;
  children?: React.ReactNode;
};

// isOpen and closeModal should be provided from a useState hook in a parent component

function SideModal({ title, isOpen, closeModal, children }: Props) {
  const classes = useStyles();
  const { right, opacity, pointerEvents } = useSpring({
    right: isOpen ? 0 : -modalWidth,
    opacity: isOpen ? 1 : 0,
    pointerEvents: isOpen ? "auto" : "none",
    config: config.default,
  });

  return (
    <animated.div className={classes.container} style={{ pointerEvents }}>
      <animated.div className={classes.background} style={{ opacity }} />
      <animated.div className={classes.modal} style={{ right }}>
        <div className={classes.header}>
          <h2 className={classes.title}>{title}</h2>
          <Icon
            name="CloseCourseManagementTray_X"
            onClick={closeModal}
            pointer
          />
        </div>
        <div className={classes.body}>{children}</div>
      </animated.div>
    </animated.div>
  );
}

export default SideModal;
