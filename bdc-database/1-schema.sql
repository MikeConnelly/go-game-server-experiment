CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "created_at" timestamp,
    "updated_at" timestamp,

    "email" varchar
);

CREATE TABLE "players" (
    "id" SERIAL PRIMARY KEY,
    "created_at" timestamp,
    "updated_at" timestamp,

    "user_id" int,
    "world_id" int,
    "alliance_id" int DEFAULT NULL,

    "name" varchar,
    "resource_worker_count" int,
    "resource_gold_count" int,
    "resource_lumber_count" int,
    "resource_mana_count" int,

    "castle_x_coord" int,
    "castle_y_coord" int,
    "castle_farm_stage" int,
    "castle_goldmine_stage" int,
    "castle_lumbermill_stage" int,
    "castle_manawell_stage" int,
    "castle_espionage_stage" int,
    "castle_armory_stage" int,
    "castle_stable_stage" int,
    "castle_heavy_stage" int
);

CREATE TABLE "worlds" (
    "id" SERIAL PRIMARY KEY,
    "created_at" timestamp,
    "updated_at" timestamp,

    "name" varchar,
    "active" boolean
);

CREATE TABLE "alliances" (
    "id" SERIAL PRIMARY KEY,
    "created_at" timestamp,
    "updated_at" timestamp,

    "name" varchar,
    "tag" varchar
);

ALTER TABLE "players" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "players" ADD FOREIGN KEY ("world_id") REFERENCES "worlds" ("id");

ALTER TABLE "players" ADD FOREIGN KEY ("alliance_id") REFERENCES "alliances" ("id");
