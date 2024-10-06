CREATE TABLE apartment (
    id SERIAL PRIMARY KEY,
    building_id INTEGER REFERENCES building(id) NOT NULL,
    number VARCHAR(50) NOT NULL,
    floor INTEGER NOT NULL,
    sq_meters INTEGER NOT NULL
);
