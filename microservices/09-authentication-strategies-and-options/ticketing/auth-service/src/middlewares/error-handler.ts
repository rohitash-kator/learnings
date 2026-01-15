import type { NextFunction, Request, Response } from 'express';
import { CustomError } from '../errors/custom-errors.ts';

const errorHandler = (
  err: Error,
  req: Request,
  res: Response,
  next: NextFunction
) => {
  if (err instanceof CustomError) {
    return res.status(err.statusCode).json({ errors: err.serializeErrors() });
  }
  res
    .status(400)
    .json({ errors: [{ message: err ?? 'Something went wrong!!' }] });
};

export { errorHandler };
