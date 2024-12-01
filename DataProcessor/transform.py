from datetime import datetime
from dateutil import parser

def calculate(data):
    return sum(data)


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
    "yo" : yo,
    "calculate" : calculate

}

def transform(df, transformation):
    for trans in transformation:
        path = trans['field_path'].split('/')
        data = df
        for p in path[1:-1]:
            data = data.get(p)
        data[path[-1]] = func[trans['func']](data[path[-1]])
    print(df)
    return df
