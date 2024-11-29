import pandas as pd

def mapping(df, mapping_data):
    df = pd.DataFrame(df)
    print(df)
    df = df.rename(columns=mapping_data['mapping'], inplace=True)
    print(df)