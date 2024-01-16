import grpc
from concurrent import futures
import time
import query_pb2_grpc as pb2_grpc
import query_pb2 as pb2

class ResponseServicer(pb2_grpc.ResponseServiceServicer):

    def __init__(self, *args, **kwargs):
        pass

    def GetResponse(self, request, context):
        message = request.message
        result = "Received message: {}".format(message)

        return pb2.ResponseBody(
            message=result,
            status=200
        )

def serve():
    print("Bootstrapping server...")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_ResponseServiceServicer_to_server(ResponseServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server listening on port 50051")
    server.wait_for_termination()

#TODO: Implement ML Model to serve requests

if __name__ == "__main__":
    serve()
