import csv
import uuid
from pymongo import MongoClient, UpdateOne
from bson.binary import Binary, UUID_SUBTYPE
from datetime import datetime

client = MongoClient('mongodb://localhost:27017/')
db = client['ppo_db']

def convert_uuid(value):
    try:
        return Binary(uuid.UUID(value).bytes, UUID_SUBTYPE)
    except ValueError:
        return value

def convert_date(value):
    try:
        return datetime.strptime(value, '%Y-%m-%d %H:%M:%S') if value else None
    except ValueError:
        return value

def convert_type(value, field_type):
    if field_type == 'uuid':
        return convert_uuid(value)
    elif field_type == 'int':
        try:
            return int(value)
        except ValueError:
            return value
    elif field_type == 'float':
        try:
            return float(value)
        except ValueError:
            return value
    elif field_type == 'bool':
        return value.lower() in ['true', '1', 'yes']
    elif field_type == 'datetime':
        return convert_date(value)
    else:
        return value

type_mapping = {
    'participants': {
        '_id': 'uuid',
        'category': 'int',
        'gender': 'bool',
        'birthdate': 'datetime'
    },
    'ratings': {
        '_id': 'uuid',
        'class': 'int',
        'blowout_cnt': 'int'
    },
    # Добавить другие коллекции по аналогии
}

def clear_collection(collection_name):
    collection = db[collection_name]
    collection.delete_many({})

def load_data_from_csv(file_path, collection_name):
    collection = db[collection_name]
    
    with open(file_path, mode='r', encoding='utf-8') as file:
        reader = csv.DictReader(file, delimiter=';')
        operations = []
        
        for row in reader:
            for field, field_type in type_mapping.get(collection_name, {}).items():
                if field in row and row[field]:
                    row[field] = convert_type(row[field], field_type)
            
            # Создаем операцию upsert
            operations.append(
                UpdateOne({'_id': row['_id']}, {'$set': row}, upsert=True)
            )
        
        if operations:
            result = collection.bulk_write(operations)
            print(f"Вставлено {result.upserted_count} новых записей.")
            print(f"Обновлено {result.modified_count} существующих записей.")
            print(f"Ошибки записи: {result.write_errors}")

data_dir = 'E:/PPO_BMSTU/db/data/'

# Очистка коллекций перед загрузкой данных
clear_collection('participants')
clear_collection('ratings')

# Загрузка данных
load_data_from_csv(data_dir + 'participants_data.csv', 'participants')
load_data_from_csv(data_dir + 'ratings_data.csv', 'ratings')

print("Данные успешно загружены в MongoDB с UUID в формате Binary!")
