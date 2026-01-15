import type { NextFunction, Request, Response } from 'express';
import { validationResult } from 'express-validator';
import { RequestValidationError } from '../errors/request-validation-error.ts';
import { User } from '../models/user.ts';

const signup = async (req: Request, res: Response, next: NextFunction) => {
  try {
    const errors = validationResult(req);

    if (!errors.isEmpty()) {
      throw new RequestValidationError(errors.array());
    }
    const { email, password } = req.body;

    const existingUser = await User.findOne({ email });

    if (existingUser) {
      console.log('Email already exists. Use different one.');
      return res
        .status(400)
        .json({ message: 'Email already exists. Use different one.' });
    }

    const user = new User({ email, password });
    await user.save();
    res.status(201).json({ message: 'User created.', user });
  } catch (error) {
    console.error(error);
    return next(error);
  }
};

const signin = (req: Request, res: Response) => {
  res.send('Hi There! from Signin');
};

const signout = (req: Request, res: Response) => {
  res.send('Hi There! from Signout');
};

export { signup, signin, signout };
