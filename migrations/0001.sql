CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE project
(
    created_at  timestamptz NOT NULL,
    deleted_at  timestamptz,
    id          uuid PRIMARY KEY,
    name        text        NOT NULL UNIQUE,
    modified_at timestamptz NOT NULL,
    version     int         NOT NULL
);

CREATE TABLE article
(
    cover_image_id uuid,
    created_at     timestamptz NOT NULL,
    deleted_at     timestamptz,
    id             uuid PRIMARY KEY,
    modified_at    timestamptz NOT NULL,
    project_id     uuid        NOT NULL,
    title          text        NOT NULL,
    version        int         NOT NULL,

    FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
);

CREATE TABLE article_block
(
    article_id  uuid        NOT NULL,
    created_at  timestamptz NOT NULL,
    data        jsonb       NOT NULL,
    deleted_at  timestamptz,
    id          uuid PRIMARY KEY,
    modified_at timestamptz NOT NULL,
    sort_rank   text        NOT NULL,
    type        text        NOT NULL,
    version     int         NOT NULL,

    UNIQUE (article_id, sort_rank),
    FOREIGN KEY (article_id) REFERENCES article (id) ON DELETE CASCADE
);

CREATE TABLE tag
(
    created_at  timestamptz NOT NULL,
    deleted_at  timestamptz,
    id          uuid PRIMARY KEY,
    name        text        NOT NULL UNIQUE,
    modified_at timestamptz NOT NULL,
    version     int         NOT NULL
);

CREATE TABLE article_tag
(
    article_id  uuid        NOT NULL,
    created_at  timestamptz NOT NULL,
    id          uuid PRIMARY KEY,
    modified_at timestamptz NOT NULL,
    sort_rank   text        NOT NULL UNIQUE,
    tag_id      uuid        NOT NULL,
    version     int         NOT NULL,

    UNIQUE (article_id, tag_id),
    UNIQUE (article_id, sort_rank),
    FOREIGN KEY (article_id) REFERENCES article (id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE
);

CREATE TABLE image
(
    created_at  timestamptz NOT NULL,
    deleted_at  timestamptz,
    type        text        NOT NULL,
    height      int         NOT NULL,
    id          uuid PRIMARY KEY,
    modified_at timestamptz NOT NULL,
    version     int         NOT NULL,
    width       int         NOT NULL
);