import { CustomError } from './custom-errors.ts';

export class DatabaseConnectionError extends CustomError {
  statusCode = 500;
  reason = 'Error connecting to database';
  constructor(message: string) {
    super('Error connecting to database');

    Object.setPrototypeOf(this, DatabaseConnectionError.prototype);
  }

  serializeErrors() {
    return [{ message: this.reason ?? 'Something went wrong!!' }];
  }
}
