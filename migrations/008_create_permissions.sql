-- CREATE TABLE permissions (
--     id SERIAL PRIMARY KEY,
--     actions TEXT[] NOT NULL
-- );

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    actions TEXT  -- Store CSV string
);