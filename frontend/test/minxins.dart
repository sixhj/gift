

// minxins  的类不能有构造函数
// 只能继承object ，父类不能 extends 其它类

class Apple {


  printNameA(){
    print('apple');
  }

  hello(){
    print('hello apple');
  }


}


class Orange{

  printNameO(){
    print('orange');
  }

  hello(){
    print('hello orange');
  }
}



class Fruit with Orange,Apple{

}


main(){

  Fruit f = new Fruit();
  f.printNameA();
  f.printNameO();
  f.hello(); // hello orange ,和 with 的顺序有关，调用最后一个with 的同名方法

}