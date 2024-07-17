BEGIN;
create table order_items
(
    id         varchar(191) not null
        primary key,
    order_id   varchar(191) null,
    product varchar(191) null,
    price      double       null,
    quantity   bigint       null,
    total      double       null,
    constraint fk_order_items_order
        foreign key (order_id) references orders (id)
);
create index idx_order_items_id
    on order_items (id);
create index idx_order_items_order_id
    on order_items (order_id);
COMMIT;