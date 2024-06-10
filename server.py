import grpc
from concurrent import futures
from model.greet.greet_pb2_grpc import GreetServicer,add_GreetServicer_to_server
from model.greet import greet_pb2

class HelloService(GreetServicer):
    def SayHello(self, request, context):
        say_hello = greet_pb2.Result()
        say_hello.message = f"hello {request.name}"
        return say_hello
       

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_GreetServicer_to_server(HelloService(), server)
    server.add_insecure_port('localhost:5001')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()