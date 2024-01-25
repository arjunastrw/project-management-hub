CREATE DATABASE IF NOT EXISTS project_management_db;

CREATE EXTENSION "uuid-ossp";

CREATE TYPE role_type AS ENUM ('ADMIN','MANAGER', 'TEAM MEMBER');

CREATE TYPE task_status AS ENUM('In Progress', 'Blocked', 'Waiting Approval', 'Accepted', 'Rejected', 'On Hold');

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role role_type NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE projects (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    manager_id UUID NOT NULL,
    deadline DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (manager_id) REFERENCES users(id)
);


CREATE TABLE project_members (
    member_id UUID NOT NULL,
    project_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (member_id, project_id),
    FOREIGN KEY (member_id) REFERENCES users(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);


CREATE TABLE tasks (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status task_status NOT NULL,
    approval BOOLEAN,
    person_in_charge UUID NOT NULL,
    deadline DATE NOT NULL,
    project_id UUID NOT NULL,
    approval_date DATE,
    feedback TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (person_in_charge) REFERENCES users(id),
    FOREIGN KEY (project_id) REFERENCES projects(id)
);


CREATE TABLE reports (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID NOT NULL,
    report TEXT NOT NULL,
    task_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);