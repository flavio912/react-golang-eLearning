import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import moment from 'moment';
import { Theme } from 'helpers/theme';
import ProfileIcon from 'components/core/ProfileIcon';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'flex-start',
    alignItems: 'center',
    margin: '10px 0px'
  },
  icon: {
    width: '30px',
    height: '30px',
    borderRadius: '20px',
    marginRight: '15px',
    backgroundColor: theme.colors.primaryGreen
  },
  heading: {
    fontSize: theme.fontSizes.default,
    fontWeight: 400,
    color: theme.colors.primaryBlack,
    width: '275px',
    whiteSpace: 'nowrap',
    overflow: 'hidden',
    textOverflow: 'ellipsis'
  },
  time: {
    fontSize: theme.fontSizes.tiny,
    fontWeight: 300,
    color: theme.colors.textGrey,
  },
  borderLeft: {
    borderLeft: `1px solid ${theme.colors.borderGrey}`
  },
  borderRight: {
    borderRight: `1px solid ${theme.colors.borderGrey}`
  }
}));

export interface Update {
    name: string;
    course: string;
    time: Date;
  }

type Props = {
  name: string;
  course: string;
  time: Date;
  className?: string;
};

function UserUpdate({ name, course, time, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <div className={classNames(classes.root, className)}>
        <ProfileIcon className={classNames(classes.icon)} name={name} size={30} />
        <div>
            <div className={classNames(classes.heading)}>{`${name} took a course in ${course}`}</div>
            <div className={classNames(classes.time)}>{moment(time.toISOString()).fromNow()}</div>
        </div>
    </div>
  );
}

export default UserUpdate;
