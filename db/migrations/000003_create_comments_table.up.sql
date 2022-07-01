BEGIN;

create table comments(
    id SERIAL primary key,
    user_id int not NULL references public.users(id) on delete cascade on update cascade,
    photo_id int not NULL references public.photos(id) on delete cascade on update cascade,
    message text not null,
    created_at DATE,
    updated_at DATE
);

COMMIT;