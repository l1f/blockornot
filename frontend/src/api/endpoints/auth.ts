import axios, { AxiosInstance } from "axios";
import { accessData, authHeader } from "../client";
import Account from "../entities/Account";

const API_PREFIX = "/api/v1/auth";

export const getInitialAccessData = async (): Promise<accessData> => {
  const response = await axios.get(API_PREFIX);
  const { data } = response;

  return {
    headers: {
      token: response.headers["x-auth-token"],
      secret: response.headers["x-auth-secret"],
    },
    accessUrl: data.access_url,
  };
};

export const completeAuth = async (
  pin: string,
  headers: authHeader
): Promise<[AxiosInstance, Account]> => {
  const response = await axios.post(
    API_PREFIX,
    {
      pin,
    },
    {
      headers: {
        "x-auth-token": headers.token,
        "x-auth-secret": headers.secret,
      },
    }
  );

  const { data } = response;

  return [
    axios.create({
      headers: {
        token: response.headers["x-auth-token"],
        secret: response.headers["x-auth-secret"],
      },
    }),
    data,
  ];
};
