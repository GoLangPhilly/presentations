package org.golangphilly.grpc.file.transfer;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.2.0)",
    comments = "Source: file/server.proto")
public final class FileTransferServiceGrpc {

  private FileTransferServiceGrpc() {}

  public static final String SERVICE_NAME = "FileTransferService";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<org.golangphilly.grpc.file.transfer.Server.FileRequest,
      org.golangphilly.grpc.file.transfer.Server.FileResponse> METHOD_UPLOAD =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.CLIENT_STREAMING,
          generateFullMethodName(
              "FileTransferService", "Upload"),
          io.grpc.protobuf.ProtoUtils.marshaller(org.golangphilly.grpc.file.transfer.Server.FileRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(org.golangphilly.grpc.file.transfer.Server.FileResponse.getDefaultInstance()));

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FileTransferServiceStub newStub(io.grpc.Channel channel) {
    return new FileTransferServiceStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FileTransferServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new FileTransferServiceBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary and streaming output calls on the service
   */
  public static FileTransferServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new FileTransferServiceFutureStub(channel);
  }

  /**
   */
  public static abstract class FileTransferServiceImplBase implements io.grpc.BindableService {

    /**
     */
    public io.grpc.stub.StreamObserver<org.golangphilly.grpc.file.transfer.Server.FileRequest> upload(
        io.grpc.stub.StreamObserver<org.golangphilly.grpc.file.transfer.Server.FileResponse> responseObserver) {
      return asyncUnimplementedStreamingCall(METHOD_UPLOAD, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_UPLOAD,
            asyncClientStreamingCall(
              new MethodHandlers<
                org.golangphilly.grpc.file.transfer.Server.FileRequest,
                org.golangphilly.grpc.file.transfer.Server.FileResponse>(
                  this, METHODID_UPLOAD)))
          .build();
    }
  }

  /**
   */
  public static final class FileTransferServiceStub extends io.grpc.stub.AbstractStub<FileTransferServiceStub> {
    private FileTransferServiceStub(io.grpc.Channel channel) {
      super(channel);
    }

    private FileTransferServiceStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileTransferServiceStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new FileTransferServiceStub(channel, callOptions);
    }

    /**
     */
    public io.grpc.stub.StreamObserver<org.golangphilly.grpc.file.transfer.Server.FileRequest> upload(
        io.grpc.stub.StreamObserver<org.golangphilly.grpc.file.transfer.Server.FileResponse> responseObserver) {
      return asyncClientStreamingCall(
          getChannel().newCall(METHOD_UPLOAD, getCallOptions()), responseObserver);
    }
  }

  /**
   */
  public static final class FileTransferServiceBlockingStub extends io.grpc.stub.AbstractStub<FileTransferServiceBlockingStub> {
    private FileTransferServiceBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private FileTransferServiceBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileTransferServiceBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new FileTransferServiceBlockingStub(channel, callOptions);
    }
  }

  /**
   */
  public static final class FileTransferServiceFutureStub extends io.grpc.stub.AbstractStub<FileTransferServiceFutureStub> {
    private FileTransferServiceFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private FileTransferServiceFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FileTransferServiceFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new FileTransferServiceFutureStub(channel, callOptions);
    }
  }

  private static final int METHODID_UPLOAD = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final FileTransferServiceImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(FileTransferServiceImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_UPLOAD:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.upload(
              (io.grpc.stub.StreamObserver<org.golangphilly.grpc.file.transfer.Server.FileResponse>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static final class FileTransferServiceDescriptorSupplier implements io.grpc.protobuf.ProtoFileDescriptorSupplier {
    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return org.golangphilly.grpc.file.transfer.Server.getDescriptor();
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (FileTransferServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new FileTransferServiceDescriptorSupplier())
              .addMethod(METHOD_UPLOAD)
              .build();
        }
      }
    }
    return result;
  }
}
