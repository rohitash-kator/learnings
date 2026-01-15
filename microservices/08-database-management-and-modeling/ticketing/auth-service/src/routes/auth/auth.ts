import { Router } from 'express';
import { signin, signout, signup } from '../../controllers/auth.ts';
import { body } from 'express-validator';

const router = Router();

router.post(
  '/signup',
  body('email').trim().isEmail().withMessage('Email must be valid'),
  body('password')
    .trim()
    .isLength({ min: 6, max: 32 })
    .withMessage('Password must be between 6 and 32 characters'),
  signup
);
router.post('/signin', signin);
router.post('/signout', signout);

export default router;
