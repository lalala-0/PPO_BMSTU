from pymongo import MongoClient
from pymongo.errors import PyMongoError
import threading

def watch_changes_for_collection(source_collection, target_collection, source_name, target_name):
    """Настраивает Change Stream для одной коллекции."""
    try:
        change_stream = source_collection.watch()

        for change in change_stream:
            # Проверяем источник, чтобы избежать циклической синхронизации
            if 'fullDocument' in change and change['fullDocument'].get('source') != source_name:
                print(f"Изменение в коллекции {source_collection.name} ({source_name}): {change}")
                if change['operationType'] == 'insert':
                    target_collection.update_one(
                        {'_id': change['fullDocument']['_id']},
                        {'$set': {**change['fullDocument'], 'source': source_name}},
                        upsert=True
                    )
                elif change['operationType'] == 'update':
                    target_collection.update_one(
                        {'_id': change['documentKey']['_id']},
                        {'$set': {**change['updateDescription']['updatedFields'], 'source': source_name}}
                    )
                elif change['operationType'] == 'delete':
                    target_collection.delete_one({'_id': change['documentKey']['_id']})
                print(f"Изменение передано из {source_name} в {target_name} (коллекция {source_collection.name})")
    except PyMongoError as e:
        print(f"Ошибка при синхронизации коллекции {source_collection.name} между {source_name} и {target_name}: {e}")

def watch_changes(source_client, target_client, source_name, target_name):
    """Настраивает Change Streams для всех коллекций базы данных."""
    try:
        source_db = source_client['ppo']
        target_db = target_client['ppo']

        # Получаем список всех коллекций в исходной базе
        collections = source_db.list_collection_names()

        threads = []

        for collection_name in collections:
            source_collection = source_db[collection_name]
            target_collection = target_db[collection_name]

            # Создаем поток для каждой коллекции
            thread = threading.Thread(
                target=watch_changes_for_collection,
                args=(source_collection, target_collection, source_name, target_name)
            )
            threads.append(thread)
            thread.start()

        for thread in threads:
            thread.join()

    except PyMongoError as e:
        print(f"Ошибка при настройке синхронизации для {source_name} и {target_name}: {e}")

def main():
    uri1 = "mongodb://localhost:27021"  # Подключение к первому репликационному набору
    uri2 = "mongodb://localhost:27023"  # Подключение ко второму репликационному набору

    client1 = MongoClient(uri1)
    client2 = MongoClient(uri2)

    try:
        # Запускаем два потока для синхронизации всех коллекций между наборами
        thread1 = threading.Thread(target=watch_changes, args=(client1, client2, 'rs1', 'rs2'))
        thread2 = threading.Thread(target=watch_changes, args=(client2, client1, 'rs2', 'rs1'))

        thread1.start()
        thread2.start()

        thread1.join()
        thread2.join()
    except Exception as e:
        print(f"Ошибка при настройке синхронизации: {e}")

if __name__ == '__main__':
    main()
