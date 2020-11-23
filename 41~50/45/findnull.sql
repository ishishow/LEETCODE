select Customers.Name as Customers
from Customers left join Orders
on Customers.Id = Orders.CustomerId
where Orders.Id is NULL


-- in 
select customers.name as 'Customers'
from customers
where customers.id not in
(
    select customerid from orders
);