

from dataclasses import dataclass
import itertools
import random
from secrets import choice
from telnetlib import TM

from gen_config import HEROES_FILE, MATCHES_FILE, MATCHES_PERFOMANCES_FILE, MAX_BLD, MAX_DMG, MAX_DN, MAX_GPM, MAX_HEAL, MAX_KDA, MAX_LH, MAX_MATCH_DURATION, MAX_NETWORTH, MAX_XPM, MIN_BLD, MIN_DMG, MIN_DN, MIN_GPM, MIN_HEAL, MIN_KDA, MIN_LH, MIN_MATCH_DURATION, MIN_NETWORTH, MIN_XPM
from tools import dataclass_to_csv

hero_pool = list(open(HEROES_FILE, "r", encoding="utf-8"))


@dataclass
class Match:
    tournament_id: int
    r_team_id: int
    d_team_id: int
    duration: int
    winner: bool


@dataclass
class MatchPerfomance:
    match_id: int
    player_id: int
    team: bool
    hero: str
    kills: int
    deaths: int
    assists: int
    networth: int
    gpm: int
    xpm: int
    dmg: int
    heal: int
    bld: int


def matches_gen(tournaments_teams, teams_players):
    tmp = open("./dataset/"+MATCHES_PERFOMANCES_FILE, "w", encoding="utf-8")
    tmp.close()

    file = open("./dataset/"+MATCHES_FILE, "w", encoding="utf-8")
    match_id = 1
    for tournament_id, teams_pool in tournaments_teams.items():
        for i in range(len(teams_pool)):
            for j in range(i+1, len(teams_pool)):
                obj = Match(tournament_id=tournament_id,
                            r_team_id=teams_pool[i],
                            d_team_id=teams_pool[j],
                            duration=random.randint(
                                MIN_MATCH_DURATION*60, MAX_MATCH_DURATION*60),
                            winner=choice([True, False]))
                match_info_gen(
                    match_id, teams_players[i+1], teams_players[j+1])
                match_id += 1
                file.write(dataclass_to_csv(obj))
    file.close()


def match_info_gen(match_id, r_players, d_players):
    h_pool = hero_pool.copy()
    i = 0
    file = open("./dataset/"+MATCHES_PERFOMANCES_FILE, "a", encoding="utf-8")
    for player_id in itertools.chain(r_players, d_players):
        obj = MatchPerfomance(match_id=match_id,
                              player_id=player_id,
                              team=i < 5,
                              hero=h_pool.pop(random.randint(
                                  0, len(h_pool)-1)).rstrip(),
                              kills=random.randint(MIN_KDA, MAX_KDA),
                              deaths=random.randint(MIN_KDA, MAX_KDA),
                              assists=random.randint(MIN_KDA, MAX_KDA),
                              networth=random.randint(
                                  MIN_NETWORTH, MAX_NETWORTH),
                              gpm=random.randint(MIN_GPM, MAX_GPM),
                              xpm=random.randint(MIN_XPM, MAX_XPM),
                              dmg=random.randint(MIN_DMG, MAX_DMG),
                              heal=random.randint(MIN_HEAL, MAX_HEAL),
                              bld=random.randint(MIN_BLD, MAX_BLD))
        i += 1
        file.write(dataclass_to_csv(obj))
    file.close()
