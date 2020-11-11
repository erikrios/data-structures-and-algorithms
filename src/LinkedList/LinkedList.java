package LinkedList;

public class LinkedList {
    private Node first;
    private Node last;

    void addFirst(int item) {
        Node node = new Node(item);

        if (first == null)
            first = last = node;
        else {
            node.next = first;
            first = node;
        }
    }

    void addLast(int item) {
        Node node = new Node(item);

        if (first == null)
            first = last = node;
        else {
            last.next = node;
            last = node;
        }
    }

    private class Node {
        private int value;
        private Node next;

        public Node(int value) {
            this.value = value;
        }
    }
}
