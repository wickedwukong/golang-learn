create keyspace golang with replication = {'class':'SimpleStrategy', 'replication_factor':1};

use golang;

create table employees (id UUID primary key, email VARCHAR, created TIMESTAMP, updated TIMESTAMP, deleted TIMESTAMP);

create table users ( id UUID primary key, name text );

insert into users (id,name) values (4fc59970-29ef-11e4-8c21-0800200c9a66, 'bob');

quit;
