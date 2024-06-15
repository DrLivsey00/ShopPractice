Create table users(
    user_id serial primary key,
    user_name text unique not null ,
    user_email text unique not null,
    user_password text not null,
    user_bio text,
    user_access_level int default 0 not null ,
    created_at timestamp not null,
    updated_at timestamp not null
)