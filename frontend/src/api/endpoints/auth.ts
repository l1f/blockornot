import axios from "axios";
import { authHeader, Client } from "../client";

const API_PREFIX = "/api/auth/";

export const getInitialAuthHeaders = async (): Promise<authHeader> => {
  const data = await axios.get(API_PREFIX);
  return {
    token: data.headers["X-AUTH-TOKEN"],
    secret: data.headers["X-AUTH-SECRET"],
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
