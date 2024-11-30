import ast
import json
import time
import datetime

import pandas as pd

def yo(time_data):
    formatted_str = time_data.replace("{", '{"').replace(":", '":').replace(", ", ', "').replace("}", '}')
    time_data = ast.literal_eval(formatted_str)

    seconds = time_data["seconds"]

    date = datetime.datetime.utcfromtimestamp(seconds)
    now = datetime.datetime.now()
    year = now.year - date.year

    return year

func = {
    "yo" : yo
}

def transform(df, transformation):
    for key, value in transformation.items():
        df[key] = df[key].map(func[value])
    return df