create table orders
(
    id       varchar(191) not null
        primary key,
    customer longtext     null,
    date     datetime(3)  null,
    total    double       null
);
