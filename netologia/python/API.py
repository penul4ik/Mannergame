import requests
import json
from iso3166 import countries
from urllib import parse # поможет нам безопасно заменить символы из url на %[число], чтобы браузер смог понять что является запросом, а что URLом

# print(parse.quote('https://cnd.mos.cms.futurecdn.net/snbrHBRigvvzjxNGuUtcck.jpg?a=4&parse=true'))

data = requests.get('https://jsonplaceholder.typicode.com/posts').json()

for i in range(5):
    print(i,f"Title: {data[i]['title']}")
    print( f"Body: {data[i]['body']}")
    print()
    
country = input("Country code [Ex: RU]: ") 
city = input("City name [Ex: Moscow]: ")
APIkey = input("OpenWeather API key: ")

# Получаем код страны по ISO3166
countryCode = str(countries.get('RU')).split(', ')[3].split('=')[1].replace("'","")

# Получаем долготу и широту
r = requests.get(f'http://api.openweathermap.org/geo/1.0/direct?q={city},{countryCode}&limit=5&appid={APIkey}')

# Делаем запрос к API используя долготу и широту из прошлого запроса
prognoz = requests.get(f'https://api.openweathermap.org/data/2.5/weather?lat={r.json()[0]['lat']}&lon={r.json()[0]['lon']}&appid={APIkey}&units=metric')

print(f"decription: {prognoz.json()['weather'][0]['description']}")
print(f"temp: {prognoz.json()['main']['temp']}")
print(f"wind: {prognoz.json()['wind']['speed']}")
