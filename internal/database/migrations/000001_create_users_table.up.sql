create table users (
    id serial primary key,
    email text not null,
    password text not null,
    created_at timestamp with time zone default current_timestamp
);