CREATE KEYSPACE golang WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};
use golang


CREATE TABLE users (id uuid PRIMARY KEY, name varchar);
