import { Router } from 'express';

import { currentUser } from '../../controllers/user.ts';

const router = Router();

router.get('/current-user', currentUser);

export default router;
