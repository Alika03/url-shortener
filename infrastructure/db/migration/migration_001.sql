create table if not exists "url_shortener"
(
    id uuid primary key,
    code varchar(10) not null,
    full_url text not null,
    expired_at          timestamp(6) not null
);