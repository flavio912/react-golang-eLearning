import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { animated, useSpring, config } from 'react-spring';
import { Theme } from 'helpers/theme';
import Icon, { IconNames } from '../../../sharedComponents/core/Icon/Icon';
import Spacer from 'sharedComponents/core/Spacers/Spacer';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    zIndex: 100,
    top: 0,
    overflow: 'hidden',
    width: 274,
    gridRow: '1 / span 2',
    position: 'relative',
    height: '100%',
    flexDirection: 'column',
    alignItems: 'center',
    boxShadow: '1px 0px 0px rgba(0,0,0,0.03)',
    background: theme.colors.primaryWhite
  },
  menu: {
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'flex-start',
    alignItems: 'center'
  },
  tab: {
    width: '274px',
    height: 72,
    cursor: 'pointer',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'flex-start',
    opacity: 0.3,
    transition: 'background-color 0.3s linear, opacity 0.3s linear'
  },
  selected: {
    backgroundColor: theme.colors.hoverGreen,
    opacity: 1,
    transition: 'background-color 0.3s linear, opacity 0.3s linear'
  },
  fold: {
    width: 17,
    height: 40,
    position: 'absolute',
    left: '-13px',
    opacity: 1,
    transition: 'opacity 1s linear',
    borderRadius: 13,
    backgroundColor: theme.colors.secondaryGreen
  },
  noFold: {
    height: '40px',
    width: '5px',
    opacity: 0,
    transition: 'visibility 0s 1s, opacity 1s linear'
  },
  logo: {
    cursor: 'pointer',
    height: '50px',
    width: '140px',
    margin: '25px 0'
  },
  arrow: {},
  arrowCont: {
    top: 30,
    right: -15,
    width: 20,
    height: 20,
    padding: 5,
    position: 'absolute',
    boxShadow: '6px 2px 10px rgba(0,0,0,0.07)',
    borderRadius: 20,
    backgroundColor: '#FFFFFF'
  },
  title: {
    fontSize: theme.fontSizes.large,
    fontWeight: '700',
    color: theme.colors.textBlue,
    margin: '0 25px'
  },
  body: {
    padding: '30px 30px',
    backgroundColor: theme.colors.backgroundGrey,
    boxShadow: theme.shadows.body,
    flexGrow: 1
  },
  row: {
    display: 'flex',
    flexDirection: 'row',
    marginLeft: '30px'
  }
}));

export interface Tab {
  id: number;
  icon: IconNames;
  title: string;
  size?: number;
}

type Props = {
  logo: string;
  tabs: Array<Tab>;
  selected: Tab;
  onClick?: (tab: Tab) => void;
  onLogoClick?: Function;
  className?: string;
};

function SideMenu({
  logo,
  tabs,
  selected,
  onClick,
  onLogoClick,
  className
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });
  const [isOpen, setIsOpen] = React.useState(true);
  const { left, opacity } = useSpring({
    left: isOpen ? 0 : -274,
    opacity: isOpen ? 1 : 0,
    config: config.default
  });

  return (
    <animated.div
      className={classNames(classes.root, className)}
      style={{ left }}
    >
      {/* <div className={classes.arrowCont}>
        <Icon
          name={isOpen ? "ArrowLeft" : "ArrowRight"}
          className={classes.arrow}
          size={15}
          onClick={() => setIsOpen(!isOpen)}
        />
      </div> */}

      <img
        className={classNames(classes.logo)}
        onClick={() => onLogoClick && onLogoClick()}
        src={logo}
        alt="Logo"
      />
      <Spacer vertical spacing={2} />
      <Spacer vertical spacing={1} />
      <animated.div className={classNames(classes.menu)} style={{ opacity }}>
        {tabs &&
          tabs.map((tab) => (
            <div
              key={tab.id}
              className={classNames(
                classes.tab,
                selected && selected.id === tab.id && classes.selected
              )}
              onClick={() => {
                if (onClick) onClick(tab);
              }}
            >
              <div
                className={classNames(
                  selected && selected.id === tab.id
                    ? classes.fold
                    : classes.noFold
                )}
              />
              <div className={classes.row}>
                <Icon
                  name={tab.icon}
                  size={tab.size ?? 20}
                  style={{ cursor: 'pointer' }}
                />
                <div className={classes.title}>{tab.title}</div>
              </div>
              <div />
            </div>
          ))}
      </animated.div>
    </animated.div>
  );
}

export default SideMenu;
