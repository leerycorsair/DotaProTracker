do
$do$
begin
   if exists (
      select from pg_catalog.pg_roles
      where  rolname = 'initial_login') then
      raise notice 'role "initial_login" already exists. skipping.';
   else
      create role initial_login with nosuperuser nocreatedb login password 'initial_login';
   end if;
end
$do$;
do
$do$
begin
   if exists (
      select from pg_catalog.pg_roles
      where  rolname = 'default_user') then
      raise notice 'role "default_user" already exists. skipping.';
   else
      create role default_user with nosuperuser nocreatedb login password 'default_user';
   end if;
end
$do$;
do
$do$
begin
   if exists (
      select from pg_catalog.pg_roles
      where  rolname = 'moderator') then
      raise notice 'role "moderator" already exists. skipping.';
   else
      create role moderator with nosuperuser nocreatedb login password 'moderator';
   end if;
end
$do$;

do
$do$
begin
   if exists (
      select from pg_catalog.pg_roles
      where  rolname = 'administrator') then
      raise notice 'role "administrator" already exists. skipping.';
   else
      create role administrator with nosuperuser nocreatedb login password 'administrator';
   end if;
end
$do$;

grant select, insert on table users to initial_login;
grant usage, select on sequence users_id_seq to initial_login;

grant select on table players to default_user;
grant select on table teams to default_user;
grant select on table tournaments to default_user;
grant select on table companies to default_user;
grant select on table teams_players to default_user;
grant select on table companies_teams to default_user;
grant select on table companies_tournaments to default_user;
grant select on table tournaments_teams to default_user;
grant select on table matches to default_user;
grant select on table match_perfomances to default_user;

grant select, insert, delete, update on table players to moderator;
grant select, insert, delete, update on table teams to moderator;
grant select, insert, delete, update on table tournaments to moderator;
grant select, insert, delete, update on table companies to moderator;
grant select, insert, delete, update on table teams_players to moderator;
grant select, insert, delete, update on table companies_teams to moderator;
grant select, insert, delete, update on table companies_tournaments to moderator;
grant select, insert, delete, update on table tournaments_teams to moderator;
grant select, insert, delete, update on table matches to moderator;
grant select, insert, delete, update on table match_perfomances to moderator;

   grant usage, select on sequence players_id_seq to moderator;
   grant usage, select on sequence teams_id_seq to moderator;
   grant usage, select on sequence tournaments_id_seq to moderator;
   grant usage, select on sequence companies_id_seq to moderator;
   grant usage, select on sequence teams_players_id_seq to moderator;
   grant usage, select on sequence companies_teams_id_seq to moderator;
   grant usage, select on sequence companies_tournaments_id_seq to moderator;
   grant usage, select on sequence tournaments_teams_id_seq to moderator;
   grant usage, select on sequence matches_id_seq to moderator;
   grant usage, select on sequence match_perfomances_id_seq to moderator;

   grant select, insert, delete, update on table players to administrator;
   grant select, insert, delete, update on table teams to administrator;
   grant select, insert, delete, update on table tournaments to administrator;
   grant select, insert, delete, update on table companies to administrator;
   grant select, insert, delete, update on table teams_players to administrator;
   grant select, insert, delete, update on table companies_teams to administrator;
   grant select, insert, delete, update on table companies_tournaments to administrator;
   grant select, insert, delete, update on table tournaments_teams to administrator;
   grant select, insert, delete, update on table matches to administrator;
   grant select, insert, delete, update on table match_perfomances to administrator;
   grant select, insert, delete, update on table users to administrator;

   grant usage, select on sequence players_id_seq to administrator;
   grant usage, select on sequence teams_id_seq to administrator;
   grant usage, select on sequence tournaments_id_seq to administrator;
   grant usage, select on sequence companies_id_seq to administrator;
   grant usage, select on sequence teams_players_id_seq to administrator;
   grant usage, select on sequence companies_teams_id_seq to administrator;
   grant usage, select on sequence companies_tournaments_id_seq to administrator;
   grant usage, select on sequence tournaments_teams_id_seq to administrator;
   grant usage, select on sequence matches_id_seq to administrator;
   grant usage, select on sequence match_perfomances_id_seq to administrator;
   grant usage, select on sequence users_id_seq to administrator;

