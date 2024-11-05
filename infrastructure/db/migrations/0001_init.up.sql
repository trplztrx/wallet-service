create type operation_type as enum('deposit', 'withdraw');

create table wallets (
    id uuid primary key,
    balance varchar(100) not null,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

create table transactions (
    id uuid primary key,
    wallet_id uuid references wallets(id),
    operationType operation_type not null,
    amount varchar(100) not null,
    created_at timestamp without time zone
);