-- name: GetSummaryData :many
SELECT 
	e.id,
	ed.surname,
	ed."name",
	latest_experience.workplace,
	latest_degree.degree_level
FROM employees AS e
JOIN employee_details AS ed ON e.id = ed.employee_id
JOIN (
	SELECT DISTINCT employee_id FROM employee_socials
) AS socials ON e.id = socials.employee_id
JOIN (
	SELECT DISTINCT ON (employee_id)
		employee_id,
		workplace
	FROM employee_work_experiences
	WHERE employee_work_experiences.language_code = sqlc.arg(language_code)
	ORDER BY employee_work_experiences.employee_id, employee_work_experiences.date_end DESC NULLS FIRST
) AS latest_experience ON e.id = latest_experience.employee_id
JOIN (
	SELECT DISTINCT ON (employee_id)
		employee_id,
		degree_level,
		speciality
	FROM employee_degrees
	WHERE employee_degrees.language_code = sqlc.arg(language_code)
	ORDER BY employee_id, date_end DESC
) AS latest_degree ON e.id = latest_degree.employee_id
WHERE
	ed.language_code = sqlc.arg(language_code)
	AND ed.is_employee_details_new = true
ORDER BY e.id;
