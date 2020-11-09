package Arrays;

public class Array {
    protected int[] numbers;
    protected int count;

    public Array(int length) {
        numbers = new int[length];
    }

    void insert(int item) {
        if (numbers.length == count) {
            int[] newNumbers = new int[count * 2];

            for (int i = 0; i < count; i++)
                newNumbers[i] = numbers[i];

            numbers = newNumbers;
        }

        numbers[count++] = item;
    }

    void removeAt(int index) {
        if (index < 0 || index >= count)
            throw new IllegalArgumentException();

        for (int i = index; i < count; i++)
            numbers[i] = numbers[i + 1];

        count--;
    }

    int indexOf(int item) {
        for (int i = 0; i < count; i++) {
            if (item == numbers[i])
                return i;
        }
        return -1;
    }

    void print() {
        for (int i = 0; i < count; i++)
            System.out.println(numbers[i]);
    }
}
