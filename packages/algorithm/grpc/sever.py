import grpc

from concurrent import futures
from proto import relation_net_pb2
from proto import relation_net_pb2_grpc
from proto import nation_pb2
from proto import nation_pb2_grpc
from proto import storage_pb2
from proto import storage_pb2_grpc
from query import Query


class RelationService(relation_net_pb2_grpc.RelationNetServiceServicer):

    def QueryRelationNet(self, request, context):
        rank = Query()
        user = rank.get_relation_for_user(request.username)
        return relation_net_pb2.RelationResponse(user)


class NationService(nation_pb2_grpc.NationServiceServicer):
    def QueryNation(self, request, context):
        rank = Query()
        user = rank.get_nation_for_user(request.username)
        return nation_pb2.RelationResponse(user)


class StorageServiceServicer(storage_pb2_grpc.StorageServiceServicer):
    def UpdateData(self, request, context):
        rank = Query()
        rank.query_Ranks()
        return storage_pb2.StorageResponse("success")


# 启动 gRPC 服务器
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    server.add_insecure_port('[::]:50051')

    nation_pb2_grpc.add_NationServiceServicer_to_server(NationService(), server)
    relation_net_pb2_grpc.add_RelationNetServiceServicer_to_server(RelationService, server)

    server.start()
    print("gRPC server is running on port 50051...")
    server.wait_for_termination()


if __name__ == '__main__':
    serve()
