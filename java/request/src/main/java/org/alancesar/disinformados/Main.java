package org.alancesar.disinformados;

import org.springframework.web.client.RestTemplate;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.*;

public class Main {

    private static final String[] URLS = { "http://localhost:3000", "http://localhost:3001" };
    private static final int REQUESTS = 100;

    public static void main(String args[]) throws InterruptedException, ExecutionException {
        long init = System.currentTimeMillis();

        ExecutorService executor = Executors.newFixedThreadPool(REQUESTS * URLS.length);
        List<Future<String>> responses = new ArrayList<>();

        for (int i = 0; i < REQUESTS; i++) {
            for (String url : URLS) {
                responses.add(executor.submit(() -> new RestTemplate().getForObject(url, String.class)));
            }
        }

        executor.shutdown();
        executor.awaitTermination(30, TimeUnit.SECONDS);

        for (Future<String> response : responses) {
            System.out.println(response.get());
        }

        long end = System.currentTimeMillis();
        System.out.println("Tempo total -> " + (end - init));
    }

}
