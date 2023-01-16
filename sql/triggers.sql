create or replace function make_deposit()
returns trigger
as
    $$
    begin
        update tournaments
        set prize_pool = prize_pool + new.deposit 
        where id = new.tournament_id;
        return new;
    end;
    $$ 
language plpgsql;

drop trigger if exists make_deposit_trigger on companies_tournaments;
create trigger make_deposit_trigger after insert on companies_tournaments
for row execute procedure make_deposit();


create or replace function make_win()
returns trigger
as
    $$
    declare
        tmp int = 0;
    begin
        if (new.is_winner = true) then 
            select into tmp prize_pool from tournaments_teams join tournaments
            on tournaments.id = tournaments_teams.id
            where tournaments.id = new.tournament_id;

            update teams
            set total_earnings = total_earnings + tmp
            where teams.id = new.team_id;
        end if;
        return new;
    end;
    $$ 
language plpgsql;

drop trigger if exists make_win on tournaments_teams;
create trigger make_win after insert on tournaments_teams
for row execute procedure make_win();

