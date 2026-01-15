import type { ValidationError } from 'express-validator';
import { CustomError } from './custom-errors.ts';

export class RequestValidationError extends CustomError {
  statusCode = 400;

  constructor(private errors: ValidationError[]) {
    super('Invalid request');

    Object.setPrototypeOf(this, RequestValidationError.prototype);
  }

  serializeErrors() {
    return this.errors.map((e) => {
      // Extract message from validation error
      // In express-validator v7, msg should be a string when using withMessage()
      let message: string = 'Invalid request parameter';
      
      // Debug: log the error structure to understand what we're getting
      console.log('Error structure:', JSON.stringify(e, null, 2));
      console.log('Error msg type:', typeof e.msg, 'msg value:', e.msg);
      
      // Handle different possible types of e.msg
      if (e.msg) {
        if (typeof e.msg === 'string' && e.msg.trim().length > 0) {
          message = e.msg;
        } else if (typeof e.msg === 'object') {
          // If msg is an object, try to extract a message
          if ('message' in e.msg && typeof e.msg.message === 'string') {
            message = e.msg.message;
          } else {
            // If it's an empty object or unknown structure, use a descriptive default
            const fieldName = 'path' in e && typeof e.path === 'string' ? e.path : 'field';
            message = `Invalid ${fieldName}`;
          }
        } else if (typeof e.msg !== 'function') {
          // Convert to string if it's not a function
          const msgStr = String(e.msg);
          message = msgStr || 'Invalid request parameter';
        }
      }
      
      // Ensure message is always a string (never an object)
      if (typeof message !== 'string') {
        message = 'Invalid request parameter';
      }
      
      return {
        type: e.type,
        value: 'value' in e ? e.value : undefined,
        message: message, // Explicitly ensure it's a string
        field: 'path' in e ? e.path : undefined,
        paramType: 'location' in e ? e.location : undefined,
      };
    });
  }
}
