>我们知道，java当中，类的加载顺序是：类静态块-类静态属性-类内部属性-类构造方法。
>但是当有内部类的时候会怎样呢？我们先看一下代码。

~~~ java
public class Singleton {

    public static class Inner{
        static {
            System.out.println("TestInner Static!");
        }
        public final static Singleton testInstance = new Singleton(3);
    }

    public static Singleton getInstance(){
        return Inner.testInstance;
    }

    public Singleton(int i ) {
        System.out.println("Test " + i +" Construct! ");
    }

    //类静态块
    static {
        System.out.println("Test Static");
    }

    //类静态属性
    public static Singleton testOut = new Singleton(1);

    public static void main(String args[]){
        Singleton t = new Singleton(2);
        Singleton.getInstance();
    }

}
~~~

