import React from "react";

import account from "../api/entities/Account";
import { accessData, authHeader } from "../api/client";
import { getInitialAccessData } from "../api/endpoints/auth";

export type actions =
  | { type: "AUTH_INIT" }
  | { type: "AUTH_READY"; accessData: accessData }
  | { type: "AUTH_COMPLETE"; header: authHeader }
  | { type: "AUTH_SUCCESS"; account: account }
  | { type: "LOGOUT" }
  | { type: "AUTH_ERROR"; error: string };

export const dispatchFetchInitialAuthData = async (
  dispatch: React.Dispatch<actions>
) => {
  try {
    const accessData = await getInitialAccessData();
    dispatch({ type: "AUTH_READY", accessData });
  } catch (error) {
    dispatch({ type: "AUTH_ERROR", error: "Error initializing auth" });
  }
};
