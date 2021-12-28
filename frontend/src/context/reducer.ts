import { actions } from "./actions";
import account from "../api/entities/Account";
import { authHeader } from "../api/client";

export type initialStateType = {
  account?: account | null;
  header?: authHeader | null;

  loading: boolean;
  error?: string | null;
};

const initialState = {
  account: null,
  header: null,
  loading: true,
  error: null,
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
        header: action.header,
        loading: false,
      };
    case "AUTH_COMPLETE":
      return {
        ...state,
        header: action.header,
        loading: true,
      };
    case "AUTH_SUCCESS":
      return {
        ...state,
        loading: false,
        account: action.account,
      };
    case "AUTH_ERROR":
      return {
        ...state,
        loading: false,
        error: action.error,
      };
    default:
      return state;
  }
};
