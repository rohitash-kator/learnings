import { Router } from 'express';
import authRoutes from './auth/auth.ts';
import userRoutes from './users/users.ts';

const router = Router();

router.use('/auth', authRoutes);
router.use('/users', userRoutes);

export default router;
