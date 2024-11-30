import json

import pandas as pd

def mapping(df, mapping_data):
    df = pd.DataFrame(df)
    print(df)
    mapping_data = mapping_data['mapping']
    df.rename(columns=mapping_data, inplace=True)
    print(df)