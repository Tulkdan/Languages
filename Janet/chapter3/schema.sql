create table people (id integer primary key, name text not null);

create table grudges (
    id integer primary key,
    holder integer not null,
    against integer not null,
    reason text not null,
    foreign key (holder) references people (id),
    foreign key (against) references people (id)
);
