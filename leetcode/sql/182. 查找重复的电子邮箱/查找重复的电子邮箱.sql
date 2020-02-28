-- https://leetcode-cn.com/problems/duplicate-emails/
SELECT Email
FROM Person
GROUP BY Email
HAVING COUNT(ID) > 1;