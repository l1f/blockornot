import React, { createContext, Dispatch, useReducer } from "react";

import { AuthReducer, initialStateType } from "./reducer";
import { AuthActions } from "./actions";

const initialState = {
  account: undefined,
  header: null,

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

//
// const AuthStateContext = createContext(initialState);
// const AuthDispatchContext = createContext(AuthReducer);
//
// export function useAuthState() {
//   const context = useContext(AuthStateContext);
//   if (context === undefined) {
//     throw new Error("useAuthState must be used within a AuthProvider");
//   }
//
//   return context;
// }
//
// export function useAuthDispatch(): Dispatch {
//   const context = useContext(AuthDispatchContext);
//   if (context === undefined) {
//     throw new Error("useAuthDispatch must be used within a AuthProvider");
//   }
//
//   return context;
// }
//
// type AuthProviderProps = {};
// export const AuthProvider: FunctionComponent<AuthProviderProps> = ({
//   children,
// }) => {
//   const [state, dispatch] = useReducer(AuthReducer, initialState);
//
//   return (
//     // @ts-ignore
//     <AuthStateContext.Provider value={state}>
//       <AuthDispatchContext.Provider value={dispatch}>
//         {children}
//       </AuthDispatchContext.Provider>
//     </AuthStateContext.Provider>
//   );
// };
