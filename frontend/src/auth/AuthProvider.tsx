export {};
// import React, {
//   FunctionComponent,
//   useContext,
//   useEffect,
//   useState,
// } from "react";
// import { Client, getClient } from "../api/client";
// import { getAuthHeaders, getInitialAuthHeaders } from "../api/endpoints/auth";
// import account from "../api/entities/Account";
//
// type AuthState = {
//   loggedIn: boolean;
//   account: account | null;
//
//   login: (pin: string) => Promise<account>;
//   logout: () => void;
//
//   client: Client;
// };
//
// type AuthProviderProps = {};
//
// const emptyAuthState = {
//   loggedIn: false,
//   account: null,
//
//   login: async (pin: string): Promise<account> => {
//     throw "not implemented";
//   },
//   logout: () => {},
//   client: getClient(),
// };
//
// const AuthContext = React.createContext<AuthState>(emptyAuthState);
//
// export const useAuth = (): AuthState => {
//   return useContext(AuthContext);
// };
//
// const AuthProvider: FunctionComponent<AuthProviderProps> = ({ children }) => {
//   const [loading, setLoading] = useState<boolean>(true);
//   const [authState, setAuthState] = useState<AuthState>({
//     ...emptyAuthState,
//     login: async (pin): Promise<account> => {
//       const header = await getAuthHeaders(authState.client, pin);
//     },
//   });
//
//   const initAuth = async () => {
//     const header = await getInitialAuthHeaders(authState.client);
//     const c = getClient(header);
//     setAuthState({ ...authState, client: c });
//     setLoading(false);
//   };
//
//   useEffect(() => {
//     initAuth();
//   }, []);
//
//   if (loading) {
//     return (
//       <div>Loading...</div>
//       // <Backdrop open={true}>
//       //   <CircularProgress color="inherit" />
//       // </Backdrop>
//     );
//   } else {
//     return (
//       <AuthContext.Provider value={authState}>{children}</AuthContext.Provider>
//     );
//   }
// };
//
// export default AuthProvider;
// @ts-ignore
// @ts-ignore
// @ts-ignore
