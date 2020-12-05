create or replace function get_balancers_id() returns setof integer as '
    select id from balancers
' language sql;
alter function get_balancers_id() owner to balance_admin;

select get_balancers_id();

create or replace function get_machines_quantity(balancer integer) returns bigint as '
    select count(*) from machines
    where balancer_id = balancer
' language sql;
alter function get_machines_quantity(integer) owner to balance_admin;

select get_machines_quantity(3);

create or replace function get_usable_machines(balancer integer) returns setof int as '
    select id from machines
    where balancer_id = balancer and state = true
' language sql;
alter function get_usable_machines(integer) owner to balance_admin;

select get_usable_machines(3);

create or replace procedure update_machine(machine_id integer, s bool) as '
    update machines
    set state = s
    where id = machine_id
' language sql;
alter procedure update_machine(integer, boolean) owner to balance_admin;

call update_machine(1, false);

alter sequence balancers_id_seq owner to balance_admin;
alter sequence machines_id_seq owner to balance_admin;
