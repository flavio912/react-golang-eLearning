import React from 'react';
import SingleUser from 'components/SingleUser';
import { Course } from 'sharedComponents/core/Input/SearchableDropdown/SearchableDropdown';

type State = {
  shown: boolean;
  courses: Course[];
};
type ActionType = 'show' | 'hide';
type Action = { type: ActionType; courses?: Course[] };
type Dispatch = (action: Action) => void;

const SideModalDispatchContext = React.createContext<Dispatch | undefined>(
  undefined
);

function sideModalReducer(state: State, action: Action) {
  switch (action.type) {
    case 'show': {
      return { shown: true, courses: action.courses };
    }
    case 'hide': {
      return { shown: false, courses: [] };
    }
  }
}

function SideModalProvider({ children }: { children: React.ReactNode }) {
  const [state, dispatch] = React.useReducer(sideModalReducer, {
    shown: false,
    courses: []
  });
  return (
    <SideModalDispatchContext.Provider value={dispatch}>
      <SingleUser
        isOpen={state.shown}
        onClose={() => {
          dispatch({ type: 'hide' });
        }}
        initialCourses={state.courses}
      />
      {children}
    </SideModalDispatchContext.Provider>
  );
}

function useSideModalDispatch() {
  const context = React.useContext(SideModalDispatchContext);
  if (context === undefined) {
    throw new Error('SideModalDispatch must be used with SideModalProvider');
  }
  return context;
}

export { SideModalProvider, useSideModalDispatch };
