create table if not exists users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

create table if not exists statistics
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    date    date                                        not null,
    views   int                                         not null,
    clicks  int                                         not null,
    cost    int                                         not null
);
