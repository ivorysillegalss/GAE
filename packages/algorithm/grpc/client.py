import grpc
from proto import storage_pb2
from proto import storage_pb2_grpc


def run():
    # 连接到 gRPC 服务
    with grpc.insecure_channel('localhost:50061') as channel:
        stub = storage_pb2_grpc.StorageServiceStub(channel)
        # 发送更新请求
        response = stub.UpdateData(storage_pb2.UpdateRequest(data="active"))

    print(f"Response from server: {response.status}")


if __name__ == "__main__":
    run()
