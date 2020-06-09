import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Card from '../Card';

const useStyles = createUseStyles((theme: Theme) => ({
  labelledCardRoot: {},
  labelHeader: {
    background: (props: { labelBackground: string }) =>
      props.labelBackground || theme.colors.primaryBlue,
    borderTopLeftRadius: theme.primaryBorderRadius,
    borderTopRightRadius: theme.primaryBorderRadius,
    display: 'flex',
    alignItems: 'center',
    padding: [4, theme.spacing(2)],
    fontWeight: 500,
    color: theme.colors.primaryWhite,
    fontSize: theme.fontSizes.small
  },
  content: {
    padding: theme.spacing(2)
  }
}));

export type PaddingOptions = 'none' | 'small' | 'medium' | 'large';

type Props = {
  children: React.ReactNode;
  padding?: PaddingOptions;
  label?: string;
  labelBackground?: string;
  className?: string;
};

function LabelledCard({
  children,
  padding = 'none',
  labelBackground,
  className,
  label = '',
  ...props
}: Props & React.PropsWithoutRef<JSX.IntrinsicElements['div']>) {
  const theme = useTheme();
  const classes = useStyles({ labelBackground, theme });

  return (
    <Card padding={'none'} className={classes.labelledCardRoot} {...props}>
      <div className={classes.labelHeader}>{label}</div>
      <div className={classes.content}>{children}</div>
    </Card>
  );
}

export default LabelledCard;
