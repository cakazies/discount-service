CREATE TABLE IF NOT EXISTS "items" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "sku" varchar(50) NOT NULL,
    "name" varchar(200) NOT NULL,
    "price" float8,
    "currency" varchar(20),
    "qty" int,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp,
    "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "users" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "username" varchar(100) NOT NULL,
    "password" varchar(200) NOT NULL,
    "name" varchar(200) NOT NULL,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp,
    "deleted_at" timestamp
);

CREATE TABLE IF NOT EXISTS "orders" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "user_id" int NOT NULL,
    "item_id" int NOT NULL,
    "item_sku" varchar(50),
    "item_name" varchar(200),
    "item_price" float8,
    "item_qty" int,
    "discount" float8,
    "total" float8,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp,
    "deleted_at" timestamp
);


CREATE TABLE "promotion_items" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "item_id" int,
    "item_sku" string,
    "min_qty" int,
    "free_item" varchar,
    "discount" float,
    "is_cashback" boolean,
    "direct_cashback" varchar(100),
    "max_count" int,
    "detail" text,
    "is_active" bool,
    "is_abnormal" bool,
    "start_date" timestamp,
    "end_date" timestamp,
    "created_at" timestamp DEFAULT (now()),
    "updated_at" timestamp,
    "deleted_at" timestamp
);

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

ALTER TABLE "promotion_items" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("id");

insert into items (sku, name, price, currency, qty)
values ('120P90', 'Google Home', 49.99, 'USD', 10),
		('43N23P', 'Macbook Pro', 5399.99, 'USD', 5),
		('A304SD', 'Alexa Speaker', 109.55, 'USD', 10),
		('234234', 'RaspBerry Pi B', 30, 'USD', 2)