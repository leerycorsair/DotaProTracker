from dataclasses import dataclass
from random import choice, randint
from gen_config import COMPANIES_AMOUNT, COMPANIES_FILE, COUNTRIES_FILE, MAX_REVENUE, MIN_REVENUE
from tools import dataclass_to_csv, random_line
from wonderwords import RandomWord

r = RandomWord()


@dataclass
class Company:
    name: str
    country: str
    website: str
    revenue: int
    industry: str


def companies_gen():
    file = open("./dataset/"+COMPANIES_FILE, "w",encoding="utf-8")
    for i in range(COMPANIES_AMOUNT):
        cname = str(r.word(include_parts_of_speech=["adjectives"])).title()+" " + str(r.word(include_parts_of_speech=["nouns"])).title() + " " + choice(["Inc.", "Company", "Team", "Group", "Cooperation", "Imternational"])
        company = Company(name=cname,
                          country=random_line(COUNTRIES_FILE),
                          website="".join(cname.split()).lower()+".com",
                          revenue=randint(MIN_REVENUE, MAX_REVENUE) // 100 * 100,
                          industry=str(r.word(include_parts_of_speech=["nouns"])).title())
        file.write(dataclass_to_csv(company))
    file.close()
