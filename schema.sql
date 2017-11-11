CREATE TABLE IF NOT EXISTS rider_events (
    action varchar(50) NOT NULL,
    station varchar(50) NOT NULL,
    line varchar(50) NOT NULL,
    timestamp timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS line_tickets (
    line varchar(50) NOT NULL PRIMARY KEY,
    ticket_price float NOT NULL DEFAULT 1.99
);

INSERT INTO line_tickets
(line, ticket_price)
VALUES
("Tomato", 1.99),
("Avocado", 2.99),
("Blueberry", 1.99),
("Orange", 1.49),
("Banana", 2.49)
;
