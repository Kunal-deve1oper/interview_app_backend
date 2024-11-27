-- Create the database
CREATE DATABASE interview_db;

-- Use the database
USE interview_db;

-- Create the Organisation table
CREATE TABLE Organizations (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    adminNo INT NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the Admin table
CREATE TABLE Admins (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    organisation uuid NOT NULL,
    position VARCHAR(255),
    avatar VARCHAR(255),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (organisation) REFERENCES Organizations(id)
);

-- Create the Role table
CREATE TABLE Roles (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    skills VARCHAR(255) NOT NULL,
    experience INT NOT NULL,
    minATS INT NOT NULL,
    createdBy uuid NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (createdBy) REFERENCES Admins(id)
);

-- Create the Candidates table
CREATE TABLE Candidates (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    organisation VARCHAR(255) NOT NULL,
    experience INT NOT NULL,
    role uuid NOT NULL,
    cv VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255) NOT NULL,
    phoneNo VARCHAR(20),
    selected BOOLEAN DEFAULT FALSE,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role) REFERENCES Roles(id)
);

-- Create the Calendar table
-- CREATE TABLE Calendar (
--     id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
--     createdBy uuid NOT NULL,
--     date DATE NOT NULL,
--     role uuid NOT NULL,
--     time TIME NOT NULL,
--     candidates JSON NOT NULL,
--     createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     FOREIGN KEY (createdBy) REFERENCES Admin(id),
--     FOREIGN KEY (role) REFERENCES Role(id)
-- );
