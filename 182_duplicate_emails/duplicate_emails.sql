
-- Solution used 637ms
-- beats 66.01% of Mysql submissions.
SELECT Email FROM Person 
GROUP BY Email 
HAVING Count(Email) > 1;