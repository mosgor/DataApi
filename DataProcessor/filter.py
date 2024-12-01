comparison = {
    "less": lambda a, b: a < b,
    "greater": lambda a, b: a > b,
    "equal": lambda a, b: a == b,
    "not equal": lambda a, b: a != b
}

def filter(df, filters_data):

    """
    filtered_df = df
    for fil in filters_data:
        if fil["arg"] in df.columns:
            filtered_df = filtered_df[
                comparison[fil["func"]](filtered_df[fil["field"]], filtered_df[fil["arg"]])
            ]
        else:
            filtered_df = filtered_df[
                comparison[fil["func"]](filtered_df[fil["field"]], fil["arg"])
            ]

    return filtered_df
    """

    for fil in filters_data:
        path1 = fil['field_path'].split('/')
        data1 = df
        for p in path1[1::]:
            data1 = data1.get(p)
        value1 = data1

        if type(fil['arg']) == str:
            path2 = fil['arg'].split('/')
            data2 = df
            for p in path2[1:-1]:
                data2 = data2.get(p)
            value2 = data2
        else:
            value2 = fil['arg']

        if comparison[fil['func']](value1, value2):
            return df
        else:
            return None