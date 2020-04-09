import * as React from "react";
import { createUseStyles } from "react-jss";
import { animated, useSpring } from "react-spring";
import { Theme } from "helpers/theme";
import Icon from "../Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    width: "100%",
    height: "100vh",
    overflow: "hidden",
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
    maxWidth: 700,
    height: "100%",
    position: "absolute",
    top: 0,
    display: "flex",
    flexDirection: "column",
  },
  header: {
    width: `calc(100% - ${2 * 40})`,
    height: 76,
    background: theme.primaryGradient,
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    padding: [0, 40],
    color: "white",
    fontWeight: "bold",
    fontSize: theme.fontSizes.veryLarge,
  },
  body: {
    backgroundColor: "white",
    flexGrow: 1,
    width: "100%",
  },
  footer: {
    width: "100%",
    height: 120,
    backgroundColor: theme.colors.backgroundGrey,
    borderTop: `2px solid ${theme.colors.borderGrey}`,
  },
}));

type Props = {
  isOpen: boolean;
  closeModal: () => void;
};

function SideModal({ isOpen, closeModal, children }: Props) {
  const classes = useStyles();
  const { right, opacity } = useSpring({
    right: isOpen ? 0 : -700,
    opacity: isOpen ? 1 : 0,
  });

  return (
    <div className={classes.container}>
      <animated.div className={classes.background} style={{ opacity }} />
      <animated.div className={classes.modal} style={{ right }}>
        <div className={classes.header}>
          <h2>Course Management</h2>
          <Icon
            name="CloseCourseManagementTray_X"
            onClick={closeModal}
            pointer
          />
        </div>
        <div className={classes.body}>{/* Children go here */}</div>
        <div className={classes.footer}></div>
      </animated.div>
    </div>
  );
}

export default SideModal;
