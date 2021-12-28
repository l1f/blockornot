import { AxiosInstance } from "axios";

export type Client = {
  axios: AxiosInstance;
  authorized: boolean;
};

export type authHeader = {
  token: string;
  secret: string;
};

export type accessUrl = {
  Scheme: string;
  Host: string;
  Path: string;
  ForceQuery: boolean;
  RawQuery: string;
};

export type accessData = {
  headers: authHeader;
  accessUrl?: accessUrl;
};
