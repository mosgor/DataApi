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

