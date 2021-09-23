-- join之前是from，之后是on
-- select * from a join b on c where d
-- left join: 如果on条件是 a.id=b.id; 则a表中与b表中的相同id才会进行连接，a表中的数据不会与b表中id不同的数据连接
-- 如果a表中的数据在b表中一条满足on条件的数据都没有，则会返回一条a表字段有数据b表字段全为null的数据
-- 如果on是多条件查询，例如：on a.id=b.id and a.name=b.name; 只有满足on条件时才会连接，如果a表中在b表一条数据都没匹配到，则会生成一条a表字段有数据b表字段全为null的数据
-- 在左连接下，如果 where 右边.字段 is null  表示左表中不满足on条件的数据； is not null则表示左表中满足on的数据

select
    session_id
from
(
    select * from playback
) as p
left join
(
    select * from ads
) as a
on p.customer_id=a.customer_id and a.timestamp between p.start_time and p.end_time
where ad_id is null;



# 1 
select
    session_id
from
    playback p left join ads a
on
    p.customer_id = a.customer_id
group by
    1
having
    sum(case when a.timestamp between p.start_time and p.end_time then 0 else 1 end) = count(1);