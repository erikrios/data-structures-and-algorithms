package Queue;

import java.util.Stack;

public class StackQueue {
    private final Stack<Integer> firstStack = new Stack<>();
    private final Stack<Integer> secondStack = new Stack<>();

    public void enqueue(int item) {
        firstStack.push(item);
    }

    public int dequeue() {
        if (isEmpty()) throw new IllegalStateException();
        moveFirstStackToSecondStack();
        return secondStack.pop();
    }

    public Integer peek() {
        if (isEmpty()) throw new IllegalStateException();
        moveFirstStackToSecondStack();
        return secondStack.peek();
    }

    public boolean isEmpty() {
        return firstStack.isEmpty() && secondStack.isEmpty();
    }

    private void moveFirstStackToSecondStack() {
        if (secondStack.isEmpty()) reverse();
    }

    private void reverse() {
        while (!firstStack.isEmpty())
            secondStack.push(firstStack.pop());
    }
}
