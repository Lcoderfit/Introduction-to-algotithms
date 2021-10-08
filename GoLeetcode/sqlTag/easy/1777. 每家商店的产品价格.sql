
-- 行转列(max+case when+group)
-- 为什么要用max：因为假如product_id分组内有三条数据（store分别对应'store1', 'store2', 'store3'）
-- 那么确定store1这一列时：case when store = 'store1' then price else null end 也就是说store1这一行数据为price
-- store2和store3这两条数据的price为null，如果不用max，则会多出两条price为null的数据
select product_id,
    max(case when store='store1' then price else null end) store1,
    max(case when store='store1' then price else null end) store1,
    max(case when store='store1' then price else null end) store1
from Products
-- 这里可以写成group by 1; 表示select选择的第一列; order by 1同理
group by product_id;


-- 列转行(max+union+group by) 
-- union会做排序和去重操作，所以效率比union all要慢
select name, '语文' as subject, scores from table
union
select name, '数学' as subject, scores from table
union
select name, '英语' as subject, scores from table
