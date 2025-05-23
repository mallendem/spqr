\c spqr-console

CREATE DISTRIBUTION ds1 COLUMN TYPES integer;

CREATE KEY RANGE krid2 FROM 101 ROUTE TO sh2 FOR DISTRIBUTION ds1;
CREATE KEY RANGE krid1 FROM 1 ROUTE TO sh1 FOR DISTRIBUTION ds1;

ALTER DISTRIBUTION ds1 ATTACH RELATION orders DISTRIBUTION KEY id;
ALTER DISTRIBUTION ds1 ATTACH RELATION delivery DISTRIBUTION KEY order_id;

\c regress

CREATE TABLE orders(id INT PRIMARY KEY);
CREATE TABLE delivery(id INT PRIMARY KEY, order_id INT, FOREIGN KEY(order_id) REFERENCES orders(id));

SET __spqr__engine_v2 TO false;

INSERT INTO orders(id) VALUES (5);
INSERT INTO delivery(id,order_id) VALUES (10, 5);
SELECT * FROM delivery;
SELECT * FROM delivery JOIN orders ON order_id = id;
SELECT * FROM delivery JOIN orders ON delivery.order_id = orders.id;

DROP TABLE orders CASCADE;
DROP TABLE delivery;

\c spqr-console
DROP DISTRIBUTION ALL CASCADE;
DROP KEY RANGE ALL;
