BEGIN;

create table photos(
    id SERIAL primary key,
    title text not null,
    caption text not null,
    photo_url text not null,
    user_id int NOT NULL references public.users(id) on delete cascade on update cascade,
    created_at DATE,
    updated_at DATE
);

COMMIT;