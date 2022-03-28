# Java 中的单例实现

在这篇文章中，我们来介绍在 Java 中如何正确的实现一个单例模式。

## 1. 饿汉式单例（不推荐）

```java
public class Singleton {

    private static Singleton INSTANCE = new Singleton();

    private String info = "Initial info class";

    private Singleton() {} // 私有构造器
    
    public static Singleton getInstance() {
        return INSTANCE;
    }
}
```

饿汉式单例在变量初始化时即对其进行一次实例化，保证了单例的唯一性，但是有一个缺点是如果是实例化一个大对象且不一定会用到此实例时，这种实现就会造成较大的资源浪费，因此不推荐此方式。

> 除上面介绍的实现方式均为懒汉式单例实现，即在需要时才会实例化对象。

## 2. 最常见的实现模式（不推荐)

```java
public class Singleton {

    private static Singleton INSTANCE;    // 唯一的私有的静态变量

    private String info = "Initial info class";

    private Singleton() {} // 私有构造器
    
    public static Singleton getInstance() {
        if (INSTANCE == null) {
            INSTANCE = new Singleton();
        }
        return INSTANCE;
    }
}
```

此写法是在Java中最常见的一种单例实现，但是由于在并发情况下
会出现意想不到的问题，从而导致会出现多个不同的实例被创建（当然也可以通过加锁实现并发安全，但是影响效率），所以不推荐使用此写法。

## 3. Enum 实现（推荐）

```java
public enum Singleton {

    INSTANCE("Initial info class");

    private String info;

    private Singleton(String info) {
        this.info = info;
    }

    public Singleton getInstance() {
        return INSTANCE;
    }
}
```

使用枚举来实现单例是一种不错的单例实现方式，这种实现不仅是懒加载的，而且是并发安全的，它借由 Java 枚举本身的机制保证变量只会被实例化一次，是较为推荐的一种实现方式。
