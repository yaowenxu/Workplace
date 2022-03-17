// 结构体派生
#[derive(Debug)]
struct Rectangle {
    width: u32,
    length: u32,
}

impl Rectangle {
    // 计算长方形面积
    fn area(&self) -> u32 {
        println!("{:?}", self);
        return self.width * self.length;
    }
    // 重新设置长和宽
    fn set(&mut self) {
        self.width = 40;
        self.length = 50;
    }
}

// 可以 impl 多个代码块;
impl Rectangle {
    // 关联函数，常用于构造函数
    fn square(size: u32) -> Rectangle {
        let t = Rectangle {
            width: size,
            length: size,
        };
        println!("{:#?}", t);
        return t; // 所有权转移
    }
}

// 枚举
#[derive(Debug)]
enum IPKind {
    V4,
    V6,
}

// 入口函数
fn main() {
    let name = "yaowen";
    let name = name;
    let guess: isize = "666".parse().expect("not a number");
    println!("Hello, {} {}!", name, guess);
    function();
    let mut rect = Rectangle {
        width: 30,
        length: 50,
    };
    rect.set();
    println!("Area: {}", rect.area());
    println!("{:?}", Rectangle::square(60));
    println!("{:#?}", IPKind::V4);
    println!("{:#?}", IPKind::V6);
    option();
    match_coin(Coin::Penny);
    match_coin(Coin::Nickel);
    match_coin(Coin::Dime);
    match_coin(Coin::Quarter);
}

// 新的函数
fn function() -> i32 {
    println!("Hello, func!");

    // 代码块， 表达式；
    let state = {
        let mut x = 1;
        println!("step 2 {}", x);
        x = x + 1;
        x
    };
    println!("step 1 {}", state);
    return 0;
}

fn option() {
    let some_i = Some(0);
    let some_x = match some_i {
        None => None,
        Some(i) => Some(i + 1),
    };
    // 更少的匹配，可以看作是 match 语句的语法糖；
    if let Some(1) = some_x {
        println!("if let match Some(1)");
    }
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn match_coin(coin: Coin) -> u32 {
    return match coin {
        Coin::Penny => {
            println!("match Penny");
            1
        }
        Coin::Nickel => {
            println!("match Nickel");
            5
        }
        _ => 10,
    };
}
