### 1

```java
public static boolean addAll(Collection<T>, T...elements);//添加多个元素
public static void shuffle(List<?> list);//打乱顺序
public static <T> void sort(List<T> list);//将集合按默认升序排序
被排序的集合必须实现Comparable接口种的compareTo方法 

    
public static <T> void sort(List<T> list, Comparator<? super T>);
Comparator相当于找裁判 比较两个
```

