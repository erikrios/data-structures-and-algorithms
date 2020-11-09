package TheBigONotation;

/**
 * Time Complexity 0(n^2) / Quadratic
 */
public class Quadratic {
    public void log(int[] numbers) {
        for (int first : numbers)
            for (int second : numbers)
                System.out.println(second);
    }
}
