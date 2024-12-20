import concurrent.futures as futures
import grpc

from data_processor import *

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    data_processor_pb2_grpc.add_DataProcessorServicer_to_server(DataProcessor(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server stared on port 50051")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
