from gen_config import COMPANIES_AMOUNT, COMPANIES_TOURNAMENTS_FILE, MAX_DEPOSIT, MAX_SPONSORS, MAX_YEAR_EVENT, MIN_DEPOSIT, MIN_SPONSORS, MIN_YEAR_EVENT, PLAYERS_AMOUNT, PLAYERS_PER_TEAM, TEAM_PLAYER_FILE, TEAMS_AMOUNT, TOURNAMENTS_AMOUNT
from tools import dataclass_to_csv, date_gen
import random
from dataclasses import dataclass


@dataclass
class CompanyTournament:
    company_id: int
    tournament_id: int
    deposit: int


def companies_tournaments_gen():
    file = open("./dataset/"+COMPANIES_TOURNAMENTS_FILE, "w", encoding="utf-8")
    for i in range(TOURNAMENTS_AMOUNT):
        companies = [i+1 for i in range(COMPANIES_AMOUNT)]
        for j in range(random.randint(MIN_SPONSORS, MAX_SPONSORS)):
            obj = CompanyTournament(company_id=companies.pop(random.randint(0, len(companies)-1)),
                                    tournament_id=i+1,
                                    deposit=random.randint(MIN_DEPOSIT, MAX_DEPOSIT)//1000*1000)
            file.write(dataclass_to_csv(obj))
    file.close()
