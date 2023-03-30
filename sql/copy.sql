copy players(nickname, realname, birthdate, country, mmr, role, signature_hero)
    from 'c:/gencsv/players.csv'
    null as ''
    csv;

copy teams(name, created_at, email, total_earnings, region, tier)
    from 'c:/gencsv/teams.csv'
    null as ''
    csv;

copy tournaments(name, tier, prize_pool, date_start, duration, dpc_points, location)
    from 'c:/gencsv/tournaments.csv'
    null as ''
    csv;

copy companies(name, country, website, revenue, industry)
    from 'c:/gencsv/companies.csv'
    null as ''
    csv;

copy teams_players(team_id, player_id, contract_date, contract_time)
    from 'c:/gencsv/teams_players.csv'
    null as ''
    csv;

copy companies_tournaments(company_id, tournament_id, deposit)
    from 'c:/gencsv/companies_tournaments.csv'
    null as ''
    csv;

copy companies_teams(company_id, team_id, contract_date, contract_time)
    from 'c:/gencsv/companies_teams.csv'
    null as ''
    csv;

copy tournaments_teams(tournament_id, team_id, participation_type, is_winner)
    from 'c:/gencsv/tournaments_teams.csv'
    null as ''
    csv;

copy matches(tournament_id, r_team_id, d_team_id, duration, winner)
    from 'c:/gencsv/matches.csv'
    null as ''
    csv;

copy match_perfomances(match_id, player_id, team, hero, kills, deaths, assists, networth, gpm, xpm, dmg, heal, bld)
    from 'c:/gencsv/match_perfomances.csv'
    null as ''
    csv;
