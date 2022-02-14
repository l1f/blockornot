import { useContext } from "react";
import { Navigate } from "react-router-dom";

import { AuthContext } from "../../context/context";

const RequireAuth = ({ children }: { children: JSX.Element }) => {
  const { state } = useContext(AuthContext);

  return state.isAuthenticated ? children : <Navigate to="/login" replace/>;
};

export default RequireAuth;
