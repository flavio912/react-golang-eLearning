import * as React from "react";
import Card, { PaddingOptions } from "../../../sharedComponents/core/Cards/Card";
import InfoField, { Field } from "../../core/Input/InfoField";
import { createUseStyles, useTheme } from "react-jss";
import classNames from "classnames";
import { Theme } from "helpers/theme";
import Icon from "sharedComponents/core/Icon";

const useStyles = createUseStyles((theme: Theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
  },
  heading: {
    fontSize: theme.fontSizes.default,
    fontWeight: 300,
    color: theme.colors.primaryBlack,
  },
  row: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    paddingBottom: "15px",
  },
}));

type Props = {
  heading: string;
  fields: Field[];
  onClick?: Function;
  padding?: PaddingOptions;
  className?: string;
};

function ProfileCard({
  heading,
  fields,
  onClick,
  padding = "none",
  className,
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ theme });

  return (
    <Card padding={padding} className={classNames(classes.root, className)}>
      <div className={classNames(classes.row)}>
        <div className={classNames(classes.heading)}>{heading}</div>
        <Icon
          name="Card_SecondaryActon_Dots"
          size={20}
          style={{ cursor: "pointer" }}
          onClick={() => onClick && onClick()}
        />
      </div>
      {fields &&
        fields.map((field, index) => (
          <InfoField
            key={field.fieldName}
            fieldName={field.fieldName}
            value={field.value}
            padding={field.padding}
            border={index + 1 !== fields.length}
          />
        ))}
    </Card>
  );
}

export default ProfileCard;
