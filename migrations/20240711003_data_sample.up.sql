BEGIN;
insert into orders (id, customer, date, total) values('046299c9-2611-4562-aac2-3fbd3f9eb812', 'Roberto', NOW(), 15025.0);
insert into order_items (id,order_id,product,price,quantity,total) values ('77204195-5b2d-4c79-a026-007dcc342256', '046299c9-2611-4562-aac2-3fbd3f9eb812','RTX 4090 TI',1500.0,1,1500.0);
insert into order_items (id,order_id,product,price,quantity,total) values ('dc59682d-2ff7-4d4e-bc4e-d3b413436157', '046299c9-2611-4562-aac2-3fbd3f9eb812','Cabo USB 50m',12.5,2,25.0);
COMMIT;
