import React from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { createTheme, CssBaseline, ThemeProvider, useMediaQuery } from '@mui/material'

import { AuthProvider } from './context/context'

import Login from './pages/login/login.page'
import Main from './pages/main/main.page'
import RequireAuth from './components/RequireAuth/RequireAuth.component'
import { teal } from '@mui/material/colors'

function App() {
  const prefersDarkMode = useMediaQuery('(prefers-color-scheme: dark)');

  const theme = React.useMemo(
    () =>
      createTheme({
        palette: {
          mode: prefersDarkMode ? 'dark' : 'light',
          ...(!prefersDarkMode ?
            {
              // palette values for light mode
              primary: {
                main: teal[400]
              }
            } :
            {
              // palette values for dark mode
              primary: {
                main: teal[900]
              }
            })
        },
      }),
    [prefersDarkMode],
  );

  return (
    <AuthProvider>
      <BrowserRouter>
        <ThemeProvider theme={theme}>
          <CssBaseline />
          <Routes>
            <Route
              path="/"
              element={
                <RequireAuth>
                  <Main />
                </RequireAuth>
              }
            />
            <Route path="/login" element={<Login />} />
          </Routes>
        </ThemeProvider>
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
