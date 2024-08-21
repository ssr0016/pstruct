
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    date_of_birth DATE NOT NULL
);


-- Drop existing foreign key constraint if it exists
ALTER TABLE user_permissions
DROP CONSTRAINT IF EXISTS user_permissions_user_id_fkey;

-- Alter column types to match
ALTER TABLE users
ALTER COLUMN id TYPE VARCHAR(255) USING id::VARCHAR;

ALTER TABLE user_permissions
ALTER COLUMN user_id TYPE VARCHAR(255) USING user_id::VARCHAR;

-- Recreate the foreign key constraint
ALTER TABLE user_permissions
ADD CONSTRAINT user_permissions_user_id_fkey
FOREIGN KEY (user_id) REFERENCES users(id);



