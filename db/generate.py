import random
import uuid
from faker import Faker

fake = Faker()

def generate_new_id(arr):                 
    return arr.append(str(uuid.uuid4()))  # Generate a UUID,

def generate_sail_num():
    return random.randint(1, 8888)

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
                
def generate_class():
     return random.randint(1, 11)


def generate_rating_name():
     return fake.word()

def generate_rule_num():
     return random.choice([10, 11, 12, 13, 15, 16, 17, 18, 19, 20, 21, 22, 23, 31, 42])

def generate_comment():
     return fake.word()

def generate_protest_status():
     return random.randint(1, 2)

def generate_crew_status_in_protest():
     return random.randint(1, 3)

def generate_spec_circ():
     return random.randint(0, 13)

def generate_points(i, step, cnt):
    while step > cnt:
        step -= cnt
    points = i + step
    if points > cnt:
        points -= cnt
    return points

