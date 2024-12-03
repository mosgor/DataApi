# ModelOrchestrator

Сервис ответственный за взаимодействие с моделями данных, составление маппингов и сбор статистики.  
Он предоставляет API для фронтенда, позволяющее добавлять, изменять, получать и удалять модели данных и их шаблоны.   
Также предоставляется API для добавления, изменения, получения и редактирования маппингов (документов указывающих на отношение ожидаемых и получаемых полей).  
Отдельным запросом можно получить статистику.

## Маршруты API

Оба API имеют одинаковые методы, но разные пути. Текущий адрес  `http://model_orchestrator:8082`.  
### Пути для моделей 
+ **`GET`** на `/model` - возвращает все модели и их шаблоны сохранённые в базах данных. Обрамляем их в список `Models`.
+ **`GET`** на `/model/{model_id}` - возвращает модель её шаблон по id.  

Схема моделей: 
```json
{
    "id": 1,
    "name": "test_model",
    "connection_string": "http://localhost:8112",
    "created_at": "2024-12-01T10:04:31.705302Z",
    "fields": [
        {
            "path": "testFolder/field_string",
            "type": "string"
        },
        {
            "path": "field_int",
            "type": "int"
        },
        {
            "path": "age",
            "type": "int"
        },
        {
            "path": "sub_int",
            "type": "int"
        }
    ],
    "model_id": 1
}
```
+ **`POST`** на `/model` - добавляет новую модель данных в базы данных и возвращает новую запись.

Схема запроса:
```json
{
  "name": "test_model",
  "connection_string": "http://localhost:8112",
  "fields": [
    {
      "path": "field_string",
      "type": "string"
    },
    {
      "path": "field_int",
      "type": "int"
    },
    {
      "path": "age",
      "type": "int"
    },
    {
      "path": "sub_int/somePath",
      "type": "int"
    }
  ]
}
```

**В процессе разработки:**
+ **`PUT`** на `/model/{model_id}` - обновляет модель по её id и возвращает изменённую модель.
+ **`DELETE`** на `/model/{model_id}` - удаляет модель по её id и возвращает её.

### Пути для маппингов 

+ **`GET`** на `/mapping` - возвращает все маппинги. Обрамляет их в список `Mappings`.
+ **`GET`** на `/mapping/{mapping_id}` - возвращает маппинг по id.

Схема маппинга:
```json
{
  "_id": "674c35d51b17a04188cedc18",
  "source_id": [
    1,
    2
  ],
  "model_id": 1,
  "mapping": [
    {
      "source_path": "1/name",
      "model_path": "nm"
    },
    {
      "source_path": "1/dt",
      "model_path": "temp/database"
    },
    {
      "source_path": "1/field/string",
      "model_path": "str"
    }
  ],
  "transformation": [
    {
      "field_path": "temp/database",
      "func": "calculate",
      "msg": ""
    }
  ],
  "filters": [
    {
      "field_path": "temp/database",
      "func": "less",
      "arg": 10
    }
  ]
}
```
+ **`POST`** на `/mapping` - добавляет новый маппинг в Mongo и возвращает новую запись.

Схема запроса:
```json
{
  "source_id": [2, 3],
  "model_id": 1,
  "mapping": [
    {"source_path":"1/name",  "model_path":"nm"},
    {"source_path":"1/dt",  "model_path":"temp/database"},
    {"source_path":"1/test/string", "model_path":"str"}
  ],
  "transformation": [
    {"field_path": "temp/database", "func": "calculate"}
  ],
  "filters": [
    {"field_path": "temp/database", "func": "less", "arg": 10}
  ]
}
```

**В процессе разработки:**
+ **`PUT`** на `/mapping/{mapping_id}` - обновляет маппинг по id и возвращает изменённую запись.
+ **`DELETE`** на `/mapping/{mapping_id}` - удаляет маппинг по его id и возвращает удалённую запись.

## Структура проекта

Проект имеет следующую структуру:  

+ В директории `ModelOrchestrator/cmd` хранится пакет main и исполняемые файлы.

+ Директория `ModelOrchestrator/config` содержит конфигурации для запуска приложения с разными настройками.   
При запуске, конфиг содержащийся в этой директории, должен быть **обязательно** указан в переменной среды `CONFIG_PATH`.

+ Директория `ModelOrchestrator/pkg` содержит весь go код, не находящийся в пакете main.
+ + В директории `ModelOrchestrator/pkg/config` содержится код ответственный за подгрузку конфигурационных файлов и их парсинг.
+ + В директории `ModelOrchestrator/pkg/gRPC` содержится код отвечающий за получение данных от DataProcessor, отправку их на модель и сбор логов.
+ + В директории `ModelOrchestrator/pkg/logs` содержится код отвечающий за работу с логами, в файле `handlers.go` описан http хэндлер, а в `storage` методы для работы с базой данных.
+ + В директории `ModelOrchestrator/pkg/mapping` содержится код отвечающий за работу с маппингами, в файле `handlers.go` описаны http хэндлеры, а в `storage` методы для работы с базой данных.
+ + В директории `ModelOrchestrator/pkg/model` содержится код отвечающий за работу с моделями, в файле `handlers.go` описаны http хэндлеры, а в `storage` методы для работы с базами данных.
+ + В директории `ModelOrchestrator/pkg/internal` содержатся служебные пакеты, доступ к которым нужен только из `pkg`.
+ + + В директории `ModelOrchestrator/pkg/internal/proto` содержатся сгенерированные из proto файлы для связи по gRPC.
+ + + В директории `ModelOrchestrator/pkg/internal/repositories` содержатся интерфейсы, определяющие методы для работы с базами данных. 
+ + + В директории `ModelOrchestrator/pkg/internal/storageClients` содержится интерфейс, выделяющий только необходимые методы из библиотеки `pgx` для работы с Postgres.
+ + + В директории `ModelOrchestrator/pkg/internal/structs` описываются структуры используемые в API

## Используемые библиотеки

В проекте были использованы следующие библиотеки:

+ `cleanenv` от `ilyakaznacheev` для парсинга `yaml` файлов.
+ `chi` от `go-shi` для удобного маршрутизирования http запросов и добавления middleware.
+ `render` от `go-chi` для удобной записи структур в тела http ответов. 
+ `pgx` от `jackc` для получения драйверов и инструментов по работе с Postgres.
+ `grpc` от `Google` для работы с системой удалённого вызова процедур `gRPC`.
+ `mongo-dricer` от `mongo` для работы с базой данных MongoDB.
+ Для логгирования использован встроенный пакет `log/slog`.