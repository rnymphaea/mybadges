create table if not exists users (
    id uuid primary key,
    email text not null,
    password text not null,
    created_at timestamp with time zone default current_timestamp
);

create table if not exists categories (
    id uuid primary key,
    category_name text not null
);

create table if not exists collections (
    id uuid primary key,
    collection_name text not null
);

create table if not exists badges (
    id uuid primary key,
    user_id uuid not null references users (id) on delete cascade,
    title varchar(255) not null,
    description text,
    image_url varchar(255),
    category_id uuid references categories (id) on delete set null,
    collection_id uuid references collections (id) on delete set null,
    material varchar(50),
    weight int,
    height int,
    width int,
    thickness int,
    coverage varchar(30),
    fastening varchar(30)
);