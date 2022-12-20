alter table movies
    alter column version type text;

alter table movies
    alter column version set default 1;

alter table movies
    alter column version type integer using (version::integer);