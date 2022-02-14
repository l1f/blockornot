import React, { FormEvent, useContext, useEffect, useState } from "react";

import { AuthContext } from "../../context/context";
import { dispatchCompleteAuth, dispatchFetchInitialAuthData } from "../../context/actions";
import { useNavigate } from "react-router-dom";

import Button from "@mui/material/Button";
import { Alert, Card, CardContent, Grid, TextField, Typography } from "@mui/material";
import Header from "../../components/Header/header.component";
import { AuthStage } from "../../entities/AuthStage";
import openInNewTab from "../../helper/open-new-tab.helper";
import Footer from "../../components/Footer/footer.component";

const Login = () => {
  const navigate = useNavigate();

  const [authStage, setAuthStage] = useState<AuthStage>(AuthStage.initial);
  const [pin, setPin] = useState("");

  const { state, dispatch } = useContext(AuthContext);

  useEffect(() => {
    dispatchFetchInitialAuthData(dispatch);
  }, [dispatch]);

  const handleAuthStart = async (event: FormEvent) => {
    event.preventDefault();

    const url = state.accessData?.accessUrl;
    openInNewTab(`${url?.Scheme}://${url?.Host}${url?.Path}?${url?.RawQuery}`);

    setAuthStage(AuthStage.userInput);
  };

  const handlePinSubmit = async (event: FormEvent) => {
    event.preventDefault();

    await dispatchCompleteAuth(pin, state.accessData!.headers!, dispatch);
    setAuthStage(AuthStage.done);
  };

  const handlePinChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPin(event.currentTarget.value);
  };

  const LoginBtn = () => <Button variant="contained" fullWidth onClick={handleAuthStart}>Authorize</Button>;
  const PinField = () => (
    <div>
      <form>
        <TextField label="PIN" variant="outlined" fullWidth onChange={handlePinChange} value={pin}/>
        <Button sx={{ marginTop: 1 }} variant="contained" fullWidth onClick={handlePinSubmit}>Confirm</Button>
      </form>
    </div>
  );


  if (authStage === AuthStage.done) {
    navigate("/");
  }

  if (state.loading) {
    return <div>Loading...</div>;
  }

  return (
    <div style={{ minHeight: "100vh", display: "flex", flexDirection: "column", justifyContent: "space-between" }}>
      <Header/>
      <Grid container spacing={0} direction="column" alignItems="center" justifyContent="center">
        <Grid item xs={3}>
          <Alert severity="error" sx={{ marginBottom: 1, display: state.error ? "" : "none" }}>{state.error}</Alert>
          <Card>
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Login with twitter
              </Typography>

              {(() => {
                  switch (authStage) {
                    case AuthStage.initial:
                      return <LoginBtn/>;
                    case AuthStage.userInput:
                      return <PinField/>;
                  }
                }
              )()}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
      <Footer/>
    </div>
  );
};

export default Login;
