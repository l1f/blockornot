import axios from "axios";
import { accessData, authHeader, Client } from "../client";

const API_PREFIX = "/api/v1/auth";

export const getInitialAccessData = async (): Promise<accessData> => {
  const response = await axios.get(API_PREFIX);
  const { data } = await response;

  return {
    headers: {
      token: response.headers["x-auth-token"],
      secret: response.headers["x-auth-secret"],
    },
    accessUrl: data.access_url,
  };
};

export const getAuthHeaders = async (
  client: Client,
  pin: string
): Promise<authHeader> => {
  const data = await client.axios.post(API_PREFIX, {
    pin: pin,
  });

  return {
    token: data.headers["X-AUTH-TOKEN"],
    secret: data.headers["X-AUTH-SECRET"],
  };
};
