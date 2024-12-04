db = db.getSiblingDB('DataApi');

db.source_schemas.insertOne({
    source_id: 1,
    fields: [
        {path: "testFolder/stringField", type: "string"},
        {path: "intField", type: "int"},
        {path: "dateField", type: "date"}
    ]
});

db.model_schemas.insertOne({
    model_id: 1,
    fields: [
        {path: "fieldString", type: "string"},
        {path: "fieldInt", type: "int"},
        {path: "age", type: "int"},
        {path: "testFolder/subInt", type: "int"}
    ]
});

db.mappings.insertOne({
    source_id: [1],
    model_id: 1,
    mapping: [
        {source_path:"1/testFolder/stringField",  model_path:"fieldString"},
        {source_path:"1/intField",  model_path:"fieldInt"},
        {source_path:"1/dateField", model_path:"age"},
		{source_path:"1/intField", model_path:"testFolder/subInt"}
    ],
    transformation: [
        {field_path: "age", func: "yo"}
    ],
    filters: [
        {field_path: "testFolder/subInt", func: "less", arg: 10}
    ]
});