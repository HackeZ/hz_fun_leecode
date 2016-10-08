-- Solution by 783ms
-- beates 67.34% mysql submissions.
DELETE Person FROM 
Person WHERE Id NOT IN (
    SELECT Id FROM (
        SELECT MIN(Id) as Id FROM
        Person GROUP BY Email
    ) p
);

-- Why use `p` in there?
-- When you INSERT/UPDATE/DELETE/ on a table, you can not 
-- reference that table in an inner query(you can 
-- however reference a field from that outer table...)