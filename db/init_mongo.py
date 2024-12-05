from pymongo import MongoClient

# Подключение к MongoDB
client = MongoClient('mongodb://localhost:27023/')
db = client["ppo"]

# Создание коллекций
db.create_collection("participants")
db.create_collection("ratings")
db.create_collection("crews")
db.create_collection("races")
db.create_collection("judges")
db.create_collection("protests")
db.create_collection("crew_protest")
db.create_collection("crew_race")
db.create_collection("participant_crew")
db.create_collection("judge_rating")

print("Данные успешно добавлены в MongoDB!")
