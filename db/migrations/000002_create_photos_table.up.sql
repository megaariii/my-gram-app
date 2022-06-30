BEGIN;

create table photos(
    id SERIAL primary key,
    title text not null,
    caption text not null,
    photo_url text not null,
    user_id int NOT NULL REFERENCES users(id),
    created_at DATE,
    updated_at DATE
);

COMMIT;