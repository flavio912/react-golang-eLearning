import * as React from 'react'
import { createUseStyles } from 'react-jss';
import CoreInput from '../../core/CoreInput';

const styles = createUseStyles({

})

type Props = {

}

function FancyInput(props: Props) {
  const classes = styles();

  const onChange = (text: string) => {

  }
  return (
    <div>
      <CoreInput placeholder="joe@blogs.com" type="email" onChange={onChange} />
    </div>
  )
}

export default FancyInput