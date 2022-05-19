#!/usr/bin/env python3
import sqlite3
import json

# Connecting to sqlite
conn = sqlite3.connect('data.db')


def get_num_from_bool(b):
    if b:
        return 1
    else:
        return 0


# Creating a cursor object using the
# cursor() method
cursor = conn.cursor()
categories = []
languages = []
limit = 1000000
game_num = 0
with open('all_games_f.json') as json_file:
    json_data = json.load(json_file)
    for game in json_data:
        if game_num > limit:
            break
        game_num += 1
        print(f"{game['fs_id']}: {game['title']}")
        if 'price_regular_f' not in game.keys():
            print('No regular price. Skipped')
            continue

        if 'paid_subscription_required_b' in game.keys():
            paid_subscription_required_b = game['paid_subscription_required_b']
        else:
            paid_subscription_required_b = False
        title = game['sorting_title'].replace('"', "'")
        if 'demo_availability' in game.keys():
            demo_availability = game['demo_availability']
        else:
            demo_availability = False
        if 'publisher' in game.keys():
            publisher = game['publisher']
        else:
            publisher = ''
        game_query = f'''INSERT INTO games VALUES (\
            {game['fs_id']},\
            "{game['change_date'][0:10]}",\
            "{game['url']}",\
            "{game.get('image_url','')}",\
            {get_num_from_bool(paid_subscription_required_b)},\
            {get_num_from_bool(game['cloud_saves_b'])},\
            {game.get('priority',"2099-09-09")[0:10]},\
            {get_num_from_bool(demo_availability)},\
            {game['age_rating_sorting_i']},\
            "{publisher}",\
            "{game['excerpt'].replace('"', "'")}",\
            "{game.get('date_from',"2017-09-09")[0:10]}",\
            {game['price_discount_percentage_f']},\
            "{title}",\
            "",\
            {game['players_to']},\
            {game['price_regular_f']},\
            {game['price_lowest_f']},\
            {game['_version_']})'''
        # print(game['fs_id'])
        # print(game_query)
        cursor.execute(game_query)
        conn.commit()
        if 'game_categories_txt' in game.keys():
            game_categories = game['game_categories_txt']
        elif 'game_category' in game.keys():
            game_categories = game['game_category'].split(",")
        else:
            game_categories = []

        i = 0
        for category in game_categories:
            categories.append(
                (category, game['pretty_game_categories_txt'], game['fs_id'])) if (category, game['pretty_game_categories_txt'][i], game['fs_id']) not in categories else categories
        if game['language_availability'][0] != '':
            langs = game['language_availability'][0].split(',')
        else:
            langs = ['english']
        # print(langs)
        for lang in langs:
            languages.append(
                (lang, game['fs_id'])) if (lang, game['fs_id']) not in languages else languages


# Queries to INSERT records.
for category in categories:
    cursor.execute(
        f'''INSERT INTO categories(name, russian_name, game_id) VALUES('{category[0]}', NULL, '{category[2]}')''')

for lang in languages:
    cursor.execute(
        f'''INSERT INTO languages(lang_name, game_id) VALUES('{lang[0]}', '{lang[1]}')''')

# Commit your changes in the database
conn.commit()

# Closing the connection
conn.close()
