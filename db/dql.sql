create or replace function get_balancers_info()
returns table(id integer, count bigint)
language sql as $$
    select distinct b.id,
    count(*) over(partition by b.id)
    from balancers b
    join machines m on b.id = m.balancer_id
    order by b.id;
$$;

select * from get_balancers_info();

create function get_usable_machines()
returns table(balancer_id integer, id integer)
language sql as $$
    select balancer_id, id from machines
    where state = true
$$;

select * from get_usable_machines();

create or replace procedure update_machine(machine_id integer, s bool)
language sql as $$
    update machines
    set state = s
    where id = machine_id
$$;

call update_machine(1, false);

alter procedure update_machine(integer, boolean) owner to balance_admin;
alter function get_balancers_info() owner to balance_admin;
alter function get_usable_machines() owner to balance_admin;

alter sequence balancers_id_seq owner to balance_admin;
alter sequence machines_id_seq owner to balance_admin;
