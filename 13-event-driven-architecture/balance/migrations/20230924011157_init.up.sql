CREATE TABLE balances (
    account_id_from varchar(255),
    account_id_to varchar(255),
    balance_account_id_from int(11),
    balance_account_id_to int(11)
);

INSERT INTO balances (account_id_from, account_id_to, balance_account_id_from, balance_account_id_to) VALUES ('5c1898f9-6831-4401-bc13-aec918913f2f', '541eb5fa-8247-49cf-a5d1-3c23b1734f1c', 10000, 5000);