package Queue;

public class Main {
    public static void main(String[] args) {
        StackQueue stackQueue = new StackQueue();
        stackQueue.enqueue(1);
        stackQueue.enqueue(2);
        stackQueue.enqueue(3);
        stackQueue.enqueue(4);
        stackQueue.enqueue(5);
        int dequeue = stackQueue.dequeue();
        int peek = stackQueue.peek();
        System.out.println(stackQueue.isEmpty());
        System.out.println(dequeue);
        System.out.println(peek);
    }
}
