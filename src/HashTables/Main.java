package HashTables;

import java.util.HashMap;
import java.util.Map;

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
    }
}
