-- # Write your MySQL query statement below
-- https://leetcode-cn.com/problems/not-boring-movies/description/
SELECT *
FROM cinema
WHERE id%2 = 1 AND description != 'boring'
ORDER BY rating DESC;