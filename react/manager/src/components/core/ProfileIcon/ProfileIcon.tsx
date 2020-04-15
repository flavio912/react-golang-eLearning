import * as React from "react";
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    backgroundRepeat: 'no-repeat',
    backgroundSize: 'cover',
    opacity: 0.5
  },
  initials: {
    color: theme.colors.primaryBlack,
    fontWeight: 900,
    fontSize: theme.fontSizes.tiny
  }
}));

type Props = {
  name?: string;
  url?: string;
  size?: number;
  className?: string;
};

const colourMap = {
  1: '#cf1d1d',
  2: '#cf4a1d',
  3: '#cf1d76',
  4: '#1d47cf',
  5:'#611dcf',
  6:'#1dcfa6',
  7:'#1fdd76',
  8:'#1f4fd1',
  9:'#37d11f',
  10:'#60d422',
  11:'#22d457',
  12:'#b9d422',
  13:'#d69524',
  14:'#d6c124',
  15:'#d66824',
  16:'#d94b27',
  17:'#d97727',
  18:'#d92739',
  19:'#c49e47',
  20:'#6249c7',
  21:'#a149c7',
  22:'#4995c7',
  23:'#cc4e93',
  24:'#cc4e54',
  25:'#c64ecc',
  26:'#1ad10a',
}

/**
 * Converts a name into two letters
 * @param name The name to convert
 */
function toInitials(name: string) {
  if (name.length > 2) {
    let initials: RegExpMatchArray = name.match(/\b\w/g) || [];
    return ((initials.shift() || '') + (initials.pop() || '')).toUpperCase();
  }
  return name.toUpperCase();
}

/**
 * Given a pair of initals it will return a unique
 * colour. It will also catch 1 letter initials and
 * defaults to a colour when no string is given
 * @param initials The intitals to use as the hash
 */
function initialToColour(initials: string) {
  if (initials.length > 1) {
    const first = initials.charCodeAt(0) - 64;
    const second = initials.charCodeAt(1) - 64;
    const hash = Math.ceil((first + second) / 2);
    return colourMap[hash];
  } else if (initials.length === 1) {
    const hash = initials.charCodeAt(0) - 64;
    return colourMap[hash];
  }
  return colourMap[0];
}

function ProfileIcon({ name = "", url, size = 30, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const initials: string = toInitials(name);
  const colour: string = initialToColour(initials);

  return (
    <div
      className={classNames(classes.root, className)}
      style={{ backgroundColor: colour, backgroundImage: `url(${url})`, width: `${size}px`, height: `${size}px`, borderRadius: `${size}px` }}
    >
      <div className={classNames(classes.initials)}>{initials}</div>
    </div>
  );
}

export default ProfileIcon;
