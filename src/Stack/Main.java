package Stack;

public class Main {
    public static void main(String[] args) {
        String name = "Erik Rio Setiawan";
        StringReverser reverser = new StringReverser();
        System.out.println(reverser.reverse(name));

        String balancedString = "(([1] + <2>)) [a]";
        Expression expression = new Expression();
        System.out.println(expression.isBalanced(balancedString));
    }
}
