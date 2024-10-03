create table badges (
    id serial primary key ,
    title varchar(255) not null,
    release_date date,
    description text,
    image_url varchar(255) not null,
    price int,
    category_id int,
    user_id int not null,
    collection_id int default 0,
    material varchar(50),
    weight int,
    height int,
    width int,
    thickness int,
    coverage varchar(30)
);