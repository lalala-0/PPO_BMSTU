import csv
import random
import uuid
import hashlib
from faker import Faker
from datetime import datetime, timedelta

fake = Faker()

# Number of rows to generate tables
ratings_rows = 2000
crews_rows = 2000
judges_rows = 2000
protests_rows = 2000
participants_rows = 2000
races_rows = 2000

crew_rece_rows = 2000
crew_protest_rows = 2000
judge_rating_rows = 2000
participant_crew_rows = 2000

ratings_id = []
crews_id = []
judges_id = []
protests_id = []
participants_id = []
races_id = []

crew_rece_id = []
crew_protest_id = []
judge_rating_id = []
participant_crew_id = []


def generate_new_id(arr):                 
    return arr.append(str(uuid.uuid4()))  # Generate a UUID,

def generate_ids():
     for _ in range(judges_rows):
          generate_new_id(judges_id)
          generate_new_id(ratings_id)
          generate_new_id(crews_id)
          generate_new_id(protests_id)
          generate_new_id(participants_id)
          generate_new_id(races_id)
          generate_new_id(crew_rece_id)
          generate_new_id(crew_protest_id)
          generate_new_id(judge_rating_id)
          generate_new_id(participant_crew_id)


# Function to generate a random full name
def generate_full_name():
    return fake.name()


# Function to generate a random full name
def generate_full_name_f():
    return fake.name_female()


# Function to generate a random international flag (boolean)
def generate_bool():
    return random.choice([True, False])


# Function to generate a random age
def generate_age():
    return random.randint(6, 18)


# Function to generate a random created date
def generate_datetime():
    year = random.randint(2000, 2023)
    month = random.randint(1, 12)
    day = random.randint(1, 28)
    hour = random.randint(8, 19)
    minuts = random.randint(1, 6)
    return f"{year}-{month:02d}-{day:02d} {hour:02d}:{minuts:02d}:23"


def generate_participant_id():
    return random.choice(participants_id)

def generate_crew_id():
    return random.choice(crews_id)

def generate_protest_id():
    return random.choice(protests_id)

def generate_race_id():
    return random.choice(races_id)

def generate_rating_id():
    return random.choice(ratings_id)

def generate_judge_id():
    return random.choice(judges_id)

def generate_judges():
    with open("E:\\PPO_BMSTU\\db\\data\\judges_data.csv", mode='w', newline="") as file:
        writer = csv.writer(file, delimiter=';')
        writer.writerow(["id", "name", "login", "password", "role", "status"])
        for i in range(judges_rows):
            writer.writerow([
                 judges_id[i],
                fake.name(),
                fake.user_name(),
                fake.password(),
                random.randint(1, 2),
                random.choice(["Chief Secretary", 'Deputy Chief Judge for the MTO', 'Deputy Chief Distance Judge', 'Chairman of the protest Committee (chief umpire in races with direct judging on the water)', 'Chairman of the Technical Committee', 'Senior judge at the start', 'The senior judge at the finish line', 'Judge of the protest committee; empire in races with direct judging on the water', 'The referee is at the starting line', 'The referee is at the finish line'])
            ])


def generate_participants():
    with open("E:\\PPO_BMSTU\\db\\data\\participants_data.csv", mode="w", newline="") as file:
        writer = csv.writer(file, delimiter=';')
        writer.writerow(["id", "name", "category", "gender", "birthdate", "coach_name"])
        for i in range(participants_rows):
            writer.writerow([
                participants_id[i],
                fake.name(),
                random.randint(1, 8),
                generate_bool(),
                generate_datetime(),
                fake.name()
        ])
                
def generate_class():
     return random.randint(1, 11)

def generate_crews():
    with open("E:\\PPO_BMSTU\\db\\data\\crews_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "rating_id", "class", "sail_num"])
            for i in range(crews_rows):
                writer.writerow([
                     crews_id[i],
                    generate_rating_id(),
                    generate_class(),
                    i
                ])


def generate_rating_name():
     return fake.word()


def generate_ratings():
    with open("E:\\PPO_BMSTU\\db\\data\\ratings_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "name", "class", "blowout_cnt"])
            for i in range(ratings_rows):
                writer.writerow([
                     ratings_id[i],
                    generate_rating_name(),
                    generate_class(),
                    random.randint(1, 7)
                ])

def generate_rule_num():
     return random.choice([10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23, 31, 42])

def generate_comment():
     return fake.word()

def generate_protest_status():
     return random.randint(1, 2)

def generate_protests():
    with open("E:\\PPO_BMSTU\\db\\data\\protests_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "race_id", "rating_id", "judge_id", "rule_num", "review_date", "status", "comment"])
            for i in range(protests_rows):
                writer.writerow([
                    protests_id[i],
                    generate_race_id(),
                    generate_rating_id(), 
                    generate_judge_id(),
                    generate_rule_num(),
                    generate_datetime(),
                    generate_protest_status(),
                    generate_comment()
                    ])


def generate_races():
    with open("E:\\PPO_BMSTU\\db\\data\\races_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "rating_id", "class", "date"])
            for i in range(protests_rows):
                writer.writerow([
                     races_id[i],
                    generate_rating_id(),
                    generate_class(),
                    generate_datetime()
                ])


######################
def generate_crew_status_in_protest():
     return random.randint(1, 3)

def generate_crew_protest():
    with open("E:\\PPO_BMSTU\\db\\data\\crew_protest_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "crew_id", "protest_id", "crew_status"])
            for i in range(protests_rows):
                writer.writerow([
                    crew_protest_id[i],
                    generate_crew_id(),
                    generate_protest_id(),
                    generate_crew_status_in_protest()
                ])

def generate_spec_circ():
     return random.randint(0, 13)

def generate_points():
     return random.randint(1, 101)

def generate_crew_race():
    with open("E:\\PPO_BMSTU\\db\\data\\crew_race_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "crew_id", "race_id", "points", "spec_circumstance"])
            for i in range(protests_rows):
                writer.writerow([
                    crew_rece_id[i],
                    generate_crew_id(),
                    generate_race_id(),
                    generate_points(),
                    generate_spec_circ()
                ])


def generate_participant_crew():
    with open("E:\\PPO_BMSTU\\db\\data\\participant_crew_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "participant_id", "crew_id", "helmsman", "active"])
            for i in range(protests_rows):
                writer.writerow([
                     participant_crew_id[i],
                    generate_participant_id(),
                    generate_crew_id(),
                    generate_bool(),
                    generate_bool()
                ])

def generate_judge_rating():
    with open("E:\\PPO_BMSTU\\db\\data\\judge_rating_data.csv", mode="w", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow(["id", "judge_id", "rating_id"])
            for i in range(protests_rows):
                writer.writerow([
                     judge_rating_id[i],
                    generate_judge_id(),
                    generate_rating_id()
                ])

if __name__ == '__main__':
    generate_ids()


    generate_judges()
    generate_participants()
    generate_crews()
    generate_ratings()
    generate_protests()
    generate_races()
    generate_crew_protest()
    generate_crew_race()
    generate_judge_rating()
    generate_participant_crew()
