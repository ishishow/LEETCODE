select E2.Name as Employee
from Employee as E1 , Employee as E2
Where E1.id = E2.Managerid AND E1.Salary < E2.Salary