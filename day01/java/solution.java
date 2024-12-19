package cmiles74.adventofcode;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.InvalidParameterException;
import java.util.List;
import java.util.Arrays;
import java.util.ArrayList;
import java.util.concurrent.Callable;

class Solution {

    int[][] loadSample() {
        return new int[][] {
            {3, 4, 2, 1, 3, 3},
            {4, 3, 5, 3, 9, 3}
        };
    }

    int[][] loadInput(String filename) throws IOException {
        var list1 = new ArrayList<Integer>();
        var list2 = new ArrayList<Integer>();

        var lines = Files.readAllLines(Paths.get(filename));
        for (String line : lines) {
            String[] values = line.split("\\s+");
            list1.add(Integer.parseInt(values[0].trim()));
            list2.add(Integer.parseInt(values[1].trim()));
        }

        return new int[][] {
            list1.stream().mapToInt(i -> i).toArray(),
            list2.stream().mapToInt(i -> i).toArray()
        };
    }

    public static void main (String[] args) throws Exception {
        Solution solution = new Solution();
        //int[][] input = solution.loadSample();
        int[][] input = solution.loadInput("../input.txt");

        System.out.println("Part 1 - Sum of Distances:");
        var timer = new Timer<Integer>(new Part1(input));
        var result = timer.time();
        System.out.println(result.getResult());

        System.out.println("Part 2 - Sum of Occurrences:");
        var timer2 = new Timer<Integer>(new Part2(input));
        var result2 = timer2.time();
        System.out.println(result2.getResult());

        System.out.println("\n----");
        System.out.println("Part 1 completed " + result.getElapsed() + " ms");
        System.out.println("Part 2 completed " + result2.getElapsed() + " ms");
    }
}

class Part1 implements Callable<Integer> {
    int[] list1;
    int[] list2;

    Part1(int[][] input) {
        list1 = input[0];
        list2 = input[1];
    }

    public Integer call() throws InvalidParameterException {
        Arrays.sort(list1);
        Arrays.sort(list2);

        int distance = 0;
        for (int index = 0; index < list1.length; index++) {
            distance += Math.abs(list1[index] - list2[index]);
        }

        return distance;
    }
}

class Part2 implements Callable<Integer> {
    int[] list1;
    int[] list2;

    Part2(int[][] input) {
        list1 = input[0];
        list2 = input[1];
    }

    int countOccurs(int[] list, int searchValue) {
        int occurs = 0;
        for (int value : list) {
            if (value == searchValue) {
                occurs += 1;
            }

            if (value > searchValue) {
                break;
            }
        }

        return occurs;
    }

    public Integer call() throws InvalidParameterException {
        Arrays.sort(list2);

        int sumOccurs = 0;
        for (int value : list1) {
            sumOccurs += (value * countOccurs(list2, value));
        }

        return sumOccurs;
    }
}

class TimedResult<Double, T> {
    Double elapsed;
    T result;

    TimedResult(Double elapsed, T result) {
        this.elapsed = elapsed;
        this.result = result;
    }

    public Double getElapsed() {
        return elapsed;
    }

    public T getResult() {
        return result;
    }
}

class Timer<T> {
    Callable<T> callable;

    Timer(Callable<T> callable) {
        this.callable = callable;
    }

    public TimedResult<Double, T> time() throws Exception {
        var start = System.nanoTime();
        T result =  callable.call();
        var stop = System.nanoTime();
        var elapsed = (stop - start) / 1000000.0;
        return new TimedResult<Double, T>(elapsed, result);
    }
}
