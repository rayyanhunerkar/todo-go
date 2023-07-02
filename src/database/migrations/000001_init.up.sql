create extension if not exists "uuid-ossp";

create table if not exists "public"."users"(
    id uuid default uuid_generate_v4() primary key,
    username varchar(256) not null unique,
    password varchar(256) not null,
    first_name varchar(64) not null,
    last_name varchar(64) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create index user_name_idx on public.users (username);
