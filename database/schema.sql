DROP TABLE IF EXISTS companies;
CREATE TABLE companies (
    id VARCHAR(256) PRIMARY KEY NOT NULL,
    name VARCHAR(256) NOT NULL,
    logo VARCHAR(256) NOT NULL,
    rating FLOAT DEFAULT 0.0,
    deleted boolean DEFAULT 0
);

DROP TABLE IF EXISTS facts;
CREATE TABLE facts (
    id VARCHAR(256) PRIMARY KEY NOT NULL,
    summary VARCHAR(1024)  NOT NULL,
    citation VARCHAR(256) NOT NULL,
    company_id VARCHAR(256) REFERENCES companies(id),
    deleted boolean DEFAULT 0
);

DROP TABLE IF EXISTS proposed_facts;
CREATE TABLE proposed_facts(
    id VARCHAR(256) PRIMARY KEY NOT NULL,
    summary VARCHAR(1024) NOT NULL,
    citation VARCHAR(256) UNIQUE NOT NULL,
    company_id VARCHAR(256),
    company_name VARCHAR(256),
    date_added DATETIME,
    approved boolean DEFAULT 0,
    rejected boolean DEFAULT 0
);
