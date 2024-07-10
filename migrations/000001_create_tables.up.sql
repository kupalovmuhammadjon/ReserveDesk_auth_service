CREATE TABLE users (
    id UUID PRIMARY KEY default gen_random_uuid(),
    full_name VARCHAR,
    is_admin BOOLEAN,
    email VARCHAR UNIQUE,
    password VARCHAR,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY default gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    token TEXT UNIQUE,
    revoked BOOLEAN,
    created_at TIMESTAMP default current_timestamp,
    updated_at TIMESTAMP default current_timestamp,
    deleted_at TIMESTAMP
);
