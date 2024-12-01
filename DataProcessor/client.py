import grpc
from google.protobuf.internal.well_known_types import Timestamp

from proto import model_orchestrator_pb2
from proto import model_orchestrator_pb2_grpc

def client(df, model_id):
    df['arrival_time'].ToDatetime()
    yield model_orchestrator_pb2.ProcessedData(
        source_id=df['source_id'],
        model_id=model_id,
        data_json=df['data_json'],
        arrival_time=df['arrival_time'],
    )


def run_client(df, model_id):
    with grpc.insecure_channel('model_orchestrator:11411') as channel:
        stub = model_orchestrator_pb2_grpc.ModelOrchestratorStub(channel)

        response = stub.SendData(client(df, model_id))

        if response.status_code == 0:
            print(f"Success: {response.message}")
        else:
            print(f"Error: {response.status_code}: {response.message}")


    # channel = grpc.insecure_channel('localhost:11411')
    # stub = model_orchestrator_pb2_grpc.ModelOrchestratorStub(channel)
    # send_stream(stub)

# def send_stream(stub):
#     """Функция для отправки клиентского потока данных."""
#     response = stub.StreamRequest(generate_requests())
#     print(f"Ответ от сервера: {response.summary}")
