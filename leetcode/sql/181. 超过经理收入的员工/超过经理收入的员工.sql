-- https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/description/

SELECT e2.Name AS Employee
FROM Employee e1 INNER JOIN Employee e2 ON e1.ID = e2.ManagerId AND e2.Salary > e1.Salary;