create or replace function get_machines_quantity(balancer integer) returns bigint as '
    select count(*) from machines
    where balancer_id = balancer;
' language sql;

select get_machines_quantity(3);

create or replace function get_usable_machines(balancer integer) returns setof integer as '
    select id from machines
    where balancer_id = balancer and state = true
' language sql;

select get_usable_machines(3);

create or replace procedure update_machine(machine_id integer, s bool) as '
    update machines
    set state = s
    where id = machine_id
' language sql;

call update_machine(1, false);
