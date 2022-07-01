BEGIN;

create table social_media(
    id SERIAL primary key,
    name varchar(250) not null,
    social_media_url text not null,
    user_id int NOT NULL references public.users(id) on delete cascade on update cascade,
    created_at DATE,
    updated_at DATE
);

COMMIT;