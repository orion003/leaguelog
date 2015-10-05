CREATE TABLE standing (
    id SERIAL PRIMARY KEY,
    season_id INTEGER NOT NULL REFERENCES season (id),
    team_id INTEGER NOT NULL REFERENCES team (id),
    wins SMALLINT DEFAULT 0,
    losses SMALLINT DEFAULT 0,
    ties SMALLINT DEFAULT 0,
    created TIMESTAMP NOT NULL,
    modified TIMESTAMP NOT NULL
);