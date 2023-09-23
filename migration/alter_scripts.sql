-- USER
CREATE table pl_user (
    id SERIAL primary KEY,
    member_id UUID UNIQUE,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(100) NOT NULL,
    email VARCHAR(200),
    mobile VARCHAR(20) NOT NULL UNIQUE,
    status VARCHAR(30) NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by uuid,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by uuid,
    delete_time TIMESTAMP
);
