from dataclasses import dataclass
from random import randint

from gen_config import MAX_EARNINGS, MAX_TIER, MAX_YEAR_TEAM, MIN_EARNINGS, MIN_MMR, MIN_TIER, MIN_YEAR_TEAM, REGIONS_FILE, TEAMS_AMOUNT, TEAMS_FILE
from tools import email_gen, random_line, dataclass_to_csv, date_gen
from wonderwords import RandomWord

r = RandomWord()


@dataclass
class Team:
    name: str
    created_at: str
    email: str
    total_earnings: int
    region: str
    tier: int


def teams_gen():
    file = open("./dataset/"+TEAMS_FILE, "w", encoding="utf-8")
    for i in range(TEAMS_AMOUNT):
        team = Team(name=str(r.word(include_parts_of_speech=["adjectives"])).title()+" " + str(r.word(include_parts_of_speech=["nouns"])).title(),
                    created_at=date_gen(MIN_YEAR_TEAM, MAX_YEAR_TEAM),
                    email=email_gen(),
                    total_earnings=randint(MIN_EARNINGS, MAX_EARNINGS),
                    region=random_line(REGIONS_FILE),
                    tier=randint(MIN_TIER, MAX_TIER))
        file.write(dataclass_to_csv(team))
    file.close()
