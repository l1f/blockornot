import React, { FormEvent, useContext, useEffect, useState } from "react";

import { AuthContext } from "../../context/context";
import {
  dispatchCompleteAuth,
  dispatchFetchInitialAuthData,
} from "../../context/actions";
import { useNavigate } from "react-router-dom";

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
  const navigate = useNavigate();

  const [authStage, setAuthStage] = useState<stage>(stage.initial);
  const [pin, setPin] = useState("");

  const { state, dispatch } = useContext(AuthContext);

  useEffect(() => {
    dispatchFetchInitialAuthData(dispatch);
  }, [dispatch]);

  const handleAuthStart = async (event: FormEvent) => {
    event.preventDefault();

    const url = state.accessData?.accessUrl;
    openInNewTab(`${url?.Scheme}://${url?.Host}${url?.Path}?${url?.RawQuery}`);

    setAuthStage(stage.userInput);
  };

  const handlePinSubmit = async (event: FormEvent) => {
    event.preventDefault();

    await dispatchCompleteAuth(pin, state.accessData!.headers!, dispatch);
    setAuthStage(stage.done);
  };

  const handlePinChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPin(event.currentTarget.value);
  };

  const LoginBtn = () => <button onClick={handleAuthStart}>login</button>;
  const PinField = () => (
    <div>
      <form>
        <label htmlFor="pin">PIN: </label>
        <input id="pin" onChange={handlePinChange} value={pin} />
        <button onClick={handlePinSubmit}>login</button>
      </form>
    </div>
  );

  if (authStage === stage.done) {
    navigate("/");
  }

  if (state.loading) {
    return <div>Loading..</div>;
  }

  return (
    <div>
      {state.error}
      <div>
        <h1>Login w/ twitter</h1>
        {authStage === stage.initial ? <LoginBtn /> : null}
        {authStage === stage.userInput ? <PinField /> : null}
      </div>
    </div>
  );
};

export default Login;
