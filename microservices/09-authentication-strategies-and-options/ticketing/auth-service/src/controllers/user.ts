import type { Request, Response } from 'express';

const currentUser = (req: Request, res: Response) => {
  res.send('Hi There! from Current User');
};

export { currentUser };
