

from dataclasses import dataclass
from gen_config import TEAMS_AMOUNT, TOURNAMENTS_AMOUNT, TOURNAMENTS_TEAMS_FILE
from random import choice, randint
from tools import dataclass_to_csv


@dataclass
class TournamentTeam:
    tournament_id: int
    team_id: int
    participation_type: str
    is_winner: bool


def tournaments_teams_gen():
    file = open("./dataset/"+TOURNAMENTS_TEAMS_FILE, "w", encoding="utf-8")
    t_t_info = []
    for i in range(TOURNAMENTS_AMOUNT):
        teams = [i+1 for i in range(TEAMS_AMOUNT)]
        teams_in_tournament = []
        for j in range(choice([8, 16,32])):
            teams_in_tournament.append(teams.pop(randint(0, len(teams)-1)))
            obj = TournamentTeam(tournament_id=i+1,
                                 team_id=teams_in_tournament[-1],
                                 participation_type=choice(
                                     ["invite", "qualification"]),
                                 is_winner=True and j == 0)
            file.write(dataclass_to_csv(obj))
        t_t_info.append(teams_in_tournament)
    file.close()

    zip_iterator = zip([i+1 for i in range(TOURNAMENTS_AMOUNT)], t_t_info)
    tournaments_teams_dict = dict(zip_iterator)
    return tournaments_teams_dict

