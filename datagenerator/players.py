from gen_config import COUNTRIES_FILE, HEROES_FILE, MAX_MMR, MAX_YEAR_BIRTH, MIN_MMR, MIN_YEAR_BIRTH, PLAYERS_AMOUNT, PLAYERS_FILE, ROLES_FILE
from dataclasses import dataclass
from nickname_generator import generate
from names import get_full_name
from tools import dataclass_to_csv, date_gen, random_line
import random


@dataclass
class Player:
    nickname: str
    realname: str
    birthdate: str
    country: str
    mmr: int
    role: str
    signature_hero: str


def players_gen():
    file = open("./dataset/"+PLAYERS_FILE, "w", encoding="utf-8")
    for i in range(PLAYERS_AMOUNT):
        player = Player(nickname=generate().lower(),
                        realname=get_full_name(),
                        birthdate=date_gen(MIN_YEAR_BIRTH, MAX_YEAR_BIRTH),
                        country=random_line(COUNTRIES_FILE),
                        mmr=random.randint(MIN_MMR, MAX_MMR),
                        role=random_line(ROLES_FILE),
                        signature_hero=random_line(HEROES_FILE))
        file.write(dataclass_to_csv(player))
    file.close()

