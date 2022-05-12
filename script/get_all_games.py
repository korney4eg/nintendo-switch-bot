#!/usr/bin/env python3
import requests
import json

ROWS_PER_PAGE = 24
all_games = []
all_games_num = 0

for i in range(0, 351):
    start_rows=i * ROWS_PER_PAGE
    print('Processing page: '+str(i))
    url = 'https://searching.nintendo-europe.com/ru/select?' +\
      'q=*&fq=type%3AGAME%20AND%20((playable_on_txt%3A%22HAC%22))'+\
      '%20AND%20sorting_title%3A*%20AND%20*%3A*&sort=deprioritise_b%20asc'+\
      '%2C%20popularity%20asc&start='+str(start_rows)+'&rows=24&wt=json'+\
      '&bf=linear(ms(priority%2CNOW%2FHOUR)%2C3.19e-11%2C0)'+\
      '&bq=!deprioritise_b%3Atrue%5E999'

    r = requests.get(url)
    if r.status_code != 200:
        print('Nothing found')
        print (r.content)
        break
    all_games += json.loads(r.content)['response']['docs']

jsonString = json.dumps(all_games)
jsonFile = open("all_games.json", "w")
jsonFile.write(jsonString)
jsonFile.close()
