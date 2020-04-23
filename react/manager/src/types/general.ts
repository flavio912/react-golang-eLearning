export type GraphError = {
  message: string;
  path: string[];
  extensions?: {
    message: string;
    type: string;
    helpText?: string;
    title?: string;
  };
};
