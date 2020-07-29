import * as React from 'react';
import Icon from 'sharedComponents/core/Icon';

type State = {
  hasError: boolean;
};

type Props = {};

class ErrorBoundary extends React.Component<Props, State> {
  constructor(props: {}) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: any) {
    // Update state so the next render will show the fallback UI.
    return { hasError: true };
  }

  componentDidCatch(error: any, errorInfo: any) {
    // You can also log the error to an error reporting service
  }

  render() {
    if (this.state.hasError) {
      // You can render any custom fallback UI
      return (
        <div>
          <h1>Ah, sorry... Something's gone arwy - we're working on it.</h1>
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
