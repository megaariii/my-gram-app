begin;

create table users(
    id SERIAL primary key,
    username varchar(50) unique not null,
    email varchar(50) unique not null,
    password TEXT not null,
    age int not null,
    created_at DATE,
    updated_at DATE
);

commit;