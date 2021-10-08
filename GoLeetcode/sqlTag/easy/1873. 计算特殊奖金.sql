-- if语句
select 
    employee_id, 
    if(mod(employee_id, 2)=1 and name not like 'M%', salary, 0) bonus
from Employees
order by employee_id asc; 


-- case when
select
    employee_id,
    (case when mod(employee_id, 2)=1 and name not like 'M%') bonus
from employees
order by employee_id asc;