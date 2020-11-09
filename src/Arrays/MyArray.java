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

    int[] intersect(int[] others) {
        int[] longestNumbers;
        int[] shortestNumbers;

        if (count > others.length) {
            longestNumbers = new int[count];
            for (int i = 0; i < count; i++)
                longestNumbers[i] = numbers[i];

            shortestNumbers = others;
        } else {
            longestNumbers = others;

            shortestNumbers = new int[count];
            for (int i = 0; i < count; i++)
                shortestNumbers[i] = numbers[i];
        }

        int[] numbers = new int[shortestNumbers.length];
        int index = 0;

        for (int i = 0; i < shortestNumbers.length; i++)
            for (int j = 0; j < longestNumbers.length; j++) {
                if (shortestNumbers[i] == longestNumbers[j])
                    numbers[index++] = shortestNumbers[i];
            }

        int[] intersectNumbers = new int[index];

        for (int i = 0; i < index; i++)
            intersectNumbers[i] = numbers[i];

        return intersectNumbers;
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
