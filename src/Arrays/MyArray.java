package Arrays;

public class MyArray extends Array {

    public MyArray(int length) {
        super(length);
    }

    int max() {
        int largest = numbers[0];

        for (int i = 0; i < count; i++)
            if (numbers[i] > largest)
                largest = numbers[i];

        return largest;
    }

    Array intersect(Array others) {
        Array intersection = new Array(count);

        for (int number : numbers)
            if (others.indexOf(number) >= 0)
                intersection.insert(number);

        return intersection;
    }

    void reverse() {
        int[] reversedNumbers = new int[count];

        for (int i = 0; i < count; i++)
            reversedNumbers[i] = numbers[count - 1 - i];

        numbers = reversedNumbers;
    }

    void insertAt(int item, int index) {
        if (index < 0 || index >= count)
            throw new IllegalArgumentException();

        int[] newNumbers = new int[count];

        if (numbers.length == count)
            newNumbers = new int[count * 2];

        count++;
        for (int i = 0; i < count; i++) {
            if (i == index)
                newNumbers[index] = item;
            else if (i < index)
                newNumbers[i] = numbers[i];
            else
                newNumbers[i] = numbers[i - 1];
        }

        numbers = newNumbers;
    }
}
