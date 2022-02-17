import { actions } from "./actions";
import account from "../entities/Account";
import { accessData } from "../api/client";

export type initialStateType = {
  account?: account | null;
  accessData?: accessData | null;

  isAuthenticated: boolean;
  loading: boolean;
  error?: string | null;
};

export const AuthReducer = (
  state: initialStateType,
  action: actions
): initialStateType => {
  switch (action.type) {
    case "AUTH_INIT":
      return {
        ...state,
        loading: true,
      };
    case "AUTH_READY":
      return {
        ...state,
        accessData: action.accessData,
        loading: false,
      };
    case "AUTH_COMPLETE":
      return {
        ...state,
        accessData: {
          accessUrl: undefined,
          headers: undefined,
          client: action.client,
        },
        account: action.account,
        isAuthenticated: true,
        loading: false,
      };
    case "AUTH_ERROR":
      return {
        ...state,
        loading: false,
        accessData: undefined,
        isAuthenticated: false,
        error: action.error,
      };
    default:
      return state;
  }
};
