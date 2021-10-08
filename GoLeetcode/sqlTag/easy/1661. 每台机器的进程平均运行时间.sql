select
    t.machine_id,
    round(avg(t.end_time-t.start_time), 3) processing_time
from
(
    select 
        machine_id,
        process_id,
        max(case when activity_type='start' then timestamp else null end) start_time,
        max(case when activity_type='end' then timestamp else null end) end_time
    from
        activity
    group by machine_id, process_id
) t
group by t.machine_id;
