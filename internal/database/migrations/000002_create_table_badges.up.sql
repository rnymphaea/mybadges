create table if not exists badges (
    id uuid primary key,
    user_id uuid not null,
    title varchar(255) not null,
    description text,
    release_date date,
    image_url varchar(255),
    price int,
    category_id uuid,
    collection_id uuid,
    material varchar(50),
    weight int,
    height int,
    width int,
    thickness int,
    coverage varchar(30),
    fastening varchar(30)
);