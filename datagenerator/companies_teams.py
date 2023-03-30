from dataclasses import dataclass
from random import randint
from gen_config import COMPANIES_AMOUNT, COMPANIES_TEAMS_FILE, MAX_SPONSORS, MAX_YEAR_EVENT, MIN_SPONSORS, MIN_YEAR_EVENT, TEAMS_AMOUNT
from tools import dataclass_to_csv, date_gen


@dataclass
class CompanyTeam:
    company_id: int
    team_id: int
    contract_date: str
    contract_time: int


def companies_teams_gen():
    file = open("./dataset/"+COMPANIES_TEAMS_FILE, "w", encoding="utf-8")
    for i in range(TEAMS_AMOUNT):
        companies = [i+1 for i in range(COMPANIES_AMOUNT)]
        for j in range(randint(MIN_SPONSORS, MAX_SPONSORS)):
            obj = CompanyTeam(company_id=companies.pop(randint(0, len(companies)-1)),
                              team_id=i+1,
                              contract_date=date_gen(
                                  MIN_YEAR_EVENT, MAX_YEAR_EVENT),
                              contract_time=randint(1, 6)*6)
            file.write(dataclass_to_csv(obj))
    file.close()
