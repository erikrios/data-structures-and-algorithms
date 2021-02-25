package HashTables;

import java.util.HashMap;
import java.util.Map;

public class Utils {
    public static char findTheFirstNonRepeatedCharacter(String sentences) {
        Map<Character, Integer> map = new HashMap<>();
        String lowerCaseSentences = sentences.toLowerCase();
        char[] chars = lowerCaseSentences.toCharArray();

        for (char character : chars) {
            int count = map.containsKey(character) ? map.get(character) : 0;
            map.put(character, count + 1);
        }

        for (char character : chars) {
            if (map.get(character) == 1)
                return character;
        }

        return Character.MIN_VALUE;
    }
}
