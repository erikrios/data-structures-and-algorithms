package HashTables;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;

public class Main {
    public static void main(String[] args) {
        Map<Integer, String> map = new HashMap<>();

        map.put(1, "Erik");
        map.put(2, "Rio");
        map.put(3, "Setiawan");
        System.out.println(map);

        for (Map.Entry<Integer, String> item : map.entrySet()) {
            System.out.println(item.getValue());
        }

        String name = "Erik Rio Setiawan";
        char firstNonRepeatedNameCharacter = Utils.findTheFirstNonRepeatedCharacter(name);
        System.out.println(firstNonRepeatedNameCharacter);

        Set<Integer> set = new HashSet<>();
        int[] numbers = {1, 2, 3, 3, 2, 1, 4};
        for (int number : numbers)
            set.add(number);
        System.out.println(set);
    }
}
