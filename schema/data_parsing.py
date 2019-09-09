from os import listdir, getenv
from os.path import isfile, join
import psycopg2

dbname=getenv("POSTGRES_DBNAME")
conn = psycopg2.connect("dbname=%s user=postgres"%dbname)
print(conn)
cur = conn.cursor()

dataPath = getenv("DEFAULT_DATA_PATH")
files = [f for f in listdir(dataPath) if isfile(join(dataPath, f))]

for filename in files:
  with open(dataPath+filename,"r") as f:
    name=f.readline()
    f.readline()
    mw, Tc, Pc, omega, Tb = map(float, f.readline().split())
    f.readline()
    props2=f.readline().split() # pcsaft data
    if len(props2) == 5 :
      m, sig, eps, k, e = map(float, props2)
      d, x = [None, None]
    elif len(props2) == 7 :
      m, sig, eps, k, e, d, x = map(float, props2)
    if d == None and x == None:
      cur.execute("""
      INSERT INTO component
      VALUES (name,\"%s\",mw, %f, Tc, %f, Tc, %f, omega, %f, Tb, %f, m, %f, sig, %f, eps, %f, k, %f, e, %f)"""
      %(name, mw, Tc, Pc, omega, Tb, m, sig, eps, k, e))
    else:
      cur.execute("""INSERT INTO component
      VALUES (name,\"%s\",mw, %f, Tc, %f, Tc, %f, omega, %f, Tb, %f, m, %f, sig, %f, eps, %f, k, %f, e, %f, d, %f, x, %f);"""
      %(name, mw, Tc, Pc, omega, Tb, m, sig, eps, k, e, d, x))
      conn.commit()

cur.close()
conn.close()