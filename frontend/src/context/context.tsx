import React, { createContext, Dispatch, useReducer } from "react";

import { AuthReducer, initialStateType } from "./reducer";

const initialState = {
  account: undefined,
  accessData: null,

  loading: true,
  error: null,
};

const AuthContext = createContext<{
  state: initialStateType;
  dispatch: Dispatch<any>;
}>({
  state: initialState,
  dispatch: () => null,
});

const AuthProvider: React.FunctionComponent = ({ children }) => {
  const [state, dispatch] = useReducer(AuthReducer, initialState);

  return (
    <AuthContext.Provider value={{ state, dispatch }}>
      {children}
    </AuthContext.Provider>
  );
};

export { AuthContext, AuthProvider };
