import axios, { AxiosInstance } from "axios";

export type Client = {
  axios: AxiosInstance;
  authorized: boolean;
};

export type authHeader = {
  token: string;
  secret: string;
};

export const getClient = (header?: authHeader): Client => {
  if (header) {
    return {
      authorized: true,
      axios: axios.create({
        headers: {
          "X-AUTH-TOKEN": header.token,
          "X-AUTH-SECRET": header.secret,
        },
      }),
    };
  }

  return {
    authorized: false,
    axios: axios,
  };
};
