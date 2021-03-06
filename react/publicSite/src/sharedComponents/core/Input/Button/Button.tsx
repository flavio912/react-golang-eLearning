import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon, { IconNames } from '../../Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  button: {
    borderRadius: theme.buttonBorderRadius,
    padding: [0, theme.spacing(2)],
    border: '1px solid',
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    cursor: 'pointer',
    height: 40,
    justifyContent: 'center',
    fontSize: theme.fontSizes.default,
    transition: '0.1s ease',
    transitionProperty: 'border-colour, background-color',
    outline: 'none',
    // it would be nice to have a light blue hover state
    '&:focus': {
      borderColor: theme.colors.primaryBlue,
    },
    '&::-moz-focus-inner': {
      border: 0,
    },
  },
  bold: {
    fontWeight: 'bold !important',
  },
  small: {
    padding: [0, theme.spacing(1)],
  },
  default: {
    color: theme.colors.primaryBlack,
    borderColor: theme.colors.borderGrey,
    backgroundColor: 'white',
    fontWeight: 200,
    borderRadius: 4,
  },
  grey: {
    color: theme.colors.primaryBlack,
    borderColor: theme.colors.borderGrey,
    backgroundColor: theme.colors.backgroundGrey,
  },
  submit: {
    color: 'white',
    borderColor: theme.colors.primaryBlue,
    backgroundColor: theme.colors.primaryBlue,
    '&:focus': {
      borderColor: '#0044db',
      backgroundColor: '#0044db',
    },
  },
  gradient: {
    color: 'white',
    backgroundImage: `linear-gradient(45deg,
      ${theme.colors.primaryBlue}, ${theme.colors.primaryGreen})`,
    fontWeight: 800,
    fontSize: theme.fontSizes.large,
    borderRadius: 4,
    border: 'none',
  },
  disabled: {
    opacity: 0.5,
    cursor: 'initial !important',
  },
}));

export type Archetypes = 'default' | 'grey' | 'submit' | 'gradient';

interface Props {
  archetype?: Archetypes;
  icon?: { left?: IconNames; right?: IconNames };
  iconSize?: number;
  bold?: boolean;
  small?: boolean;
  className?: string;
  noIconPadding?: boolean;
  disabled?: Boolean;
  onClick?: () => void;
}

function Button({
  archetype,
  icon,
  iconSize,
  noIconPadding,
  bold,
  small,
  children,
  className = '',
  disabled,
  onClick,
  ...props
}: Props & React.PropsWithoutRef<JSX.IntrinsicElements['button']>) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <button
      className={classNames(
        disabled && classes.disabled,
        classes.button,
        classes[archetype || 'default'],
        bold && classes.bold,
        small && classes.small,
        className,
      )}
      onClick={() => {
        if (!disabled && onClick) onClick();
      }}
      {...props}
    >
      {/* replace with actual icon */}
      {/* prop should also be a string (icon name) */}
      {icon?.left && (
        <Icon
          style={noIconPadding ? {} : { marginRight: small ? 5 : 10 }}
          name={icon.left}
          size={iconSize ? iconSize : small ? 12 : 15}
        />
      )}
      {children}
      {icon?.right && (
        <Icon
          style={noIconPadding ? {} : { marginLeft: small ? 5 : 10 }}
          name={icon.right}
          size={iconSize ? iconSize : small ? 12 : 15}
        />
      )}
    </button>
  );
}

export default Button;
