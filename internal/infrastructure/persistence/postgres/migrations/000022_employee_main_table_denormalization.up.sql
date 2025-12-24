begin
;

ALTER TABLE public.employees
  ADD COLUMN IF NOT EXISTS highest_academic_degree varchar,
  ADD COLUMN IF NOT EXISTS speciality varchar,
  ADD COLUMN IF NOT EXISTS current_workplace varchar;

-- Degrees: only fill missing values
with
    latest_degree as (
        select distinct
            on (ed.employee_id) ed.employee_id, ed.degree_level, ed.speciality
        from public.employee_degrees ed
        order by ed.employee_id, ed.date_end desc nulls last, ed.id desc
    )
    update public.employees e
    set highest_academic_degree = coalesce(e.highest_academic_degree, ld.degree_level),
    speciality = coalesce(e.speciality, ld.speciality)
from latest_degree ld
where ld.employee_id = e.id
;

-- Work: only fill missing value
with
    latest_work as (
        select distinct on (we.employee_id) we.employee_id, we.workplace
        from public.employee_work_experiences we
        where we.on_going is true
        order by we.employee_id, we.date_start desc nulls last, we.id desc
    )
    update public.employees e
    set current_workplace = coalesce(e.current_workplace, lw.workplace)
from latest_work lw
where lw.employee_id = e.id
;

commit
;
