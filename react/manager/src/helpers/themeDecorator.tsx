import * as React from "react"
import { ThemeProvider } from 'react-jss'
import theme from './theme'

const ThemeDecorator = (storyFn: any) => (
  <ThemeProvider theme={theme}>{storyFn()}</ThemeProvider>
)

export default ThemeDecorator;