void main() {
  print('hello');

  // var 可以类型推导
  var str = 'hello';
  var str2 = 'hello2';
  // str3 = 'hello2';   ❌ ，
  var i = 1;
  String str3 = 'hello3';

  print(str2);
  print(str3);

  int i2 = 3;
  print(i2);
  const A = 1; // 常量  一开始就要赋值
  final B; //  常量 final  可以赋值一次  ， 第一次使用时才初始化
  B = 2;

  print(B);

  //
  double d2 = 1.1;

  var f = true;

  if (f) {
    print('yes');
  } else {
    print('no');
  }


  var l2 = ['apple', false, 1];
  l2.add(2);
  print(l2);

  var l3 = <String>['a', 'b']; // 指定列表类型
  print(l3);


  var apple = {
    "color": "red",
    "price": 3,

  };
  print(apple);
  print(apple["color"]);

 // map
  var m1 = new Map();
  m1['color'] = 'yellow';

  // 类型判断
  if (m1 is Map) {
    print('map');
  }
}
