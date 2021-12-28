import { FormEvent, useContext, useEffect, useState } from "react";

import { AuthContext } from "../../context/context";
import { getAuthHeaders } from "../../context/actions";

const Login = () => {
  const { state, dispatch } = useContext(AuthContext);

  useEffect(() => {
    getAuthHeaders(dispatch);
  }, [getAuthHeaders, dispatch]);

  const handleLogin = async (event: FormEvent) => {
    event.preventDefault();
    try {
      console.log("test..");
    } catch (error) {
      console.log("error.. :D", error);
    }
  };

  return (
    <div>
      {state.error}
      <div>
        <h1>Login w/ twitter</h1>
        <button onClick={handleLogin}>login</button>
      </div>
    </div>
  );
};

export default Login;
