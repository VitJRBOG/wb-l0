DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS deliveries;

DROP TABLE IF EXISTS payments;

DROP TABLE IF EXISTS items;

CREATE TABLE orders (
    order_uid text NOT NULL PRIMARY KEY,
    track_number text,
    entry text,
    delivery text,
    payment text,
    items text[],
    locale text,
    internal_signature text,
    customer_id text,
    delivery_service text,
    shardkey text,
    sm_id smallint,
    date_created text,
    oof_shard text
);

CREATE TABLE deliveries (
    name text,
    phone text NOT NULL PRIMARY KEY,
    zip text,
    city text,
    address text,
    region text,
    email text
);

CREATE TABLE payments (
    transaction text NOT NULL PRIMARY KEY,
    request_id text,
    currency text,
    provider text,
    amount integer,
    payment_dt bigint,
    bank text,
    delivery_cost integer,
    goods_total smallint,
    custom_fee smallint
);

CREATE TABLE items (
    chrt_id bigint,
    track_number text,
    price integer,
    rid text NOT NULL PRIMARY KEY,
    name text,
    sale integer,
    size text,
    total_price integer,
    nm_id bigint,
    brand text,
    status smallint
);