-- https://leetcode-cn.com/problems/classes-more-than-5-students/description/
SELECT class
FROM courses
GROUP BY class
HAVING COUNT(DISTINCT student) >= 5;