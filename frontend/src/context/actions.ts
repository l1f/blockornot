import React from "react";

import account from "../api/entities/Account";
import { authHeader } from "../api/client";
import { getInitialAuthHeaders } from "../api/endpoints/auth";

export type actions =
  | { type: "AUTH_INIT" }
  | { type: "AUTH_READY"; header: authHeader }
  | { type: "AUTH_COMPLETE"; header: authHeader }
  | { type: "AUTH_SUCCESS"; account: account }
  | { type: "LOGOUT" }
  | { type: "AUTH_ERROR"; error: string };

type ActionMap<M extends { [index: string]: any }> = {
  [Key in keyof M]: M[Key] extends undefined
    ? {
        type: Key;
      }
    : {
        type: Key;
        payload: M[Key];
      };
};

export type AuthActions = ActionMap<actions>[keyof ActionMap<actions>];

export const getAuthHeaders = async (dispatch: React.Dispatch<actions>) => {
  try {
    const header = await getInitialAuthHeaders();
    dispatch({ type: "AUTH_READY", header });
  } catch (error) {
    dispatch({ type: "AUTH_ERROR", error: "Error initializing auth" });
  }
};
