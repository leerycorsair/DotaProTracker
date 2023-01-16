
insert into players(nickname, realname, birthdate, country, mmr, role, signature_hero) values
('chexybatup','Barbra Collier','1995-8-18','Dominica',5002,'Full Support','Enigma'),
('jyzoky','Michael Lastrapes','1989-6-21','Australia',9119,'Support','Clinkz');

insert into teams(name, created_at, email, total_earnings, region, tier) values
('Abounding Hall','2005-10-4','coolcommission@mail.com',222467,'Australia',3),
('Cool Team','2005-10-4','coolcommission@mail.com',2267,'Germany',3);

insert into tournaments(name, tier, prize_pool, date_start, duration, dpc_points, location) values
('Verdant Minor',1,19811000,'2018-1-27',6,2100,'College Station');

insert into companies(name, country, website, revenue, industry) values
('Historical Pearl Cooperation','Iran','historicalpearlcooperation.com',146600,'Click');

insert into teams_players(team_id, player_id, contract_date, contract_time) values
(1,1,'2010-6-4',18);

insert into companies_tournaments(company_id, tournament_id, deposit) values
(1,1,764000);

insert into companies_teams(company_id, team_id, contract_date, contract_time) values
(1,1,'2014-5-5',36);

insert into tournaments_teams(tournament_id, team_id, participation_type, is_winner) values
(1,1,'qualification',True);

insert into matches(tournament_id, r_team_id, d_team_id, duration, winner) values
(1,1,2,4429,False);

insert into match_perfomances(match_id, player_id, team, hero, kills, deaths, assists, networth, gpm, xpm, dmg, heal, bld) values
(1,1,True,'Terrorblade',28,12,2,36993,711,889,76442,11702,1870);

insert into users(name, birthdate, login, password, email, privilege_level) values
('Cool Guy', '2001-01-01', 'coolguy', 'qwertyasdsdfh', 'email@mail.com', 1);

