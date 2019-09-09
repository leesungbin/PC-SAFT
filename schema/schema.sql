CREATE TABLE component (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  mw float(8) NOT NULL,
  Tc float(8) NOT NULL,
  Pc float(8) NOT NULL,
  omega float(8) NOT NULL,
  Tb float(8) NOT NULL,
  m float(8) NOT NULL,
  sig float(8) NOT NULL,
  eps float(8) NOT NULL,
  k float(8) NOT NULL,
  e float(8) NOT NULL,
  d float(8),
  x float(8)
);