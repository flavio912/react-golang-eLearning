import * as React from "react";
import { createUseStyles } from "react-jss";
import { Theme } from "helpers/theme";
import ProfileIcon from "components/core/ProfileIcon";

const useStyles = createUseStyles((theme: Theme) => ({
  userLabelRoot: {
    display: "flex",
    alignItems: "center",
  },
  name: {
    marginLeft: theme.spacing(1),
    fontSize: theme.fontSizes.default,
    fontWeight: 300,
  },
}));

type Props = {
  name: string;
  profileUrl: string | undefined;
};

function UserLabel({ name, profileUrl }: Props) {
  const classes = useStyles();
  return (
    <div className={classes.userLabelRoot}>
      <ProfileIcon url={profileUrl} name={name} size={30} />
      <p className={classes.name}>{name}</p>
    </div>
  );
}

export default UserLabel;
