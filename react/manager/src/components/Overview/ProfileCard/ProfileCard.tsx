import * as React from "react";
import Card, { PaddingOptions } from "../../core/Card";
import InfoField, { Field } from "../../core/InfoField";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    width: 340,
  },
  heading: {
    fontSize: theme.fontSizes.default,
    fontWeight: 300,
    color: theme.colors.primaryBlack,
    paddingBottom: "15px",
  },
}));

type Props = {
  heading: string;
  fields: Array<Field>;
  padding?: PaddingOptions;
  className?: string;
};

function ProfileCard({ heading, fields, padding = "none", className }: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <Card padding={padding} className={classNames(classes.root, className)}>
      <div className={classNames(classes.heading)}>{heading}</div>
      {fields.map((field) => (
        <InfoField
          fieldName={field.fieldName}
          value={field.value}
          padding={field.padding}
        />
      ))}
    </Card>
  );
}

export default ProfileCard;
