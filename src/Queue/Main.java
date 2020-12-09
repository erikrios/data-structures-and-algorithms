package Queue;

public class Main {
    public static void main(String[] args) {
        ArrayQueue queue = new ArrayQueue(5);
        queue.enqueue(1);
        queue.enqueue(2);
        queue.enqueue(3);
        queue.dequeue();
        int front = queue.dequeue();
        queue.enqueue(4);
        queue.enqueue(5);
        queue.enqueue(6);
        queue.enqueue(7);
        queue.dequeue();
        queue.enqueue(8);
        System.out.println(front);
        System.out.println(queue.peek());
        System.out.println(queue.isEmpty());
        System.out.println(queue.isFull());
        System.out.println(queue);
    }
}
