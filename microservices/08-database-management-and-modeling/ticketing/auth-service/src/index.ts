import express, {
  type Express,
  type NextFunction,
  type Request,
  type Response,
} from 'express';
import mongoose from 'mongoose';

import routes from './routes/routes.ts';
import { errorHandler } from './middlewares/error-handler.ts';
import { NotFoundError } from './errors/not-found-error.ts';

const app: Express = express();

app.use(express.json());

app.use('/api', routes);

app.get('/', (req: Request, res: Response) => {
  res.status(200).json({ message: 'App is running...' });
});

app.use(async (req: Request, res: Response, next: NextFunction) => {
  throw new NotFoundError();
});
app.use(errorHandler);

// connect to mongodb and start the server
mongoose
  .connect('mongodb://auth-mongo-service.default:27017/micro-auth-service')
  .then(() => {
    console.log('Connected to MongoDB database!!');
    app.listen(3000, () => {
      console.log('Auth service is listening on port 3000');
    });
  })
  .catch((error) => {
    console.error(error);
  });

// alternative way using async/await
// const MONGO_URI =
//   process.env.MONGO_URI ??
//   'mongodb://auth-mongo-service.default:27017/micro-auth-service';
// const PORT = Number(process.env.PORT ?? 3000);

// async function start(): Promise<void> {
//   try {
//     await mongoose.connect(MONGO_URI);
//     console.log('Connected to the database');
//     app.listen(PORT, () =>
//       console.log(`Auth service listening on port ${PORT}`)
//     );
//   } catch (err) {
//     console.error('Failed to connect to MongoDB', err);
//     process.exit(1);
//   }
// }

// await start();

// // graceful shutdown
// const gracefulShutdown = async (signal: string) => {
//   console.log(`Received ${signal}, closing mongoose connection...`);
//   try {
//     await mongoose.disconnect();
//   } finally {
//     process.exit(0);
//   }
// };

// process.on('SIGINT', () => gracefulShutdown('SIGINT'));
// process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));

// process.on('uncaughtException', (err) => {
//   console.error('Uncaught Exception', err);
//   process.exit(1);
// });

// process.on('unhandledRejection', (reason) => {
//   console.error('Unhandled Rejection', reason);
//   // optionally exit: process.exit(1);
// });
