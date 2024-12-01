import json

import pandas as pd

def normilize_data(df):
    data = json.loads(df['data_json'].replace("'", '"'))
    data = pd.json_normalize(data, sep='/')
    return data