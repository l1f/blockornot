import React from "react";
import { AxiosInstance } from "axios";

import account from "../entities/Account";
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
    const data = await getInitialAccessData();
    dispatch({ type: "AUTH_READY", accessData: data });
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
    const [client, responseData] = await completeAuth(pin, headers);
    dispatch({ type: "AUTH_COMPLETE", client, account: responseData });
  } catch (error) {
    dispatch({ type: "AUTH_ERROR", error: "Error completing auth" });
  }
};
