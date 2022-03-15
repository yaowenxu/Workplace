fn main() {
    let name = "yaowen";
    let name = name;
    let guess: isize =  "666".parse().expect("not a number");
    println!("Hello, {} {}!", name, guess);
    function();
}

// 新的函数
fn function() -> i32{
    println!("Hello, func!");
    let state = {
        let mut x = 1;
        println!("step 2 {}", x);
        x = x + 1;
        x
    };
    println!("step 1 {}", state);
    return 0
}