alter table "players" add constraint "players_id" primary key ("id"); 
alter table "teams" add constraint "teams_id" primary key ("id");
alter table "tournaments" add constraint "tournaments_id" primary key ("id");
alter table "companies" add constraint "companies_id" primary key ("id");
alter table "teams_players" add constraint "teams_players_id" primary key ("id");
alter table "companies_tournaments" add constraint "companies_tournaments_id" primary key ("id");
alter table "companies_teams" add constraint "companies_teams_id" primary key ("id");
alter table "tournaments_teams" add constraint "tournaments_teams_id" primary key ("id");
alter table "matches" add constraint "matches_id" primary key ("id");
alter table "users" add constraint "users_id" primary key ("id");
alter table "match_perfomances" add constraint "matches_perfomances_id" primary key ("id");

alter table "teams_players" add foreign key ("team_id") references "teams" ("id") on delete cascade;
alter table "teams_players" add foreign key ("player_id") references "players" ("id") on delete cascade;
alter table "companies_tournaments" add foreign key ("company_id") references "companies" ("id") on delete cascade;
alter table "companies_tournaments" add foreign key ("tournament_id") references "tournaments" ("id") on delete cascade;
alter table "companies_teams" add foreign key ("company_id") references "companies" ("id") on delete cascade;
alter table "companies_teams" add foreign key ("team_id") references "teams" ("id") on delete cascade;
alter table "tournaments_teams" add foreign key ("tournament_id") references "tournaments" ("id") on delete cascade;
alter table "tournaments_teams" add foreign key ("team_id") references "teams" ("id") on delete cascade;
alter table "matches" add foreign key ("tournament_id") references "tournaments" ("id") on delete cascade;
alter table "matches" add foreign key ("r_team_id") references "teams" ("id") on delete cascade;
alter table "matches" add foreign key ("d_team_id") references "teams" ("id") on delete cascade;

alter table "match_perfomances" add foreign key ("match_id") references "matches" ("id") on delete cascade;
alter table "match_perfomances" add foreign key ("player_id") references "players" ("id") on delete cascade;

alter table "players" alter column "nickname" set not null;
alter table "players" alter column "realname" set not null;
alter table "players" alter column "birthdate" set not null;
alter table "players" alter column "country" set not null;
alter table "players" alter column "mmr" set not null;
alter table "players" alter column "role" set not null;
alter table "players" alter column "signature_hero" set not null;
alter table "players" add constraint "players_mmr_check" check (mmr > 0);

alter table "teams" alter column "name" set not null;
alter table "teams" alter column "created_at" set not null;
alter table "teams" alter column "email" set not null;
alter table "teams" alter column "total_earnings" set not null;
alter table "teams" alter column "region" set not null;
alter table "teams" alter column "tier" set not null;
alter table "teams" add constraint "teams_total_earnings_check" check (total_earnings > 0);
alter table "teams" add constraint "teams_tier_check" check (tier > 0 and tier < 5);

alter table "tournaments" alter column "name" set not null;
alter table "tournaments" alter column "tier" set not null;
alter table "tournaments" alter column "prize_pool" set not null;
alter table "tournaments" alter column "date_start" set not null;
alter table "tournaments" alter column "duration" set not null;
alter table "tournaments" alter column "dpc_points" set not null;
alter table "tournaments" alter column "location" set not null;
alter table "tournaments" add constraint "tournaments_tier_check" check (tier > 0 and tier < 5);
alter table "tournaments" add constraint "tournaments_prize_pool_check"check (prize_pool > 0);
alter table "tournaments" add constraint "tournaments_date_check" check (duration > 0);
alter table "tournaments" add constraint "tournaments_dpc_points_check"check (dpc_points >= 0);

alter table "companies" alter column "name" set not null;
alter table "companies" alter column "country" set not null;
alter table "companies" alter column "website" set not null;
alter table "companies" alter column "revenue" set not null;
alter table "companies" alter column "industry" set not null;
alter table "companies" add constraint "companies_revenue" check (revenue > 0);

alter table "teams_players" alter column "team_id" set not null;
alter table "teams_players" alter column "player_id" set not null;
alter table "teams_players" alter column "contract_date" set not null;
alter table "teams_players" alter column "contract_time" set not null;
alter table "teams_players" add constraint "teams_players_contract_time_check" check (contract_time > 0 and contract_time <= 36);

alter table "companies_tournaments" alter column "company_id" set not null;
alter table "companies_tournaments" alter column "tournament_id" set not null;
alter table "companies_tournaments" alter column "deposit" set not null;
alter table "companies_tournaments" add constraint "companies_tournaments_deposit_check" check (deposit > 0);

alter table "companies_teams" alter column "company_id" set not null;
alter table "companies_teams" alter column "team_id" set not null;
alter table "companies_teams" alter column "contract_date" set not null;
alter table "companies_teams" alter column "contract_time" set not null;
alter table "companies_teams" add constraint "companies_teams_contract_time_check" check (contract_time > 0 and contract_time <= 36);

alter table "tournaments_teams" alter column "tournament_id" set not null;
alter table "tournaments_teams" alter column "team_id" set not null;
alter table "tournaments_teams" alter column "participation_type" set not null;
alter table "tournaments_teams" alter column "is_winner" set not null;
alter table "tournaments_teams" add constraint "tournaments_teams_participation_type_check" check (participation_type in ('invite', 'qualification'));

alter table "matches" alter column "tournament_id" set not null;
alter table "matches" alter column "r_team_id" set not null;
alter table "matches" alter column "d_team_id" set not null;
alter table "matches" alter column "duration" set not null;
alter table "matches" alter column "winner" set not null;
alter table "matches" add constraint "matches_duration_check" check (duration > 0);

alter table "users" alter column "name" set not null;
alter table "users" alter column "birthdate" set not null;
alter table "users" alter column "login" set not null;
alter table "users" alter column "password" set not null;
alter table "users" alter column "email" set not null;
alter table "users" alter column "privilege_level" set not null;
alter table "users" add constraint "users_login" unique (login);
alter table "users" add constraint "users_lvl_check" check (privilege_level > 0 and privilege_level < 4);

alter table "match_perfomances" alter column "match_id" set not null;
alter table "match_perfomances" alter column "player_id" set not null;
alter table "match_perfomances" alter column "team" set not null;
alter table "match_perfomances" alter column "hero" set not null;
alter table "match_perfomances" alter column "kills" set not null;
alter table "match_perfomances" alter column "deaths" set not null;
alter table "match_perfomances" alter column "assists" set not null;
alter table "match_perfomances" alter column "networth" set not null;
alter table "match_perfomances" alter column "gpm" set not null;
alter table "match_perfomances" alter column "xpm" set not null;
alter table "match_perfomances" alter column "dmg" set not null;
alter table "match_perfomances" alter column "heal" set not null;
alter table "match_perfomances" alter column "bld" set not null;
alter table "match_perfomances" add constraint "match_perfomances_kills_check" check (kills >= 0);
alter table "match_perfomances" add constraint "match_perfomances_deaths_check" check (deaths >= 0);
alter table "match_perfomances" add constraint "match_perfomances_assists_check" check (assists >= 0);
alter table "match_perfomances" add constraint "match_perfomances_networth_check" check (networth > 0);
alter table "match_perfomances" add constraint "match_perfomances_gpm_check" check (gpm > 0);
alter table "match_perfomances" add constraint "match_perfomances_xpm_check" check (xpm > 0);
alter table "match_perfomances" add constraint "match_perfomances_dmg_check" check (dmg >= 0);
alter table "match_perfomances" add constraint "match_perfomances_heal_check" check (heal >= 0);
alter table "match_perfomances" add constraint "match_perfomances_bld_check" check (bld >= 0);