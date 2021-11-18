create database users_db;
create user pgUser with password 'pg_password';
grant all privileges on database users_db to pgUser;