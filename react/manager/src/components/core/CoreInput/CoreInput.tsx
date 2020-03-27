import React, {useState, ChangeEvent} from 'react'
import { createUseStyles } from 'react-jss';
import classNames from 'classnames';

const styles = createUseStyles({
  defaultStyles: {
    border: 'none',
    outline: 'none'
  }

  
})

export type InputTypes = 'text' | 'password' | 'email' | 'search' | 'tel' | 'number'
type Props = {
  type: InputTypes
  className?: string
  placeholder?: string
  onChange: (text: string) => string | void, // Function given changed text and should return what to update it to 
}

function CoreInput({type, className, placeholder = '', onChange}: Props) {
  const classes = styles();
  const [value, setValue] = useState('');
  
  const updateInput = (event: ChangeEvent<HTMLInputElement>) => {
    const newValue = onChange(event.target.value) || event.target.value;
    setValue(newValue)
  }

  return (
    <input type={type} className={classNames(className, classes.defaultStyles)} value={value} placeholder={placeholder} onChange={updateInput}/>
  )
}

export default CoreInput