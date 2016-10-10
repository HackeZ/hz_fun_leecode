-- Solution 1
-- Used time:   Beats: 25.86%
SELECT A.Id 
FROM Weather A, Weather B
WHERE A.Temperature > B.Temperature AND 
      TO_DAYS(A.Date)-TO_DAYS(B.Date)=1;