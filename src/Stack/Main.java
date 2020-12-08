package Stack;

public class Main {
    public static void main(String[] args) {
        String name = "Erik Rio Setiawan";
        StringReverser reverser = new StringReverser();
        System.out.println(reverser.reverse(name));

        String balancedString = "(([1] + <2>)) [a]";
        Expression expression = new Expression();
        System.out.println(expression.isBalanced(balancedString));

        Stack stack = new Stack();
        stack.push(1);
        stack.push(2);
        stack.push(3);
        int popValue = stack.pop();
        int peekValue = stack.peek();
        System.out.println(popValue);
        System.out.println(peekValue);
        System.out.println(stack.isEmpty());
        System.out.println(stack.min());
        System.out.println(stack);
    }
}
