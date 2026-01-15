import { Schema, model } from 'mongoose';

const userSchema = new Schema(
  {
    name: { type: String, required: false, default: 'User' },
    email: { type: String, required: true, unique: true },
    password: { type: String, required: true },
  },
  { timestamps: true }
);

const User = model('User', userSchema);

export { User };
