import { CustomError } from './custom-errors.ts';

export class BadRequestError extends CustomError {
  statusCode: number = 400;

  constructor(public message: string) {
    super(message);

    Object.setPrototypeOf(this, BadRequestError.prototype);
  }

  serializeErrors(): {
    type?: string;
    value?: string | undefined;
    message: string;
    field?: string | undefined;
    paramType?: string | undefined;
  }[] {
    return [
      {
        message: this.message,
      },
    ];
  }
}
