CREATE TABLE league (
    id SERIAL PRIMARY KEY,
    name VARCHAR (64) UNIQUE NOT NULL,
    sport VARCHAR (32) NOT NULL,
    created TIMESTAMP NOT NULL,
    modified TIMESTAMP NOT NULL
);

CREATE TABLE season (
    id SERIAL PRIMARY KEY,
    league_id INTEGER NOT NULL REFERENCES league (id),
    name VARCHAR (64),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    created TIMESTAMP NOT NULL,
    modified TIMESTAMP NOT NULL
);

CREATE TABLE team (
    id SERIAL PRIMARY KEY,
    league_id INTEGER NOT NULL REFERENCES league (id),
    name VARCHAR (64) NOT NULL,
    created TIMESTAMP NOT NULL,
    modified TIMESTAMP NOT NULL
);

CREATE TABLE game (
    id SERIAL PRIMARY KEY,
    season_id INTEGER NOT NULL REFERENCES season (id),
    start_time TIMESTAMP NOT NULL,
    home_team_id INTEGER NOT NULL REFERENCES team (id),
    away_team_id INTEGER NOT NULL REFERENCES team (id),
    home_score SMALLINT DEFAULT 0,
    away_score SMALLINT DEFAULT 0
);