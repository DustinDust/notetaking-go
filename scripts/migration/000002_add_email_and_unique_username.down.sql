ALTER TABLE IF EXISTS users 
DROP COLUMN IF EXISTS email;

ALTER TABLE users 
DROP CONSTRAINT IF EXISTS username_unique