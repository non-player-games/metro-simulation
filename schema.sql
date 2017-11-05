CREATE TABLE rider_events (
    action varchar(50) NOT NULL,
    station varchar(50) NOT NULL,
    line varchar(50) NOT NULL,
    timestamp timestamp DEFAULT CURRENT_TIMESTAMP
);
