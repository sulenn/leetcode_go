-- https://leetcode-cn.com/problems/second-highest-salary/submissions/
SELECT (SELECT Salary
FROM Employee
ORDER BY Salary DESC
LIMIT 1, 1) AS SecondHighestSalary;