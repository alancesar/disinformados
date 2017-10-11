import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;

public class Main {

    public static void main(String[] args) throws IOException, InterruptedException, ExecutionException {
        long init = System.currentTimeMillis();

        String[] values = new String(Files.readAllBytes(Paths.get("numbers.txt"))).split("\n");
        int[] numbers = new int[values.length];

        for (int i = 0; i < values.length; i++) {
            numbers[i] = Integer.valueOf(values[i]);
        }

        ExecutorService executor = Executors.newFixedThreadPool(numbers.length);
        List<Future<Double>> responses = new ArrayList<>();

        for (int number : numbers) {
            responses.add(executor.submit(() -> calc(number)));
        }

        executor.shutdown();
        executor.awaitTermination(30, TimeUnit.SECONDS);

        for (Future<Double> response : responses) {
            System.out.println(response.get());
        }

        long end = System.currentTimeMillis();
        System.out.println("Tempo total: " + (end - init) + " milissegundos");
    }

    // (x² * √x) / (√x² + 1)
    public static double calc(int number) {
        return (Math.pow(number, 2) * Math.sqrt(number)) / (Math.sqrt(Math.pow(number, 2) + 1));
    }
}
