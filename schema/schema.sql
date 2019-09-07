CREATE TABLE IF NOT EXISTS pcsaft_migration (
  id INTEGER PRIMARY KEY,
  created TIMESTAMPTZ NOT NULL
);

DO $$
  DECLARE
    migration_id INTEGER;
  BEGIN
    IF( SELECT count(*) = 0 FROM pcsaft_migration WHERE id=1) THEN
      CREATE TABLE component (
        name TEXT NOT NULL,
        mw float(8) NOT NULL,
        omega float(8) NOT NULL,
        Tc float(8) NOT NULL,
        Pc float(8) NOT NULL,
        Tb float(8) NOT NULL,
        m float(8) NOT NULL,
        sig float(8) NOT NULL,
        eps float(8) NOT NULL,
        k float(8) NOT NULL,
        e float(8) NOT NULL,
        d float(8) NOT NULL,
        x float(8) NOT NULL
      );
      INSERT INTO pcsaft_migration VALUES (1, now());
    END IF;
  END;
$$