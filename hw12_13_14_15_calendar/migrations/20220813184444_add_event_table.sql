-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table event
(
    id                    bigserial primary key,
    title                 text      not null,
    notification_date     timestamp not null,
    start_date            timestamp not null,
    end_date              timestamp,
    notification_interval bigint,
    description           text,
    owner_id              bigint    not null,
    created_at            timestamp not null default now(),
    updated_at            timestamp
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table event;
