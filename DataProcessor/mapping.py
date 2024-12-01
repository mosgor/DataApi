
def mapping(df, mapping_data):
    model = {}
    for m in mapping_data:
        path_model = m['model_path'].split('/')
        temp = model
        for p in path_model[:-1]:
            if temp.get(p) is None:
                temp[p] = {}
            temp = temp.get(p)


        path_source = m['source_path'].split('/')
        data = df
        for p in path_source[1::]:
            data = data.get(p)
        temp[path_model[-1]] = data
    print(model)
    return model
