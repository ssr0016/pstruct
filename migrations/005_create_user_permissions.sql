

CREATE TABLE user_permissions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    permission_id INT NOT NULL REFERENCES permissions(id),
    action VARCHAR(255),
    scope VARCHAR(255)
);

-- -- Drop existing foreign key constraint if it exists
-- ALTER TABLE user_permissions
-- DROP CONSTRAINT IF EXISTS user_permissions_user_id_fkey;

-- -- Alter column types to match
-- ALTER TABLE users
-- ALTER COLUMN id TYPE VARCHAR(255) USING id::VARCHAR;

-- ALTER TABLE user_permissions
-- ALTER COLUMN user_id TYPE VARCHAR(255) USING user_id::VARCHAR;

-- -- Recreate the foreign key constraint
-- ALTER TABLE user_permissions
-- ADD CONSTRAINT user_permissions_user_id_fkey
-- FOREIGN KEY (user_id) REFERENCES users(id);
