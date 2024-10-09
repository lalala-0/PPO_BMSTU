from pymongo import MongoClient
from bson.objectid import ObjectId
from datetime import datetime

# Подключение к MongoDB
client = MongoClient('mongodb://localhost:27017/')
db = client['ppo_db']  # Замените на имя вашей базы данных

# Создание коллекций и добавление документов

# participants
db.participants.insert_one({
    "_id": str,
    "name": "Participant Name",
    "category": 1,
    "gender": True,
    "birthdate": datetime(2000, 1, 1),
    "coach_name": "Coach Name"
})

# ratings
db.ratings.insert_one({
    "_id": str,
    "name": "Rating Name",
    "class": 1,
    "blowout_cnt": 0
})

# crews
db.crews.insert_one({
    "_id": str,
    "rating_id": str,  # Ссылка на коллекцию ratings
    "class": 1,
    "sail_num": 12345
})

# races
db.races.insert_one({
    "_id": str,
    "rating_id": str,  # Ссылка на коллекцию ratings
    "number": 1,
    "class": 1,
    "date": datetime(2024, 9, 9)
})

# judges
db.judges.insert_one({
    "_id": str,
    "name": "Judge Name",
    "login": "judge_login",
    "password": "hashed_password",
    "role": 1,
    "post": "Post description"
})

# protests
db.protests.insert_one({
    "_id": str,
    "race_id": str,  # Ссылка на коллекцию races
    "rating_id": str,  # Ссылка на коллекцию ratings
    "judge_id": str,  # Ссылка на коллекцию judges
    "rule_num": 123,
    "review_date": datetime(2024, 9, 9),
    "status": 1,
    "comment": "Some comment"
})

# crew_protest
db.crew_protest.insert_one({
    "_id": str,
    "crew_id": str,  # Ссылка на коллекцию crews
    "protest_id": str,  # Ссылка на коллекцию protests
    "crew_status": 1
})

# crew_race
db.crew_race.insert_one({
    "_id": str,
    "crew_id": str,  # Ссылка на коллекцию crews
    "race_id": str,  # Ссылка на коллекцию races
    "points": 10,
    "spec_circumstance": 0
})

# participant_crew
db.participant_crew.insert_one({
    "_id": str,
    "participant_id": str,  # Ссылка на коллекцию participants
    "crew_id": str,  # Ссылка на коллекцию crews
    "helmsman": True,
    "active": True
})

# judge_rating
db.judge_rating.insert_one({
    "_id": str,
    "judge_id": str,  # Ссылка на коллекцию judges
    "rating_id": str  # Ссылка на коллекцию ratings
})

print("Данные успешно добавлены в MongoDB!")
