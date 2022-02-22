-- Table: Employee
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | id          | int  |
-- | salary      | int  |
-- +-------------+------+
-- id is the primary key column for this table.
-- Each row of this table contains information about the salary of an employee.
-- Write an SQL query to report the second highest salary from the Employee table.
-- If there is no second highest salary, the query should report null.
-- The query result format is in the following example.
--
-- Results:
--   Runtime: 788 ms, faster than 89.73% of MS SQL Server online submissions for Second Highest Salary.
--   Memory Usage: 0B, less than 100.00% of MS SQL Server online submissions for Second Highest Salary.
SELECT MAX(salary) AS "SecondHighestSalary"
FROM Employee
WHERE salary < (SELECT MAX(salary) FROM Employee)