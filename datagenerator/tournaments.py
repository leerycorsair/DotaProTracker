from dataclasses import dataclass
from random import choice, randint
from gen_config import CITIES_FILE, MAX_DURATION, MAX_PRIZE, MAX_TIER, MAX_YEAR_EVENT, MIN_PRIZE, MIN_TIER, MIN_YEAR_EVENT, TOURNAMENTS_AMOUNT, TOURNAMENTS_FILE

from tools import email_gen, random_line, dataclass_to_csv, date_gen
from wonderwords import RandomWord

r = RandomWord()


@dataclass
class Tournament:
    name: str
    tier: int
    prize_pool: int
    start_date: str
    duration: int
    dpc_points: int
    location: str


def tournaments_gen():
    file = open("./dataset/"+TOURNAMENTS_FILE, "w", encoding="utf-8")
    for i in range(TOURNAMENTS_AMOUNT):
        tournament = Tournament(name=str(r.word(include_parts_of_speech=["adjectives"])).title()+" " + choice(["Dota Championship", "Major", "Minor", "Ladder", "Dreamhack"]),
                                tier=randint(MIN_TIER, MAX_TIER),
                                prize_pool=randint(
                                    MIN_PRIZE, MAX_PRIZE)//1000*1000,
                                start_date=date_gen(
                                    MIN_YEAR_EVENT, MAX_YEAR_EVENT),
                                duration=randint(1, MAX_DURATION),
                                dpc_points=randint(300, 2500)//100*100,
                                location=random_line(CITIES_FILE))
        file.write(dataclass_to_csv(tournament))
    file.close()
