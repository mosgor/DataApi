import grpc
from google.protobuf.internal.well_known_types import Timestamp

from proto import model_orchestrator_pb2
from proto import model_orchestrator_pb2_grpc

def client(df, model_id):
    ts = Timestamp()
    ts.ToDatetime(df['arrival_time'])
    yield model_orchestrator_pb2.ProcessedData(
        source_id=int(df['source_id']),
        model_id=int(model_id),
        data_json=str(df['data_json']),
        arrival_time=ts,
    )


def run_client(df, model_id):
    with grpc.insecure_channel('localhost:11411') as channel:
        stub = model_orchestrator_pb2_grpc.ModelOrchestratorStub(channel)

        response = stub.SendData(client(df, model_id))

        if response.status_code == 0:
            print(f"Success: {response.message}")
        else:
            print(f"Error: {response.status_code}: {response.message}")
