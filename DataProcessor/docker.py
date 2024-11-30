import json

from pymongo import MongoClient

from mapping import *
from transform import *
from filter import *

def docker(df):
    client = MongoClient('mongodb://localhost:27017/DataApi')
    mapp = client['DataApi']
    collection = mapp['mappings']
    
    source_id = df['source_id']
    find = {'source_id': source_id}

    documents = collection.find()

    data = json.loads(df['data_json'].replace("'", '"'))
    data = pd.DataFrame([data])

    for document in documents:
        data = mapping(data, document['mapping'])
        data = transform(data, document['transformation'])
        data = filter(data, document['filters'])
        print(data)