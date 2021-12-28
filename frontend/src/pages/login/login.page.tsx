import { FormEvent, useContext, useEffect, useState } from "react";

import { AuthContext } from "../../context/context";
import { dispatchFetchInitialAuthData } from "../../context/actions";

const openInNewTab = (url: string) => {
  const newWindow = window.open(url, "_blank", "noopener,noreferrer");
  if (newWindow) newWindow.opener = null;
};

enum stage {
  initial,
  userInput,
  done,
}

const Login = () => {
  const [authStage, setAuthStage] = useState<stage>(stage.initial);
  const { state, dispatch } = useContext(AuthContext);

  useEffect(() => {
    dispatchFetchInitialAuthData(dispatch);
  }, [dispatch]);

  const handleLogin = async (event: FormEvent) => {
    event.preventDefault();

    const url = state.accessData?.accessUrl;

    openInNewTab(`${url?.Scheme}://${url?.Host}${url?.Path}?${url?.RawQuery}`);
    setAuthStage(stage.userInput);
  };

  if (authStage == stage.userInput) {
    return (
      <div>
        <label htmlFor="pin">PIN: </label>
        <input id="pin" />
      </div>
    );
  }

  return (
    <div>
      {state.error}
      {state.loading ? <div>Loading..</div> : null}
      <div>
        <h1>Login w/ twitter</h1>
        {state.accessData?.accessUrl?.Path}
        <button onClick={handleLogin}>login</button>
      </div>
    </div>
  );
};

export default Login;
