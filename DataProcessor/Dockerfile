FROM python:latest

WORKDIR /DataProcessor

COPY . .

RUN pip install -r requirements.txt

EXPOSE 50051

ENTRYPOINT python main.py