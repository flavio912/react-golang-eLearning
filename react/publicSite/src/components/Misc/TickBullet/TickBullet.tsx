import * as React from "react";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
    root: {
        display: 'flex'
    },
    text: {
        fontSize: theme.fontSizes.large,
        fontweight: '500',
        color: theme.colors.secondaryBlack,
        marginLeft: '25px'
    }
}));

type Props = {
    text: string;
    size?: number;
    className?: string;
};

function TickBullet({ text, size = 13, className }: Props) {
    const theme = useTheme()
    const classes = useStyles({ theme });
  return (
      <div className={classNames(classes.root, className)}>
          <Icon name="Tick" size={size} />
          <div className={classes.text}>{text}</div>
      </div>
  );
}

export default TickBullet;