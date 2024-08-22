-- name: ListOrders :many
select * from orders limit ? offset ?;

-- name: CreateOrder :exec
insert into orders (id, price, tax, final_price) values(?, ?, ?, ?);