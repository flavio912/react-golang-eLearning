import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import { Theme } from 'helpers/theme';
import classnames from 'classnames';
import CoreInput, { InputTypes } from '../CoreInput';

const useStyles = createUseStyles((theme: Theme) => ({
  container:{
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'flex-start'
  },
  label:{
    flex: 1,
    fontSize: theme.fontSizes.default,
    color: theme.colors.primaryBlack,
    marginBottom: theme.spacing(1)
  },
  input:{
    flex: 2,
    border: `1px solid ${theme.colors.borderBlack}`,
    borderRadius: 5,
    padding: '15px 10px'
  }
}));

type Props = {
  label: string;
  labelClassName?: string;
  placeholder?: string;
  type?: InputTypes;
  onChange?: (text: string) => string | void;
};

const EasyInput = ({
  label,
  labelClassName,
  placeholder = '',
  type = 'text',
  onChange = () => {}
}: Props) => {
  const theme = useTheme();
  const classes = useStyles({ theme });

  const onTextChange = (text : string) => {
    return onChange(text);
  };

  return (
    <div className={classes.container}>
      <span className={classnames(classes.label, labelClassName)}>{label}</span>
      <CoreInput
        placeholder={placeholder}
        type={type}
        onChange={onTextChange}
        className={classes.input}
      />
    </div>
  );
};

export default EasyInput;