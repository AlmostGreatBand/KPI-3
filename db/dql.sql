create or replace function get_machines_quantity(balancer integer) returns bigint as '
    select count(*) from machines
    where balancer_id = balancer;
' language sql;

select get_machines_quantity(3);

create or replace function get_usable_machines(balancer integer) returns setof integer as '
    select id from machines
    where balancer_id = balancer and usable = true
' language sql;

select get_usable_machines(3);