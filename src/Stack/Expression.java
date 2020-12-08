package Stack;

import java.util.Stack;

public class Expression {

    char[] openingChars = {'(', '{', '<', '['};
    char[] closingChars = {')', '}', '>', ']'};

    public boolean isBalanced(String input) {
        if (input == null) throw new IllegalArgumentException();

        Stack<Character> stack = new Stack<>();

        for (char character : input.toCharArray()) {

            if (isOpeningChars(character))
                stack.push(character);


            if (isClosingChars(character)) {
                if (stack.empty()) return false;
                char openingChar = stack.peek();

                if (!isMatch(openingChar, character)) return false;
                stack.pop();
            }
        }

        return stack.empty();
    }

    private boolean isOpeningChars(char character) {
        for (char openingChar : openingChars) if (character == openingChar) return true;
        return false;
    }

    private boolean isClosingChars(char character) {
        for (char closingChar : closingChars) if (character == closingChar) return true;
        return false;
    }

    private boolean isMatch(char openingChar, char closingChar) {
        for (int i = 0; i < openingChars.length; i++) {
            if (openingChar == openingChars[i] && closingChar == closingChars[i])
                return true;
        }
        return false;
    }
}
