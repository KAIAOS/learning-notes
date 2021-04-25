//----构造函数方式
// 在子类型构造函数的内部调用超类型构造函数；使用 apply() 或 call() 方法将父对象的构造函数绑定在子对象上
function SuperType(){
    // 定义引用类型值属性
    this.colors = ["red","green","blue"];
}

function SubType(){
    // 继承 SuperType，在这里还可以给超类型构造函数传参
    SuperType.call(this);
}
var instance1 = new SubType();
instance1.colors.push("purple")

//----原型混合方式
function SuperType(name){
    this.name = name;
    this.colors = ["red","green","blue"];
}
SuperType.prototype.sayName = function(){
    alert(this.name);
};
function SubType(name,age){
    // 借用构造函数方式继承属性
    SuperType.call(this,name);
    this.age = age;
}
// 原型链方式继承方法
SubType.prototype = new SuperType();
SubType.prototype.sayAge = function(){
    alert(this.age);
};
var instance1 = new SubType("luochen",22);
instance1.colors.push("purple");
alert(instance1.colors);      // "red,green,blue,purple"
instance1.sayName();
instance1.sayAge();