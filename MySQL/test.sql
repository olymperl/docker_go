FROM mysql:5.6
COPY test.sql /docker-entrypoint-initdb.d/test.sql

root@kanet:~/docker/MySQL# cat test.sql 
Use test_db;

Create table `test_tb` (id int, name varchar(20));

Insert into test_tb (id, name) values (1, 'test-user');