
-- Soltution by 509ms
-- beates 26.51% mysql submissions.
SELECT Name as Customers FROM Customers
WHERE Id NOT IN (
    SELECT customerId FROM Orders
    GROUP BY CustomerId
);

-- Soltution by 602ms
SELECT Name as Customers FROM Customers 
WHERE Id NOT IN (
    SELECT CustomerId FROM Orders
);

-- Soltution by 496ms
-- beates 44.61% mysql submissions.
SELECT Name as Customers FROM Customers as c
LEFT JOIN Orders as o ON c.Id=o.CustomerId
WHERE o.CustomerId IS NULL;
