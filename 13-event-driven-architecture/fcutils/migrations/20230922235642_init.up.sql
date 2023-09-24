CREATE TABLE clients (
    id varchar(255),
    name varchar(255),
    email varchar(255),
    created_at date
);

CREATE TABLE accounts (
    id varchar(255),
    client_id varchar(255),
    balance int(11),
    created_at date
);

CREATE TABLE transactions (
    id varchar(255),
    account_id_from varchar(255),
    account_id_to varchar(255),
    amount int(11),
    created_at date
);

INSERT INTO clients (id, name, email, created_at) VALUES ('455965ee-c29b-4971-850c-a41f28ea0167', 'Jeff Bezos', 'jeff.bezos@email.com', '2023-09-22 23:56:42');
INSERT INTO clients (id, name, email, created_at) VALUES ('7d1bd97c-9f3d-4b2a-a7ce-44a979408e61', 'Elon Musk', 'elon.musk@email.com', '2023-09-22 23:56:42');

INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('5c1898f9-6831-4401-bc13-aec918913f2f', '455965ee-c29b-4971-850c-a41f28ea0167', 10000, '2023-09-22 23:56:42');
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('541eb5fa-8247-49cf-a5d1-3c23b1734f1c', '7d1bd97c-9f3d-4b2a-a7ce-44a979408e61', 5000, '2023-09-22 23:56:42');