import * as React from "react";
import { createUseStyles } from "react-jss";
import Icon, { IconNames } from "sharedComponents/core/Icon";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: "inline-flex",
    flexDirection: "column",    
    alignItems: "center",
  },
  text: {
    color: theme.colors.primaryBlack,
    fontSize: theme.fontSizes.extraLarge,
    fontWeight: "bold",
    lineHeight: "41px",    
    margin: 0,
    letterSpacing: "-0.45px",
  },
  value: {
    position: "relative",
    fontWeight: 800,
    color: theme.colors.primaryBlack,
    fontSize: "60px",
    margin: "0 0 15px 0",
    lineHeight: "46px",
    letterSpacing: "2.95px"
  },
  valueArrow: {
    position: "absolute",
    transform: "translateY(-50%)",
    left: "calc(100% + 3px)",
    top: "50%",
    width: 0,
    height: 0, 
    borderLeft: "5px solid transparent",
    borderRight: "5px solid transparent",    
  },
  up: {
    borderBottom: `10px solid ${theme.colors.secondaryGreen}`,  
  },
  down: {
    borderTop: `10px solid ${theme.colors.secondaryDanger}`,  
  },
  footer: {
    display: "flex",
    alignItems: "flex-start",
    height: "41px",
    lineHeight: "20px",
    fontSize: theme.fontSizes.default,
    color: theme.colors.textBlue,
    margin: 0,
    letterSpacing: "-0.35px",
  }
}));

type Props = {
  icon: IconNames;
  text: string;
  value: number | { h: number; m: number };
  footer?: string;
  valueArrow?: string;
};

function QuickInfo({ icon, text, value, footer, valueArrow }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.container}>
      <Icon name={icon} size={25} />
      <p className={classes.text}>{text}</p>
      {typeof value === "number" ? (        
        <p className={classes.value}>
          {value}
          {(valueArrow?.toLowerCase() === 'up' || valueArrow?.toLowerCase() === 'down') && <span className={`${classes.valueArrow} ${classes[valueArrow?.toLowerCase()]}`}></span>}          
        </p>
      ) : (
        <p className={classes.value} style={{lineHeight: "37px", fontSize: "38px", marginTop: "7px", marginBottom:"18px"}}>
          {value.h}
          <span>h </span>
          {value.m}
          <span>m</span>
        </p>
      )}
      <p className={classes.footer}>{footer}</p>
    </div>
  );
}

export default QuickInfo;
