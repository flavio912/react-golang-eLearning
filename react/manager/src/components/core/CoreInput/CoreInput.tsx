import React, {useState, ChangeEvent} from 'react'
import { createUseStyles } from 'react-jss';

const styles = createUseStyles({

})

type Props = {
  type: 'text' | 'password' | 'email' | 'search' | 'tel' | 'number'
  className?: string
  placeholder?: string
  onChange: (text: string) => string | void, // Function given changed text and should return what to update it to 
}

function CoreInput({type, className, placeholder = '', onChange}: Props) {
  const classes = styles();
  const [value, setValue] = useState(placeholder);
  
  const updateInput = (event: ChangeEvent<HTMLInputElement>) => {
    const newValue = onChange(event.target.value) || event.target.value;
    setValue(newValue)
  }

  return (
    <input type={type} className={className} value={value} onChange={updateInput}/>
  )
}

export default CoreInput