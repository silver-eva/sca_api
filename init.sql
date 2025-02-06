create schema if not exists app;

create extension if not exists "uuid-ossp";

-- Create `cat` table
CREATE TABLE IF NOT EXISTS app.cat (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    experience FLOAT CHECK (experience BETWEEN 1 AND 15),
    breed VARCHAR(48) NOT NULL,
    salary FLOAT DEFAULT 1000
);

-- Create `target` table
CREATE TABLE IF NOT EXISTS app.target (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    country VARCHAR(255) NOT NULL,
    notes TEXT,
    complited BOOLEAN DEFAULT FALSE
);

-- Create `mission` table
CREATE TABLE IF NOT EXISTS app.mission (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) UNIQUE NOT NULL,
    complited BOOLEAN DEFAULT FALSE
);

-- Create `mission_targets` join table
CREATE TABLE IF NOT EXISTS app.mission_targets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mission_id UUID NOT NULL,
    target_id UUID NOT NULL,
    FOREIGN KEY (mission_id) REFERENCES app.mission (id) ON DELETE CASCADE,
    FOREIGN KEY (target_id) REFERENCES app.target (id) ON DELETE CASCADE
);

-- Create `mission_cats` join table
CREATE TABLE IF NOT EXISTS app.mission_cats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mission_id UUID NOT NULL,
    cat_id UUID NOT NULL,
    FOREIGN KEY (mission_id) REFERENCES app.mission (id) ON DELETE CASCADE,
    FOREIGN KEY (cat_id) REFERENCES app.cat (id) ON DELETE CASCADE
);
