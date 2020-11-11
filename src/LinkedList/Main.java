package LinkedList;

public class Main {

    public static void main(String[] args) {
        LinkedList list = new LinkedList();
        list.addLast(10);
        list.addLast(15);
        list.addLast(8);
        list.addFirst(9);
        list.removeFirst();
        list.removeLast();
        System.out.println(list.indexOf(15));
        System.out.println(list.contains(10));
        System.out.println(list.size());
    }
}
