


// 使用抽象类定义接口
// 接口：抽象类
// 不能实例化
// 使用 implements 实现接口
// implements 接口时，  需要实现全部的方法
// 接口中的方法没有方法体




abstract  class DataBase{
  conn();
  open();
  close();
}



// 实现接口的所有方法
class Mysql implements DataBase{
  @override
  close() {
    // TODO: implement close
    // throw UnimplementedError();
    print('mysql close ok');
  }

  @override
  conn() {
    // TODO: implement conn
    // throw UnimplementedError();
    print('mysql conn ok');
  }

  @override
  open() {
    // TODO: implement open
    // throw UnimplementedError();
    print( 'mysql open ok');
  }

}


main(){


}
