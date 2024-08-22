create table orders
(
    id varchar(191) not null primary key,
    price double not null,
    tax double not null,
    final_price double not null
);
