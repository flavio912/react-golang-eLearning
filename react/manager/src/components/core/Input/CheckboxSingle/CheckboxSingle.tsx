import * as React from 'react';
import { createUseStyles } from 'react-jss';
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
  onChange?: (value: boolean) => void;
};

function CheckboxSingle({
  label = '',
  defaultChecked = false,
  onChange
}: Props) {
  const classes = useStyles();
  const [isChecked, setChecked] = React.useState(defaultChecked);

  const _onClick = () => {
    const newChecked = !isChecked;
    setChecked(newChecked);
    if (onChange) onChange(newChecked);
  };

  return (
    <div className={classes.container}>
      <div className={classes.checkbox} key={label} onClick={() => _onClick()}>
        <Icon
          name={isChecked ? 'FormCheckbox_Checked' : 'FormCheckbox_Unchecked'}
          size={15}
          pointer
        />
        <p className={classes.label}>{label}</p>
      </div>
    </div>
  );
}

export default CheckboxSingle;
