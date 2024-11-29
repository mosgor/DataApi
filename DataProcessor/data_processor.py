import pandas as pd

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
