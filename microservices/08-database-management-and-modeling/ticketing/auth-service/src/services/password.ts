import { scrypt, randomBytes } from 'crypto';
import { promisify } from 'util';

const scryptAsync = promisify(scrypt);

export class Password {
  static async hash(password: string): Promise<string> {
    const salt = randomBytes(8).toString('hex');
    const buffer = (await scryptAsync(password, salt, 64)) as Buffer;

    return `${buffer.toString('hex')}.${salt}`;
  }

  static async compare(
    storedPasswordHash: string,
    suppliedPassword: string
  ): Promise<boolean> {
    const [hashedPassword, salt] = storedPasswordHash.split('.');
    const buffer = (await scryptAsync(
      suppliedPassword,
      salt as string,
      64
    )) as Buffer;

    return buffer.toString('hex') === hashedPassword;
  }
}
