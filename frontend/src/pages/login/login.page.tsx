import React, { FormEvent, useContext, useEffect, useState } from 'react'

import { AuthContext } from '../../context/context'
import { dispatchCompleteAuth, dispatchFetchInitialAuthData } from '../../context/actions'
import { useNavigate } from 'react-router-dom'

import Button from '@mui/material/Button'
import { Alert, Card, CardContent, Grid, TextField, Typography } from '@mui/material'
import Header from '../../components/Header/header.component'

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

  const LoginBtn = () => <Button variant="contained" fullWidth onClick={handleAuthStart}>Authorize</Button>;
  const PinField = () => (
    <div>
      <form>
        <TextField label="PIN" variant="outlined" fullWidth onChange={handlePinChange} value={pin} />
        <Button sx={{ marginTop: 1 }} variant="contained" fullWidth onClick={handlePinSubmit}>Confirm</Button>
      </form>
    </div>
  );

  if (authStage === stage.done) {
    navigate("/");
  }

  if (state.loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <Header />
      <Grid container spacing={0} direction="column" alignItems="center" justifyContent="center" style={{ minHeight: '100vh'}}>
        <Grid item xs={3}>
          <Alert severity="error" sx={{ marginBottom: 1, display: state.error ? '' : 'none' }}>{state.error}</Alert>
          <Card>
            <CardContent>
              <Typography gutterBottom variant="h5" component="div">
                Login with twitter
              </Typography>

              {(() => {
                  switch (authStage) {
                    case stage.initial:
                      return <LoginBtn />;
                    case stage.userInput:
                      return <PinField />;
                  }
                }
              )()}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </div>
  );
};

export default Login;
