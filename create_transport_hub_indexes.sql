-- ============================
-- Indexes for Transport_Hubs
-- ============================
CREATE INDEX idx_transport_hubs_iata_code ON Transport_Hubs(iata_code);
CREATE INDEX idx_transport_hubs_icao_code ON Transport_Hubs(icao_code);
CREATE INDEX idx_transport_hubs_city ON Transport_Hubs(city);
CREATE INDEX idx_transport_hubs_country ON Transport_Hubs(country);

-- ============================
-- Indexes for Operators
-- ============================
CREATE INDEX idx_operators_name ON Operators(name);
CREATE INDEX idx_operators_type ON Operators(type);
CREATE INDEX idx_operators_country ON Operators(country);

-- ============================
-- Indexes for Service_Lines
-- ============================
CREATE INDEX idx_service_lines_mode ON Service_Lines(mode);
CREATE INDEX idx_service_lines_name ON Service_Lines(name);

-- ============================
-- Indexes for Service_Stops
-- ============================
CREATE INDEX idx_service_stops_hub_id ON Service_Stops(hub_id);
CREATE INDEX idx_service_stops_service_id ON Service_Stops(service_id);

-- ============================
-- Indexes for Routes
-- ============================
CREATE INDEX idx_routes_origin_id ON Routes(origin_id);
CREATE INDEX idx_routes_destination_id ON Routes(destination_id);
CREATE INDEX idx_routes_operator_id ON Routes(operator_id);
CREATE INDEX idx_routes_mode ON Routes(mode);
