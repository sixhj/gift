class Person {
  late String name;

  late num age;

  Person(this.name, this.age);

  void run() {
    print("person ${this.name}  runing ");
  }
}

class Student extends Person {
  late num id;

  Student(String name, num age, num id) : super(name, age){
    this.id = id;
  }

  // Student(String name, num age,num id ) : super(name, age);

  @override
  void run() {
    // TODO: implement run
    // super.run();
    print("student=${this.name}  id=${this.id} runing ");
  }
}

main() {
  print('hello');

  Person p = new Person('a', 3);
  p.run();

  Student s = new Student('b', 4, 001);

  s.run();
}
