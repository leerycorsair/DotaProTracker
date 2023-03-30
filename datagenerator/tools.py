from wonderwords import RandomWord
import string
import random


def date_gen(min_year: int, max_year: int) -> str:
    return str(random.randint(min_year, max_year)) + "-" + str(random.randint(1, 12)) + "-" + str(random.randint(1, 28))


def random_line(filename: str) -> str:
    return random.choice(list(open(filename))).rstrip()


def dataclass_to_csv(dataclass) -> str:
    string = ""
    for field in dataclass.__dataclass_fields__:
        string += str(getattr(dataclass, field)) + ","
    return string[:-1]+"\n"


domains = ["hotmail.com", "gmail.com", "aol.com",
           "mail.com", "mail.kz", "yahoo.com"]


def get_one_random_domain(domains):
    return random.choice(domains)


r = RandomWord()


def get_one_random_name(letters):
    return ''.join(random.choice(letters) for i in range(7))


def email_gen():
    return str(r.word(include_parts_of_speech=["adjectives"])+r.word(include_parts_of_speech=["nouns"])) + '@' + get_one_random_domain(domains)
