import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import classNames from 'classnames';
import moment from 'moment';
import { Theme } from 'helpers/theme';
import ProfileIcon from 'sharedComponents/core/ProfileIcon';
import { createFragmentContainer, graphql } from 'react-relay';
import { UserUpdate_activity } from './__generated__/UserUpdate_activity.graphql';

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'flex-start',
    alignItems: 'center',
    margin: '13px 0px'
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
    fontWeight: 300,
    color: theme.colors.primaryBlack,
    width: '275px',
    whiteSpace: 'nowrap',
    overflow: 'hidden',
    textOverflow: 'ellipsis'
  },
  time: {
    fontSize: theme.fontSizes.tiny,
    fontWeight: 300,
    color: theme.colors.textGrey
  },
  borderLeft: {
    borderLeft: `1px solid ${theme.colors.borderGrey}`
  },
  borderRight: {
    borderRight: `1px solid ${theme.colors.borderGrey}`
  }
}));

type Props = {
  activity: UserUpdate_activity;
  className?: string;
};

function UserUpdate({ activity, className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  let description = '';
  switch (activity.type) {
    case 'activated':
      description = `Account created for ${activity.user?.firstName} ${activity.user?.lastName}`;
    case 'completedCourse':
      description = `${activity.user?.firstName} ${activity.user?.lastName} completed ${activity.course?.name}`;
    case 'failedCourse':
      description = `${activity.user?.firstName} ${activity.user?.lastName} failed ${activity.course?.name}`;
    case 'newCourse':
      description = `${activity.user?.firstName} ${activity.user?.lastName} started ${activity.course?.name}`;
  }
  return (
    <div className={classNames(classes.root, className)}>
      <ProfileIcon
        className={classNames(classes.icon)}
        name={`${activity.user?.firstName} ${activity.user?.lastName}`}
        url={activity.user?.profileImageUrl ?? undefined}
        size={30}
      />
      <div>
        <div className={classNames(classes.heading)}>{description}</div>
        <div className={classNames(classes.time)}>
          {moment(activity.createdAt).fromNow()}
        </div>
      </div>
    </div>
  );
}

export default createFragmentContainer(UserUpdate, {
  activity: graphql`
    fragment UserUpdate_activity on Activity {
      type
      createdAt
      course {
        name
      }
      user {
        firstName
        lastName
        profileImageUrl
      }
    }
  `
});
