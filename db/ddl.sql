create database kpi3;

create user balance_admin with password '3kpi';
grant all privileges on database kpi3 to balance_admin;
alter database kpi3 owner to balance_admin;

drop schema if exists lab3 cascade;
create schema lab3;

set search_path to lab3;

drop table if exists machines;
create table machines(
  id serial not null,
  balancer_id integer not null,
  usable boolean not null default false
);
alter table machines add constraint pk_machines primary key(id);

drop table if exists balancers;
create table balancers(
    id serial not null,
    name varchar(255) not null
);
alter table balancers add constraint pk_balancers primary key(id);

alter table machines add constraint fk_machines_balancers foreign key(balancer_id) references balancers(id);
