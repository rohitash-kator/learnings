export abstract class CustomError extends Error {
  abstract statusCode: number;

  constructor(message: string) {
    super(message);
    Object.setPrototypeOf(this, CustomError.prototype);
  }

  abstract serializeErrors(): {
    type?: string;
    value?: string | undefined;
    message: string;
    field?: string | undefined;
    paramType?: string | undefined;
  }[];
}
