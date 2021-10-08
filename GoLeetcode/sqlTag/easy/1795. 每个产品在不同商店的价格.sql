-- where条件中不能使用列的别名
-- 列转行
select product_id, 'store1' store, store1 price from Products where store1 is not null
union
select product_id, 'store2' store, store2 price from Products where store2 is not null
union
select product_id, 'store3' store, store3 price from Products where store3 is not null
