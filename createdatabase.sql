-- ============================
-- Create Database (optional)
-- ============================
CREATE DATABASE transport_hub_crawler;
\c transport_hub_crawler;

-- ============================
-- Table: Transport_Hubs
-- ============================
CREATE TABLE Transport_Hubs (
    hub_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT CHECK (type IN ('airport', 'train_station', 'metro_station')),
    iata_code CHAR(3),
    icao_code CHAR(4),
    city TEXT,
    country TEXT,
    latitude DECIMAL(9,6),
    longitude DECIMAL(9,6),
    wiki_url TEXT
);

-- ============================
-- Table: Operators
-- ============================
CREATE TABLE Operators (
    operator_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT CHECK (type IN ('airline', 'rail', 'metro')),
    iata_code CHAR(2),
    icao_code CHAR(3),
    country TEXT,
    wiki_url TEXT
);

-- ============================
-- Table: Service_Lines
-- ============================
CREATE TABLE Service_Lines (
    service_id SERIAL PRIMARY KEY,
    operator_id INT NOT NULL REFERENCES Operators(operator_id) ON DELETE CASCADE,
    mode TEXT CHECK (mode IN ('air', 'rail', 'metro')),
    name TEXT NOT NULL,
    seasonal BOOLEAN DEFAULT FALSE,
    notes TEXT
);

-- ============================
-- Table: Service_Stops
-- ============================
CREATE TABLE Service_Stops (
    service_id INT NOT NULL REFERENCES Service_Lines(service_id) ON DELETE CASCADE,
    stop_order INT NOT NULL,
    hub_id INT NOT NULL REFERENCES Transport_Hubs(hub_id) ON DELETE CASCADE,
    arrival_time TEXT,
    departure_time TEXT,
    PRIMARY KEY (service_id, stop_order)
);

-- ============================
-- Optional Derived Table: Routes
-- ============================
CREATE TABLE Routes (
    origin_id INT NOT NULL REFERENCES Transport_Hubs(hub_id) ON DELETE CASCADE,
    destination_id INT NOT NULL REFERENCES Transport_Hubs(hub_id) ON DELETE CASCADE,
    service_id INT NOT NULL REFERENCES Service_Lines(service_id) ON DELETE CASCADE,
    operator_id INT NOT NULL REFERENCES Operators(operator_id) ON DELETE CASCADE,
    mode TEXT CHECK (mode IN ('air', 'rail', 'metro')),
    PRIMARY KEY (origin_id, destination_id, service_id)
);
