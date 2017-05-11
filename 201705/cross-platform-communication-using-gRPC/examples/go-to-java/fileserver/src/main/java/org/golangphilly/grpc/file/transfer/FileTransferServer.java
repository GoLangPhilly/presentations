package org.golangphilly.grpc.file.transfer;

import com.google.protobuf.ByteString;
import io.grpc.stub.StreamObserver;

import java.io.FileOutputStream;


/**
 * Created by robertojrojas on 4/8/17.
 */
public class FileTransferServer extends FileTransferServiceGrpc.FileTransferServiceImplBase {

    private String rootDir;

    public FileTransferServer(String rootDir) {
        System.out.println("Root Dir: " + rootDir);
        this.rootDir = rootDir;
    }

    @Override
    public StreamObserver<Server.FileRequest> upload(final StreamObserver<Server.FileResponse> responseObserver) {

        return new StreamObserver<Server.FileRequest>() {

            private ByteString result;
            private String filename;

            @Override
            public void onNext(Server.FileRequest fileRequest) {
                if (this.result == null) {
                    this.result = fileRequest.getData();
                } else {
                    this.result = this.result.concat(fileRequest.getData());
                }
                System.out.println("Received messages with " +
                        fileRequest.getData().size() + " bytes");

                if (this.filename == null) {
                    this.filename = fileRequest.getFilename();
                    System.out.println("setting filename: " + this.filename);
                }
            }

            @Override
            public void onError(Throwable throwable) {
                System.err.println(throwable);
            }

            @Override
            public void onCompleted() {
                System.out.println("Total bytes received: " + result.size());
                boolean uploadSuccessful = false;
                try {
                    System.out.println("Creating file: " + this.filename);
                    String fullPath = String.format("%s/%s", FileTransferServer.this.rootDir, this.filename);
                    FileOutputStream fo = new FileOutputStream(fullPath);
                    fo.write(this.result.toByteArray());
                    fo.close();
                    uploadSuccessful = true;
                } catch (Exception e) {
                    System.err.println(e);
                }

                responseObserver.onNext(
                        Server.FileResponse.newBuilder()
                                .setFilename(this.filename)
                                .setSize(this.result.size())
                                .setIsOk(uploadSuccessful)
                                .build()

                );
                responseObserver.onCompleted();
            }
        };
    }
}
