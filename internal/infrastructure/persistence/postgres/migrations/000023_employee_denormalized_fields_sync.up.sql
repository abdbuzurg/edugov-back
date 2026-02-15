CREATE OR REPLACE FUNCTION refresh_employee_denormalized_fields(p_employee_id BIGINT)
RETURNS VOID
LANGUAGE plpgsql
AS $$
DECLARE
  v_highest_academic_degree VARCHAR;
  v_speciality VARCHAR;
  v_current_workplace VARCHAR;
BEGIN
  IF p_employee_id IS NULL THEN
    RETURN;
  END IF;

  -- Rule from request:
  -- highest_academic_degree <- latest employee_degrees.speciality
  -- speciality <- latest employee_degrees.degree_level
  SELECT
    ed.speciality,
    ed.degree_level
  INTO
    v_highest_academic_degree,
    v_speciality
  FROM employee_degrees ed
  WHERE ed.employee_id = p_employee_id
  ORDER BY ed.date_end DESC NULLS LAST, ed.id DESC
  LIMIT 1;

  -- current_workplace <- latest ongoing employee_work_experiences.workplace
  SELECT
    we.workplace
  INTO
    v_current_workplace
  FROM employee_work_experiences we
  WHERE we.employee_id = p_employee_id
    AND we.on_going IS TRUE
  ORDER BY we.date_start DESC NULLS LAST, we.id DESC
  LIMIT 1;

  UPDATE employees e
  SET
    highest_academic_degree = v_highest_academic_degree,
    speciality = v_speciality,
    current_workplace = v_current_workplace
  WHERE e.id = p_employee_id;
END;
$$;

CREATE OR REPLACE FUNCTION trg_refresh_employee_denormalized_fields()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
DECLARE
  v_employee_id BIGINT;
BEGIN
  IF TG_OP = 'DELETE' THEN
    v_employee_id := OLD.employee_id;
  ELSE
    v_employee_id := NEW.employee_id;
  END IF;

  PERFORM refresh_employee_denormalized_fields(v_employee_id);
  RETURN NULL;
END;
$$;

DROP TRIGGER IF EXISTS employee_degrees_refresh_employee_denormalized_fields_trg ON employee_degrees;
CREATE TRIGGER employee_degrees_refresh_employee_denormalized_fields_trg
AFTER INSERT OR UPDATE OR DELETE ON employee_degrees
FOR EACH ROW
EXECUTE FUNCTION trg_refresh_employee_denormalized_fields();

DROP TRIGGER IF EXISTS employee_work_experiences_refresh_employee_denormalized_fields_trg ON employee_work_experiences;
CREATE TRIGGER employee_work_experiences_refresh_employee_denormalized_fields_trg
AFTER INSERT OR UPDATE OR DELETE ON employee_work_experiences
FOR EACH ROW
EXECUTE FUNCTION trg_refresh_employee_denormalized_fields();

WITH latest_degree AS (
  SELECT DISTINCT ON (ed.employee_id)
    ed.employee_id,
    ed.speciality AS highest_academic_degree,
    ed.degree_level AS speciality
  FROM employee_degrees ed
  ORDER BY ed.employee_id, ed.date_end DESC NULLS LAST, ed.id DESC
),
latest_work AS (
  SELECT DISTINCT ON (we.employee_id)
    we.employee_id,
    we.workplace AS current_workplace
  FROM employee_work_experiences we
  WHERE we.on_going IS TRUE
  ORDER BY we.employee_id, we.date_start DESC NULLS LAST, we.id DESC
)
UPDATE employees e
SET
  highest_academic_degree = src.highest_academic_degree,
  speciality = src.speciality,
  current_workplace = src.current_workplace
FROM (
  SELECT
    COALESCE(ld.employee_id, lw.employee_id) AS employee_id,
    ld.highest_academic_degree,
    ld.speciality,
    lw.current_workplace
  FROM latest_degree ld
  FULL JOIN latest_work lw ON lw.employee_id = ld.employee_id
) src
WHERE e.id = src.employee_id;
