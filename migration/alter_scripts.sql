-- USER
CREATE table if not EXISTS "user" (
    id SERIAL primary KEY,
    member_id UUID UNIQUE,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(100) NOT NULL,
    email VARCHAR(200),
    mobile VARCHAR(20) NOT NULL UNIQUE,
    status VARCHAR(30) NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID,
    delete_time TIMESTAMP
);

-- GROUP
CREATE TABLE if not exists "group" (
    id SERIAL primary key,
    group_id UUID UNIQUE,
    name VARCHAR(255) not NULL,
    description VARCHAR(1000),
    "size" smallint,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID
);
-- GROUP MEMBER
CREATE table if not EXISTS group_member (
    id SERIAL,
    group_id UUID NOT NULL,
    member_id UUID NOT NULL,
    status VARCHAR(10) NOT NULL,
    is_admin bool DEFAULT false,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by uuid,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by uuid,
    delete_time TIMESTAMP
);

ALTER TABLE group_member ADD CONSTRAINT uniq_member_group UNIQUE(group_id, member_id);
ALTER TABLE group_member ADD CONSTRAINT fk_group_member_group FOREIGN KEY(group_id) REFERENCES "group"(group_id);
ALTER TABLE group_member ADD CONSTRAINT fk_group_pl_user FOREIGN KEY(member_id) REFERENCES "user"(member_id);
CREATE INDEX IF NOT EXISTS group_member_ids ON group_member (group_id, member_id);

-- EVENT
CREATE TABLE if not exists "event" (
    id SERIAL,
    event_id UUID UNIQUE,
    group_id UUID NOT NULL,
    venue_id UUID NOT NULL,
    name VARCHAR(500) NOT NULL,
    type VARCHAR(10) NOT NULL,
    status VARCHAR(10) NOT NULL,
    params JSONB NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by uuid,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by uuid,
    delete_time TIMESTAMP
);

-- EVENT MEMBER
CREATE TABLE if not exists event_member (
    id SERIAL,
    event_id UUID NOT NULL,
    group_id UUID NOT NULL,
    member_id UUID NOT NULL,
    action VARCHAR(10) NOT NULL,
    status VARCHAR(10) NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by uuid,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by uuid,
    delete_time TIMESTAMP
);

ALTER TABLE event_member ADD CONSTRAINT uniq_member_event UNIQUE(event_id, member_id);
CREATE INDEX IF NOT EXISTS event_member_index ON event_member (event_id, member_id);
alter table event_member add CONSTRAINT fk_event_id foreign key(event_id) references "event"(event_id);
alter table event_member add constraint fk_event_member_id foreign key(member_id) references "user"(member_id);