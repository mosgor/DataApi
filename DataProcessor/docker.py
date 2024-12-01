import json

from pymongo import MongoClient

from client import run_client
from mapping import *
from transform import *
from filter import *

def docker(df):
    client = MongoClient('mongodb://mongo:27017/DataApi')
    mapp = client['DataApi']
    collection = mapp['mappings']
    
    source_id = df['source_id']

    for s_id in source_id:
        query = {'source_id': s_id}
        documents = collection.find(query)

        data = json.loads(df['data_json'].replace("'", '"'))

        for document in documents:
            data = mapping(data, document['mapping'])
            data = transform(data, document['transformation'])
            data = filter(data, document['filters'])
            df['data_json'] = str(data)
            run_client(df, document['model_id'])