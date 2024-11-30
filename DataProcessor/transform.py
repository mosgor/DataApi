from datetime import datetime
from dateutil import parser

def yo(time_data):
    birthdate = parser.parse(time_data)
    current_date = datetime.now()
    age = current_date.year - birthdate.year

    if (current_date.month, current_date.day) < (birthdate.month, birthdate.day):
        age -= 1
    if age < 0:
        return
    return age

func = {
    "yo" : yo
}

def transform(df, transformation):
    for key, value in transformation.items():
        df[key] = df[key].map(func[value])
    return df