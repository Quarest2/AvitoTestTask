CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO
$$
    BEGIN
        IF NOT EXISTS (SELECT 1
                       FROM pg_type
                       WHERE typname = 'organization_type'
                         and typnamespace = (select oid from pg_namespace where nspname = 'public')) THEN
            CREATE TYPE organization_type AS ENUM (
                'IE',
                'LLC',
                'JSC');
        END IF;

        IF NOT EXISTS (SELECT 1
                       FROM pg_type
                       WHERE typname = 'service_type'
                         and typnamespace = (select oid from pg_namespace where nspname = 'public')) THEN
            CREATE TYPE service_type AS ENUM (
                'Construction',
                'Delivery',
                'Manufacture'
                );
        END IF;

        IF NOT EXISTS (SELECT 1
                       FROM pg_type
                       WHERE typname = 'status'
                         and typnamespace = (select oid from pg_namespace where nspname = 'public')) THEN
            CREATE TYPE status AS ENUM (
                'Created',
                'Published',
                'Closed'
                );
        END IF;

        IF NOT EXISTS (SELECT 1
                       FROM pg_type
                       WHERE typname = 'bid_author_type'
                         and typnamespace = (select oid from pg_namespace where nspname = 'public')) THEN
            CREATE TYPE bid_author_type AS ENUM (
                'Organization',
                'User');
        END IF;
    END
$$;

CREATE TABLE IF NOT EXISTS employee
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username   VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name  VARCHAR(50),
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS organization
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(100) NOT NULL,
    description TEXT,
    type        organization_type,
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS organization_responsible
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    organization_id UUID REFERENCES organization (id) ON DELETE CASCADE,
    user_id         UUID REFERENCES employee (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tender
(
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name             VARCHAR(50),
    description      TEXT,
    status           status           DEFAULT 'Created',
    service_type     service_type,
    organization_id  UUID,
    creator_username VARCHAR(50),
    created_at       TIME,
    version          INTEGER          DEFAULT 1
);

-- CREATE TABLE IF NOT EXISTS tender_history (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     name VARCHAR(50),
--     description TEXT,
--     status status DEFAULT 'Created',
--     service_type service_type,
--     organization_id UUID,
--     creator_username VARCHAR(50),
--     created_at TIME,
--     version SERIAL DEFAULT 1
-- );

CREATE TABLE IF NOT EXISTS bid
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR(50),
    creator_username VARCHAR(50),
    description   TEXT,
    status        status           DEFAULT 'Created',
    tender_id     UUID,
    author_type   bid_author_type,
    author_id     UUID,
    created_at    TIME,
    version       INTEGER          DEFAULT 1,
    approve_count SERIAL,
    reject_count  SERIAL
);

-- CREATE TABLE IF NOT EXISTS bid_history (
--     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
--     name VARCHAR(50),
--     creator_username VARCHAR(50),
--     description TEXT,
--     status status DEFAULT 'Created',
--     tender_id UUID,
--     author_type bid_author_type,
--     author_id UUID,
--     created_at TIME,
--     version SERIAl DEFAULT 1,
--     approve_count SERIAL,
--     reject_count  SERIAL
-- );

CREATE TABLE IF NOT EXISTS bid_feedback
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    bid_id          uuid,
    author_username VARCHAR(50),
    comment         TEXT,
    created_at      TIME
);