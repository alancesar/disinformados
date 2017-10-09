package org.alancesar.disinformados;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.URL;
import java.net.URLConnection;
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;

public class Main {

    private static final String[] URLS = { "http://localhost:3000", "http://localhost:3001" };
    private static final int REQUESTS = 100;

    public static void main(String args[]) throws InterruptedException, ExecutionException {
        long init = System.currentTimeMillis();

        ExecutorService executor = Executors.newFixedThreadPool(REQUESTS * URLS.length);
        List<Future<String>> responses = new ArrayList<>();

        for (int i = 0; i < REQUESTS; i++) {
            for (String url : URLS) {
                responses.add(executor.submit(() -> new Request().ping(url)));
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

class Request {

    public String ping(String url) throws IOException {

        StringBuilder response = new StringBuilder();
        URLConnection connection = new URL(url).openConnection();
        BufferedReader reader = new BufferedReader(
                new InputStreamReader(connection.getInputStream()));

        String line;

        while ((line = reader.readLine()) != null) {
            response.append(line);
        }

        reader.close();
        return response.toString();
    }
}
