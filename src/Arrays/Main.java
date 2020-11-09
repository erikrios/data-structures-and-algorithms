package Arrays;

public class Main {
    public static void main(String[] args) {
        MyArray numbers = new MyArray(5);
        numbers.insert(1);
        numbers.insert(3);
        numbers.insert(5);
        numbers.insert(9);
        numbers.insert(10);

        numbers.print();
    }
}
