BEGIN;

create table comments(
    id SERIAL primary key,
    user_id int not NULL references users(id) on delete cascade on update cascade,
    photo_id int not NULL references photos(id) on delete cascade on update cascade,
    message text not null,
    created_at DATE,
    updated_at DATE
);

COMMIT;