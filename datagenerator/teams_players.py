from gen_config import MAX_YEAR_EVENT, MIN_YEAR_EVENT, PLAYERS_AMOUNT, PLAYERS_PER_TEAM, TEAM_PLAYER_FILE, TEAMS_AMOUNT
from tools import dataclass_to_csv, date_gen
import random
from dataclasses import dataclass

@dataclass
class TeamPlayer:
    team_id: int
    player_id: int
    contract_date: str
    contract_time: int


def teams_players_gen():
    file = open("./dataset/"+TEAM_PLAYER_FILE, "w", encoding="utf-8")
    players = [i+1 for i in range(PLAYERS_AMOUNT)]
    t_p_info = []
    for i in range(TEAMS_AMOUNT):
        players_in_team = [] 
        for j in range(PLAYERS_PER_TEAM):
            players_in_team.append(players.pop(random.randint(0, len(players)-1)))
            obj = TeamPlayer(team_id = i+1,
                            player_id = players_in_team[-1],
                            contract_date = date_gen(MIN_YEAR_EVENT, MAX_YEAR_EVENT),
                            contract_time = random.randint(1,6)*6)
            file.write(dataclass_to_csv(obj))
        t_p_info.append(players_in_team)
    file.close()

    zip_iterator = zip([i+1 for i in range(TEAMS_AMOUNT)], t_p_info)
    team_players_dict = dict(zip_iterator)
    return team_players_dict

