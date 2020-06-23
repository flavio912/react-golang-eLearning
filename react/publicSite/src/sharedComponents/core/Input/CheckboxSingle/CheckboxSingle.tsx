import * as React from 'react';
import { createUseStyles } from 'react-jss';
import classNames from 'classnames';
import { Theme } from 'helpers/theme';
import Icon from '../../../../sharedComponents/core/Icon';

const useStyles = createUseStyles((theme: Theme) => ({
  container: {
    display: 'flex',
    flexDirection: 'column'
  },
  checkbox: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    cursor: 'pointer'
  },
  label: {
    margin: theme.spacing(1),
    fontSize: theme.fontSizes.small,
    maxWidth: 500,
    color: theme.colors.primaryBlack
  }
}));

type Props = {
  label?: string;
  defaultChecked?: boolean;
  size?: number;
  onChange?: (value: boolean) => void;
  fontStyle?: string;
  className?: string;
};

function CheckboxSingle({
  label = '',
  defaultChecked = false,
  size = 15,
  onChange,
  fontStyle,
  className
}: Props) {
  const classes = useStyles();
  const [isChecked, setChecked] = React.useState(defaultChecked);

  const _onClick = () => {
    const newChecked = !isChecked;
    setChecked(newChecked);
    if (onChange) onChange(newChecked);
  };

  return (
    <div className={classNames(classes.container, className)}>
      <div className={classNames(classes.checkbox, className)} key={label} onClick={() => _onClick()}>
        <Icon
          name={isChecked ? 'FormCheckbox_Checked' : 'FormCheckbox_Unchecked'}
          size={size}
          pointer
        />
        <p className={classNames(classes.label, fontStyle)}>{label}</p>
      </div>
    </div>
  );
}

export default CheckboxSingle;
