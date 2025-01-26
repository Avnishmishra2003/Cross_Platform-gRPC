import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import calculator.CalculatorGrpc;
import calculator.SquareRequest;
import calculator.SquareResponse;


public class CalculatorClient {
    public static void main(String[] args) {
        // Create a channel to the server
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 50051)
                .usePlaintext()
                .build();

        // Create a stub for making RPC calls
        CalculatorGrpc.CalculatorBlockingStub stub = CalculatorGrpc.newBlockingStub(channel);

        // Create a request
        SquareRequest request = SquareRequest.newBuilder().setNumber(5).build();

        // Call the Square method
        SquareResponse response = stub.square(request);

        // Print the response
        System.out.println("Square of 5: " + response.getResult());

        // Shutdown the channel
        channel.shutdown();
    }
}
