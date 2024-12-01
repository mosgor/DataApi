import json

import pandas as pd
from pymongo import MongoClient

from client import run_client
from mapping import *
from transform import *
from filter import *

def docker(df):
    client = MongoClient('mongodb://localhost:27017/DataApi')
    mapp = client['DataApi']
    collection = mapp['mappings']
    
    source_id = df['source_id']

    query = {'source_id': source_id}
    documents = collection.find(query)

    data = json.loads(df['data_json'].replace("'", '"'))

    for document in documents:
        data = mapping(data, document['mapping'])
        #data = transform(data, document['transformation'])
        #data = filter(data, document['filters'])
        #df['data_json'] = str(data)
        #print(df)
        #run_client(df, document['model_id'])