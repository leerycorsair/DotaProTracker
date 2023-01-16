drop table if exists players cascade;
drop table if exists teams cascade;
drop table if exists tournaments cascade;
drop table if exists companies cascade;
drop table if exists teams_players cascade;
drop table if exists companies_teams cascade;
drop table if exists companies_tournaments cascade;
drop table if exists tournaments_teams cascade;
drop table if exists matches cascade;
drop table if exists match_perfomances cascade;
drop table if exists users cascade;

create table "players" (
  "id" serial,
  "nickname" varchar,
  "realname" varchar,
  "birthdate" date,
  "country" varchar,
  "mmr" int,
  "role" varchar,
  "signature_hero" varchar
);

create table "teams" (
  "id" serial,
  "name" varchar,
  "created_at" date,
  "email" varchar,
  "total_earnings" int,
  "region" varchar,
  "tier" int
);

create table "tournaments" (
  "id" serial,
  "name" varchar,
  "tier" int,
  "prize_pool" int,
  "date_start" date,
  "duration" int,
  "dpc_points" int,
  "location" varchar
);

create table "companies" (
  "id" serial,
  "name" varchar,
  "country" varchar,
  "website" varchar,
  "revenue" int,
  "industry" varchar
);

create table "teams_players" (
  "id" serial,
  "team_id" int,
  "player_id" int,
  "contract_date" date,
  "contract_time" int
);

create table "companies_tournaments" (
  "id" serial,
  "company_id" int,
  "tournament_id" int,
  "deposit" int
);

create table "companies_teams" (
  "id" serial,
  "company_id" int,
  "team_id" int,
  "contract_date" date,
  "contract_time" int
);

create table "tournaments_teams" (
  "id" serial,
  "tournament_id" int,
  "team_id" int,
  "participation_type" varchar,
  "is_winner" boolean
);

create table "matches" (
  "id" serial,
  "tournament_id" int,
  "r_team_id" int,
  "d_team_id" int,
  "duration" int, 
  "winner" boolean
);

create table "users" (
  "id" serial,
  "name" varchar,
  "birthdate" date,
  "login" varchar, 
  "password" varchar,
  "email" varchar, 
  "privilege_level" int
);

create table "match_perfomances" (
  "id" serial,
  "match_id" int,
  "player_id" int, 
  "team" boolean,
  "hero" varchar,
  "kills" int,
  "deaths" int,
  "assists" int,
  "networth" int,
  "gpm" int,
  "xpm" int,
  "dmg" int,
  "heal" int,
  "bld" int
);