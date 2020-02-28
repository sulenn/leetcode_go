-- https://leetcode-cn.com/problems/customers-who-never-order/description/

SELECT Name AS Customers
FROM Customers
WHERE Id NOT IN (SELECT DISTINCT CustomerId
FROM Orders
);