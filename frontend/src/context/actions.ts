import React from "react";
import { AxiosInstance } from "axios";

import account from "../api/entities/Account";
import { accessData, authHeader } from "../api/client";
import { completeAuth, getInitialAccessData } from "../api/endpoints/auth";

export type actions =
  | { type: "AUTH_INIT" }
  | { type: "AUTH_READY"; accessData: accessData }
  | { type: "AUTH_COMPLETE"; client: AxiosInstance; account: account }
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

export const dispatchCompleteAuth = async (
  pin: string,
  headers: authHeader,
  dispatch: React.Dispatch<actions>
) => {
  try {
    const [client, account] = await completeAuth(pin, headers);
    dispatch({ type: "AUTH_COMPLETE", client, account });
  } catch (error) {
    dispatch({ type: "AUTH_ERROR", error: "Error completing auth" });
  }
};
