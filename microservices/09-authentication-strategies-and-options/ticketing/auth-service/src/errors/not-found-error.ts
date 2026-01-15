import { CustomError } from './custom-errors.ts';

export class NotFoundError extends CustomError {
  statusCode: number = 404;

  constructor() {
    super('Route not found');

    Object.setPrototypeOf(this, NotFoundError.prototype);
  }

  serializeErrors() {
    return [{ message: 'Not Found' }];
  }
}
