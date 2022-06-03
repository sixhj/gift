
// main 没有返回值
void main(List<String> arguments) {
  print('Hello world!');
  var name = 'Bob';
  dynamic name2 = 'Bob'; // 动态类型
  final name3 = 'Bob'; // Without a type annotation  不能被修改


  printInteger(1);

}


// 定义一个函数
printInteger(int aNumber) {
  print('The number is $aNumber.'); // 打印到控制台。
}
