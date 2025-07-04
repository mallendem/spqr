CREATE DISTRIBUTION ds1 COLUMN TYPES integer;
CREATE DISTRIBUTION ds2 COLUMN TYPES integer;

ALTER DISTRIBUTION ds1 ATTACH RELATION a DISTRIBUTION KEY a_id;
SHOW relations;

ALTER DISTRIBUTION ds1 ALTER RELATION a DISTRIBUTION KEY another;
SHOW relations;

ALTER DISTRIBUTION ds1 ALTER RELATION a DISTRIBUTION KEY another SCHEMA schema_name;
SHOW relations;

ALTER DISTRIBUTION ds1 ADD DEFAULT SHARD sh_no_exist;

ALTER DISTRIBUTION ds1 ADD DEFAULT SHARD sh1;
SHOW distributions;

CREATE DISTRIBUTED RELATION aa DISTRIBUTION KEY another2 SCHEMA schema_name2 IN ds2;
SHOW relations;

DROP DISTRIBUTION ALL CASCADE;
DROP KEY RANGE ALL;
