import json

import pandas as pd

def mapping(df, mapping_data):
    df.rename(columns=mapping_data, inplace=True)
    return df