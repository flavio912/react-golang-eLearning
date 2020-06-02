import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import classnames from 'classnames';

const useStyles = createUseStyles((theme: Theme) => ({
  headingRoot: {
    fontWeight: (props: { font: Font }) => props.font.weight,
    fontSize: (props: { font: Font }) => props.font.size,
    color: theme.colors.primaryBlack,
    marginTop: 0
  }
}));

type Props = {
  text: string;
  size: 'large' | 'medium';
  className?: string;
};

type Font = {
  size: number;
  weight: number;
};

const fonts: {
  [key: string]: Font;
} = {
  large: {
    size: 41,
    weight: 900
  },
  medium: {
    size: 20,
    weight: 600
  }
};

function Heading({ text, size, className }: Props) {
  const theme = useTheme();

  const font = fonts[size];
  const classes = useStyles({ font, theme });

  return <h1 className={classnames(classes.headingRoot, className)}>{text}</h1>;
}

export default Heading;
