package Stack;

import java.util.Arrays;

public class Stack {
    private int[] items;
    private int length;
    private final int capacity = 1;

    public Stack() {
        items = new int[capacity];
    }

    public void push(int item) {
        resizeIfRequired();
        items[length++] = item;
    }

    public int pop() {
        if (isEmpty()) throw new IllegalStateException();
        int popValue = items[length - 1];
        items[--length] = 0;
        return popValue;
    }

    public int peek() {
        if (isEmpty()) throw new IllegalStateException();
        return items[length - 1];
    }

    public int min() {
        if (isEmpty()) throw new IllegalStateException();
        int min = items[0];
        for (int i = 0; i < length; i++) {
            if (items[i] < min) min = items[i];
        }
        return min;
    }

    public int[] toArray() {
        int[] result = new int[length];

        for (int i = 0; i < length; i++)
            result[i] = items[i];

        return result;
    }

    public boolean isEmpty() {
        return length == 0;
    }

    private void resizeIfRequired() {
        if (length >= capacity) {
            int[] newItems = new int[length * 2];
            for (int i = 0; i < length; i++)
                newItems[i] = items[i];

            items = newItems;
        }
    }

    @Override
    public String toString() {
        return Arrays.toString(toArray());
    }
}
