CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,       -- Auto-incrementing ID
    user_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    UNIQUE (user_id, role_id)    -- Ensure a user cannot have the same role more than once
);
