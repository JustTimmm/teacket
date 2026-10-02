CREATE TYPE status AS ENUM (
    'unspecified',
    'open',
    'in_progress',
    'resolved',
    'closed'
);

CREATE TABLE tickets (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title TEXT NOT NULL, content TEXT NOT NULL,
    author_id BIGINT NOT NULL,
    status status NOT NULL default 'open',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now() NOT NULL
);