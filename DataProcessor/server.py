import grpc
from concurrent import futures

from proto import model_orchestrator_pb2_grpc
from proto import common_pb2


class ModelOrchestratorServicer(model_orchestrator_pb2_grpc.ModelOrchestratorServicer):
    def SendData(self, request_iterator, context):
        print("Send Messages:")
        for request in request_iterator:
            print(f"source_id={request.source_id}, model_id={request.model_id}, "
                  f"data_json={request.data_json}, arrival_time={request.arrival_time}")

        return common_pb2.Status(
            message="All data has been successfully received"
        )


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    model_orchestrator_pb2_grpc.add_ModelOrchestratorServicer_to_server(ModelOrchestratorServicer(), server)
    server.add_insecure_port('[::]:11411')
    server.start()
    print("Server stared on port 11411")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
