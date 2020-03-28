import * as React from 'react';
import { createUseStyles, useTheme } from 'react-jss';
import CoreInput, { InputTypes } from '../../core/CoreInput';
import classnames from 'classnames';

const useStyles = createUseStyles((theme: any) => ({
  root: {
    position: 'relative',
    display: 'flex',
    marginTop: theme.spacing(2)
  },
  fancyInput: {
    fontSize: 14,
    border: '1px solid #08080814',
    borderRadius: 5,
    padding: '15px 10px'
  },
  label: {
    position: 'absolute',
    background: 'white',
    left: 12,
    top: '-7px',
    fontSize: 13,
    color: (props: any) => props.labelColor,
    padding: '0px 14px'
  }
}));

type Props = {
  label: string; // Text for the label
  labelColor?: string;
  labelClassName?: string;
  placeholder?: string;
  type?: InputTypes;
  onChange?: (text: string) => string | void;
};

function FancyInput({
  label,
  labelColor = 'black',
  labelClassName,
  placeholder = '',
  type = 'text',
  onChange = () => {}
}: Props) {
  const theme = useTheme();
  const classes = useStyles({ labelColor, theme });

  const onTextChange = (text: string) => {
    return onChange(text);
  };

  return (
    <div className={classes.root}>
      <span className={classnames(classes.label, labelClassName)}>{label}</span>
      <CoreInput
        placeholder={placeholder}
        type={type}
        onChange={onTextChange}
        className={classes.fancyInput}
      />
    </div>
  );
}

export default FancyInput;
