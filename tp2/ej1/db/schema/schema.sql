create table users (
    id serial primary key,
    name varchar(50) not null unique,
    email varchar(100) not null unique,
    created_at timestamp with time zone default current_timestamp
);