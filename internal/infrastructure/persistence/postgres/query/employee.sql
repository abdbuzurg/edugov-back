-- name: CreateEmployee :one
INSERT INTO employees (
  unique_id,
  user_id,
  gender,
  tin
) VALUES(
  $1, $2, $3, $4
) RETURNING id, created_at, updated_at;

-- name: DeleteEmployee :exec
delete from employees
where id = $1
;

-- name: GetEmployeeByID :one
select *
from employees
where id = $1
;

-- name: GetEmployeeByUniqueIdentifier :one
select *
from employees
where unique_id = $1
;

-- name: GetEmployeeByUserID :one
select *
from employees
where user_id = $1
;

-- name: GetPersonnelPaginated :many
select
    e.id,
    e.unique_id,
    e.gender,
    e.tin,
    e.highest_academic_degree,
    e.speciality,
    e.current_workplace,
    ed.surname,
    ed.name,
    ed.middlename
from employees e
join
    employee_details ed
    on ed.employee_id = e.id
    and ed.is_employee_details_new is true
    and ed.language_code = sqlc.arg(language_code)
where
    -- must exist in employee_socials
    exists (select 1 from employee_socials es where es.employee_id = e.id)
    -- required non-null denormalized fields
    and e.highest_academic_degree is not null
    and e.speciality is not null
    and e.current_workplace is not null

    -- optional filters (pass NULL to ignore)
    and (nullif(sqlc.arg(uid)::text, '') is null or e.unique_id = sqlc.arg(uid))
    and (
        nullif(sqlc.arg(name)::text, '') is null
        or ed.name ilike '%' || sqlc.arg(name) || '%'
    )
    and (
        nullif(sqlc.arg(surname)::text, '') is null
        or ed.surname ilike '%' || sqlc.arg(surname) || '%'
    )
    and (
        nullif(sqlc.arg(middlename)::text, '') is null
        or ed.middlename ilike '%' || sqlc.arg(middlename) || '%'
    )
    and (
        nullif(sqlc.arg(workplace)::text, '') is null
        or e.current_workplace = sqlc.arg(workplace)
    )
    and (
        nullif(sqlc.arg(academic_degree)::text, '') is null
        or e.highest_academic_degree = sqlc.arg(academic_degree)
    )
order by e.id
limit sqlc.arg('limit')
offset sqlc.arg(page)
;

-- name: CountPersonnel :one
select count(*)::bigint as total
from employees e
join
    employee_details ed
    on ed.employee_id = e.id
    and ed.is_employee_details_new is true
    and ed.language_code = sqlc.arg(language_code)
where
    exists (select 1 from employee_socials es where es.employee_id = e.id)
    and e.highest_academic_degree is not null
    and e.speciality is not null
    and e.current_workplace is not null
    and (nullif(sqlc.arg(uid)::text, '') is null or e.unique_id = sqlc.arg(uid))
    and (
        nullif(sqlc.arg(name)::text, '') is null
        or ed.name ilike '%' || sqlc.arg(name) || '%'
    )
    and (
        nullif(sqlc.arg(surname)::text, '') is null
        or ed.surname ilike '%' || sqlc.arg(surname) || '%'
    )
    and (
        nullif(sqlc.arg(middlename)::text, '') is null
        or ed.middlename ilike '%' || sqlc.arg(middlename) || '%'
    )
    and (
        nullif(sqlc.arg(workplace)::text, '') is null
        or e.current_workplace = sqlc.arg(workplace)
    )
    and (
        nullif(sqlc.arg(academic_degree)::text, '') is null
        or e.highest_academic_degree = sqlc.arg(academic_degree)
    )
;

-- name: ListUniqueWorkplaces :many
select distinct e.current_workplace
from employees e
join
    employee_details ed
    on ed.employee_id = e.id
    and ed.is_employee_details_new is true
    and ed.language_code = sqlc.arg(language_code)
where
    exists (select 1 from employee_socials es where es.employee_id = e.id)
    and e.highest_academic_degree is not null
    and e.speciality is not null
    and e.current_workplace is not null
    and e.current_workplace <> ''
order by e.current_workplace asc
;

-- name: ListUniqueHighestAcademicDegrees :many
select distinct ed.speciality
from employee_degrees ed
join employees e on e.id = ed.employee_id
where
    ed.speciality is not null
    and ed.speciality <> ''
    and e.degree_level in (
        'Номзади илм',
        'PhD (Доктори фалсафа)',
        'Доктори илм'
    )
order by ed.speciality asc
;
