import csv
import random
from faker import Faker
import generate as g

fake = Faker()

participant_cnt = 1000
judge_cnt = 1000
rating_cnt = 1000
crew_cnt = rating_cnt * 120
race_cnt = rating_cnt * 10
crew_race_cnt = rating_cnt * 1200
protest_cnt = rating_cnt * 100
crew_protest_cnt = rating_cnt * 700
judge_rating_cnt = rating_cnt * 5
participant_crew_cnt = rating_cnt * 360

participant_ids = []
judge_ids = []
rating_ids = []
crew_ids = []
race_ids = []
crew_race_ids = []
protest_ids = []
crew_protest_ids = []
judge_rating_ids = []
participant_crew_ids = []

def generate_ids():
    for _ in range(participant_cnt):
        g.generate_new_id(participant_ids)
    for _ in range(judge_cnt):
        g.generate_new_id(judge_ids)
    for _ in range(rating_cnt):
        g.generate_new_id(rating_ids)
    for _ in range(crew_cnt):
        g.generate_new_id(crew_ids)
    for _ in range(race_cnt):
        g.generate_new_id(race_ids)
    for _ in range(crew_race_cnt):
        g.generate_new_id(crew_race_ids)
    for _ in range(protest_cnt):
        g.generate_new_id(protest_ids)
    for _ in range(crew_protest_cnt):
        g.generate_new_id(crew_protest_ids)
    for _ in range(judge_rating_cnt):
        g.generate_new_id(judge_rating_ids)
    for _ in range(participant_crew_cnt):
        g.generate_new_id(participant_crew_ids)

def Generate_participants():
    with open(".\\data\\participants_data.csv", mode="a", newline="") as file:
        writer = csv.writer(file, delimiter=';')
        for i in range(participant_cnt):
            writer.writerow([
                participant_ids[i],
                fake.name(),
                random.randint(1, 8),
                g.generate_bool(),
                g.generate_datetime(),
                fake.name()
        ])

def Generate_judges():
    with open(".\\data\\judges_data.csv", mode='a', newline="") as file:
        writer = csv.writer(file, delimiter=';')
        for i in range(judge_cnt):
            writer.writerow([
                 judge_ids[i],
                fake.name(),
                fake.user_name(),
                fake.password(),
                random.randint(1, 2),
                random.choice(["Chief Secretary", 'Deputy Chief Judge for the MTO', 'Deputy Chief Distance Judge', 'Chairman of the protest Committee (chief umpire in races with direct judging on the water)', 'Chairman of the Technical Committee', 'Senior judge at the start', 'The senior judge at the finish line', 'Judge of the protest committee; empire in races with direct judging on the water', 'The referee is at the starting line', 'The referee is at the finish line'])
            ])

def generate_crew(crewID, ratingID, ratingClass):
    with open(".\\data\\crews_data.csv", mode="a", newline="") as file:
        sailNum = g.generate_sail_num()
        writer = csv.writer(file, delimiter=';')
        writer.writerow([
            crewID,
            ratingID,
            ratingClass,
            sailNum
        ])
    return [crewID, ratingID, ratingClass, sailNum]


def generate_race(raceID, ratingID, ratingClass, num):
    with open(".\\data\\races_data.csv", mode="a", newline="") as file:
            date = g.generate_datetime()
            writer = csv.writer(file, delimiter=';')
            writer.writerow([
                raceID,
                ratingID,
                num,
                ratingClass,
                date
            ])
    return [raceID, ratingID, ratingClass, date]

def generate_crew_race(crew_receID, crewID, raceID, points):
    with open(".\\data\\crew_race_data.csv", mode="a", newline="") as file:
        writer = csv.writer(file, delimiter=';')
        writer.writerow([
            crew_receID,
            crewID,
            raceID,
            points,
            g.generate_spec_circ()
        ])

def generate_protest(protestID, raceID, ratingID):
    with open(".\\data\\protests_data.csv", mode="a", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow([
                protestID,
                raceID,
                ratingID, 
                generate_judge_id(),
                g.generate_rule_num(),
                g.generate_datetime(),
                g.generate_protest_status(),
                g.generate_comment()
                ])


def generate_crew_protest(i, witneses_cnt, crews, protestID):
    with open(".\\data\\crew_protest_data.csv", mode="a", newline="") as file:
        writer = csv.writer(file, delimiter=';')
        writer.writerow([
                crew_protest_ids[i],
                crews[random.randint(0, len(crews) - 1)][0],
                protestID,
                1
            ])
        i += 1
        writer.writerow([
                crew_protest_ids[i],
                crews[random.randint(0, len(crews) - 1)][0],
                protestID,
                2
            ])
        i += 1
        for _ in range(witneses_cnt):
            writer.writerow([
                crew_protest_ids[i],
                crews[random.randint(0, len(crews) - 1)][0],
                protestID,
                3
            ])
            i += 1
    return i

rating_classes = []
def generate_rating_classes():
    for _ in range(rating_cnt):
         rating_classes.append(g.generate_class())  # Generate a UUID,

def Generate_ratings():
    with open(".\\data\\ratings_data.csv", mode='a', newline="") as file:
        crew_counter = 0
        race_counter = 0
        crew_race_counter = 0
        protest_counter = 0
        crew_protest_counter = 0
        for i in range(rating_cnt):
            writer = csv.writer(file, delimiter=';')
            generate_rating_classes()
            writer.writerow([
                rating_ids[i],
                g.generate_rating_name(),
                rating_classes[i],
                random.randint(1, 7)
            ])

            crews = []
            races = []
            for j in range (random.randint(10, 120)):
                crews.append(generate_crew(crew_ids[crew_counter], rating_ids[i], rating_classes[i]))
                crew_counter += 1
            for j in range(random.randint(2, 10)):
                races.append(generate_race(race_ids[race_counter], rating_ids[i], rating_classes[i], j+1))
                race_counter += 1
            k = 0
            points_list = []
            for race in races:
                numbers = list(range(1, len(crews)))
                random.shuffle(numbers)
                points_list.append(numbers)

            for crew in crews:
                for j in range(len(races)):
                    generate_crew_race(crew_race_ids[crew_race_counter], crew[0], races[j][0], points_list[j][k-1])
                    crew_race_counter += 1
                k += 1
            for race in races:
                for k in range(random.randint(1, 10)):
                    generate_protest(protest_ids[protest_counter], race[0], rating_ids[i])
                    crew_protest_counter = generate_crew_protest(crew_protest_counter, random.randint(1, 5), crews, protest_ids[protest_counter])
                    protest_counter += 1
    global crew_cnt
    crew_cnt = crew_counter
    global race_cnt
    race_cnt = race_counter
    global crew_race_cnt
    crew_race_cnt = crew_race_counter
    global protest_cnt
    protest_cnt = protest_counter
    global crew_protest_cnt
    crew_protest_cnt = crew_protest_counter

def generate_judge_id():
    return random.choice(judge_ids)

def generate_participant_crew(participant_crewID, crewID, helmsman):
        with open(".\\data\\participant_crew_data.csv", mode="a", newline="") as file:
            writer = csv.writer(file, delimiter=';')
            writer.writerow([
                participant_crewID,
                generate_participant_id(),
                crewID,
                helmsman,
                g.generate_bool()
            ])

def Generate_participant_crews():
    participant_crew_counter = 0
    for i in range(crew_cnt):
        helmsman = True
        for j in range(random.randint(1, 3)):
            generate_participant_crew(participant_crew_ids[participant_crew_counter], crew_ids[i], helmsman)
            helmsman = False
            participant_crew_counter += 1
    global participant_crew_cnt
    participant_crew_cnt = participant_crew_counter

def generate_participant_id():
    return random.choice(participant_ids)

def generate_judge_rating(judge_ratingID, ratingID):
    with open(".\\data\\judge_rating_data.csv", mode="a", newline="") as file:
        writer = csv.writer(file, delimiter=';')
        writer.writerow([
            judge_ratingID,
            generate_judge_id(),
            ratingID
        ])

def Generate_judge_ratings():
    judge_rating_counter = 0
    for i in range(rating_cnt):
        for j in range(random.randint(1, 5)):
            generate_judge_rating(judge_rating_ids[judge_rating_counter], rating_ids[i])
            judge_rating_counter += 1
    global judge_rating_cnt 
    judge_rating_cnt = judge_rating_counter

def create_free_files():
    f = open(".\\data\\ratings_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "name", "class", "blowout_cnt"])
    f.close()
    f = open(".\\data\\participants_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "name", "category", "gender", "birthdate", "coach_name"])
    f.close()
    f = open(".\\data\\judges_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "name", "login", "password", "role", "status"])
    f.close()
    f = open(".\\data\\crews_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "rating_id", "class", "sail_num"])
    f.close()
    f = open(".\\data\\races_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "rating_id", "number", "class", "date"])
    f.close()
    f = open(".\\data\\crew_race_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "crew_id", "race_id", "points", "spec_circumstance"])
    f.close()
    f = open(".\\data\\crew_protest_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "crew_id", "protest_id", "crew_status"])
    f.close()
    f = open(".\\data\\participant_crew_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "participant_id", "crew_id", "helmsman", "active"])
    f.close()
    f = open(".\\data\\judge_rating_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "judge_id", "rating_id"])
    f.close()
    f = open(".\\data\\protests_data.csv", mode="w", newline="")
    writer = csv.writer(f, delimiter=';')
    writer.writerow(["id", "race_id", "rating_id", "judge_id", "rule_num", "review_date", "status", "comment"])
    f.close()


if __name__ == '__main__':
    create_free_files()
    generate_ids()

    Generate_judges()
    Generate_participants()
    Generate_ratings()
    Generate_participant_crews()
    Generate_judge_ratings()