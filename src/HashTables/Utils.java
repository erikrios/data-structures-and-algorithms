package HashTables;

import java.util.HashMap;
import java.util.Map;

public class Utils {
    public static char findTheFirstNonRepeatedCharacter(String sentences) {
        Map<Character, Boolean> chars = new HashMap<>();
        String lowerCaseSentences = sentences.toLowerCase();

        for (char character : lowerCaseSentences.toCharArray()) {
            if (!chars.containsKey(character))
                chars.put(character, true);
            else
                chars.put(character, false);
        }

        for (char character : lowerCaseSentences.toCharArray()) {
            if (chars.get(character))
                return character;
        }
        return '1';
    }
}
