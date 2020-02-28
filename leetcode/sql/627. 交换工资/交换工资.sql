-- # Write your MySQL query statement below
-- https://leetcode-cn.com/problems/swap-salary/
UPDATE salary
SET sex = CHAR(ASCII(sex) ^ ASCII('f') ^ ASCII('m'));