CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists participants cascade;
create table participants
(
    id           uuid primary key default uuid_generate_v4(),
    name         varchar(128),
    category      integer,
    gender        boolean,
    birthdate timestamp,
    coach_name      varchar(128)
);

drop table if exists ratings cascade;
create table ratings
(
    id               uuid primary key default uuid_generate_v4(),
    name varchar(256),
    class             int,
    blowout_cnt int
);

drop table if exists crews cascade;
create table crews
(
    id           uuid primary key default uuid_generate_v4(),
    rating_id    uuid   references ratings (id)  ON DELETE CASCADE,
    class         int,
    sail_num      int
);

drop table if exists races cascade;
create table races
(
    id            uuid primary key default uuid_generate_v4(),
    rating_id    uuid   references ratings (id)  ON DELETE CASCADE,
    number serial,
    class int,
    date timestamp
);

drop table if exists judges cascade;
create table judges
(
    id      uuid primary key default uuid_generate_v4(),
    name varchar(128) ,
    login varchar(64),
    password text,
    role int,
    post text
);

drop table if exists protests cascade;
create table protests
(
    id       uuid primary key default uuid_generate_v4(),
    race_id uuid references races (id) ON DELETE CASCADE,
    rating_id  uuid references ratings (id) ON DELETE CASCADE,
    judge_id  uuid references judges (id) ON DELETE CASCADE,
    rule_num int ,
    review_date timestamp,
    status int,
    comment text
);

drop table if exists crew_protest cascade;
create table crew_protest
(
    id      uuid primary key default uuid_generate_v4(),
    crew_id uuid references crews (id) ON DELETE CASCADE,
    protest_id  uuid references protests (id) ON DELETE CASCADE,
    crew_status int
);

drop table if exists crew_race cascade;
create table crew_race
(
    id      uuid primary key default uuid_generate_v4(),
    crew_id uuid references crews (id) ON DELETE CASCADE,
    race_id  uuid references races (id) ON DELETE CASCADE,
    points int,
    spec_circumstance int
);

drop table if exists participant_crew cascade;
create table participant_crew
(
    id       uuid primary key default uuid_generate_v4(),
    participant_id uuid references participants (id) ON DELETE CASCADE,
    crew_id  uuid references crews (id) ON DELETE CASCADE,
    helmsman boolean,
    active boolean
);

drop table if exists judge_rating cascade;
create table judge_rating
(
    id      uuid primary key default uuid_generate_v4(),
    judge_id uuid references judges (id) ON DELETE CASCADE,
    rating_id  uuid references ratings (id) ON DELETE CASCADE
);

