// src/models/ApiData.ts

export interface ApiParameter {
  name: string;
  in: string;
  description?: string;
  required: boolean;
  type: string;
}

export interface ApiMethod {
  summary: string;
  description: string;
  parameters: ApiParameter[];
}

export interface ApiPath {
  [method: string]: ApiMethod;
}

export interface ApiData {
  info: {
    title: string;
    description: string;
  };
  paths: {
    [path: string]: ApiPath;
  };
}
