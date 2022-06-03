// 抽象类 abstract
// extends 继承抽象类，可以复用抽象类里面的默认方法

abstract class Animal {
  eat(); // 抽象方法
  int getNum();

  printInfo() {
    print('抽象类里面的公共方法');
  }
}


class Dog extends Animal {
  late String name;

  Dog(this.name);

  @override
  eat() {
    print("${this.name} eating");
  }

  @override
  int getNum() {

    return 1;
  }
}

main() {
  // 多态
  Animal d = Dog('a');
  d.eat();
  d.printInfo();
}
