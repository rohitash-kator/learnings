import { Document, Model, Schema, model } from 'mongoose';
import { Password } from '../services/password.ts';

// An interface that describes the properties that are required to create a new User
interface UserAttributes {
  name?: string;
  email: string;
  password: string;
}

// An interface that describes the properties that a User Model has
interface UserModel extends Model<UserDocument> {
  createUser(userAttributes: UserAttributes): UserDocument;
}

// An interface that describes the properties that a User Document has
interface UserDocument extends Document {
  name?: string;
  email: string;
  password: string;
}

const userSchema = new Schema(
  {
    name: { type: String, required: false, default: 'User' },
    email: { type: String, required: true, unique: true },
    password: { type: String, required: true },
  },
  { timestamps: true }
);

userSchema.pre('save', async function () {
  if (this.isModified('password')) {
    const hashedPassword = await Password.hash(this.get('password'));
    this.set('password', hashedPassword);
  }
});

userSchema.statics.createUser = (userAttributes: UserAttributes) => {
  return new User(userAttributes);
};

const User = model<UserDocument, UserModel>('User', userSchema);

export { User };
