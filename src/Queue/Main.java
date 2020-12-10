package Queue;

public class Main {
    public static void main(String[] args) {
        PriorityQueue queue = new PriorityQueue(5);
        queue.add(5);
        queue.add(3);
        queue.add(8);
        queue.add(2);
        queue.add(1);
        System.out.println(queue);
    }
}
