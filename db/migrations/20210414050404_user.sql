-- migrate:up
create table users (
  id integer not null primary key,
  name varchar(255),
  email varchar(255) not null,
  created_at datetime,
  updated_at datetime,
  deleted_at datetime
);

-- migrate:down
drop table users;