import concurrent.futures as futures
import grpc
import pandas as pd

from proto import data_processor_pb2
from proto import data_processor_pb2_grpc
from proto import common_pb2


class DataProcessor(data_processor_pb2_grpc.DataProcessorServicer):
    def ProcessData(self, request, context):
        data_list = []

        for req in request:
            data_dict = {
                "source_id": req.source_id,
                "data_json": req.data_json,
                "arrival_time": req.arrival_time
            }

            data_list.append(data_dict)

        df = pd.DataFrame(data_list)
        print(df)
        return common_pb2.Status(message="OK")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    data_processor_pb2_grpc.add_DataProcessorServicer_to_server(DataProcessor(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server stared on port 50051")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
