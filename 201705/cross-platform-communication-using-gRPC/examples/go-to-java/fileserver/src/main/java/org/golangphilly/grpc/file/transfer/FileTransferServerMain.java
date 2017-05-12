package org.golangphilly.grpc.file.transfer;

import io.grpc.ServerBuilder;
import io.grpc.Server;

/**
 * Created by robertojrojas on 4/8/17.
 */
public class FileTransferServerMain {

    private Server server;

    public static void main(String[] args) {

        try {
            FileTransferServerMain server = new FileTransferServerMain();
            server.start(args);
        } catch (Exception e) {
            System.err.println(e);
        }

    }

    private void start(String[] args) throws Exception {

        final int port = 9080;

        String rootDir = "/tmp";
        if (args.length >= 1) {
            rootDir = args[0];
        }
        FileTransferServer svr = new FileTransferServer(rootDir);

        server = ServerBuilder
                .forPort(port)
                .addService(svr)
                .build()
                .start();
        System.out.println(String.format("Listening on port %d", port));

        Runtime.getRuntime().addShutdownHook(new Thread(){
            @Override
            public void run() {
                System.out.println("Shutting down server");
                FileTransferServerMain.this.stop();
            }
        });

        server.awaitTermination();

    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

}
