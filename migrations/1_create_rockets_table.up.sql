create table rockets if not exists {
    id serial not null primary key,
    type varchar (50),
    name varchar (50)
};