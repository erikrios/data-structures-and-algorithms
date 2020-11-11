package LinkedList;

public class LinkedList {
    private Node first;
    private Node last;

    public void addFirst(int item) {
        Node node = new Node(item);

        if (isEmpty())
            first = last = node;
        else {
            node.next = first;
            first = node;
        }
    }

    public void addLast(int item) {
        Node node = new Node(item);

        if (isEmpty())
            first = last = node;
        else {
            last.next = node;
            last = node;
        }
    }

    public int indexOf(int item) {
        int index = 0;
        Node current = first;

        while (current != null) {
            if (current.value == item)
                return index;
            current = current.next;
            index++;
        }

        return -1;
    }

    private boolean isEmpty() {
        return first == null;
    }

    private class Node {
        private int value;
        private Node next;

        public Node(int value) {
            this.value = value;
        }
    }
}
