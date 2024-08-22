BEGIN;
insert into orders (id, price, tax, final_price) values('046299c9-2611-4562-aac2-3fbd3f9eb812', 10, 0.5, 10.5);
insert into orders (id, price, tax, final_price) values('ea02aa88-b4ba-4fc1-b95d-57c1eeaa97f6', 100, 50, 150);
insert into orders (id, price, tax, final_price) values('f9d8b55b-4d44-46db-99fb-6aeca6643f63', 13, 5, 18);
COMMIT;
