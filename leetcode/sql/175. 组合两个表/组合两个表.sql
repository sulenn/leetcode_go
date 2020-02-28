-- https://leetcode-cn.com/problems/combine-two-tables/description/
-- 需要用到左连接
SELECT p.FirstName, p.LastName, a.City, a.State
FROM Person p LEFT JOIN Address a ON p.PersonId = a.PersonId;