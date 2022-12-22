CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
alter table movies
    alter column version type text;

alter table movies
    alter column version set default uuid_generate_v4();

alter table movies
    alter column version type uuid using version::uuid;
