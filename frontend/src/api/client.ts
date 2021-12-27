import axios, { AxiosInstance } from "axios";

export type Client = {
  axios: AxiosInstance;
  authorized: boolean;
};

type AuthHeader = {
  token: string;
  secret: string;
};

export const getPreAuthHeader = async (): Promise<AuthHeader> => {
  // TODO:
  const data = await axios.get("localhost:8080/api/v1/auth");
  return {
    token: data.headers["X-AUTH-TOKEN"],
    secret: data.headers["X-AUTH-SECRET"],
  };
};

export const getClient = async (
  header?: AuthHeader,
  authorized: boolean = false
): Promise<Client> => {
  let _header = header;
  if (!_header) {
    _header = await getPreAuthHeader().then((preHeader) => preHeader);
  }

  return {
    axios: axios.create({
      headers: {
        "X-AUTH-TOKEN": _header!.token,
        "X-AUTH-SECRET": _header!.secret,
      },
    }),
    authorized: authorized,
  };
};
