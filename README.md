# Data API

Этот проект представляет с собой MVP инструмента для поставления данных из любых источников на production модель машинного обучения. Предназначен для аналитиков данных, которым нужно no-code решение по быстрой настроке передачи данных из источников до модели.

## Архитектура сервиса

Сервис состоит из 6 элементов: двух баз данных (PostgreSQL и  MongoDB), трех сервисов отвечающих за внутреннюю логику и фронтенда.

### Source Manager

Первый сервис, задействованный в цепочке обработки данных. Предоставляет пользователю инструменты по добавлению, редактированию, получению и удалению источников данных и их шаблонов.  
Сервис по запросу или по расписанию запрашивает данные с источников, преоразует их в читаемый для следующего сервиса формат и отправляет посредством *gRPC*.

### Data Processor

Является вторым в цепочке сервисом. На этом этапе данные, полученные с Source Manager преобразуются (включая трансформации данных) в читаемый для production-модели формат, фильтруются и посредством того же *gRPC* отправляются на следующий сервис.

### Model Orchestrator

Третий и заключительный в цепочке сервис. С помощью него пользователь может редактировать, получать, добавлять и удалять production модели, шаблоны желаемых ими данных, маппинги для соотношения полей модели и источников.  
Также даёт возможность посмотреть статистику выполненых операций.

### PostgreSQL

Первая база данных, используется для хранения информации строгого формата, наприме логи, имена моделей и источников и ссылки на них, даты добавления.

### MongoDB

Вторая база данных, используется для хранения информации более свободного формата с множеством вложенных полей, например шаблоны ожидаемой моделью информации, шаблоны получаемой из источников информации и маппинги полей.

### Frontend (предстоит реализовать)

Визуальный интерфейс, позволяющий аналитикам данных и другим потенциальным пользователям легко управлять источниками данных и моделами и выстраивать связи между ними.

## Запуск

Чтобы запустить программу, необходимо скачать исходный код проекта, установить Docker и Docker Compose, в терминале перейти внутрь директории с исходным кодом проекта (если увидете `docker-compose.yml` файл - вы там, где нужно) и прописать команду:
```shell
docker-compose up -d --build
```

## Над проектом работали 
  
1) Михайлов Максим - teamlead, golang backend, архитектор, ответственный за ModelOrchestrator.
2) Кучин Илья - python backend, ответственный за DataProcessor.
3) Стенин Кирилл - golang backend, ответственный за SourceManager.
