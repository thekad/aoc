use std::num::ParseIntError;
use std::path::PathBuf;

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    let mut depth1: i32 = 0;
    let mut depth2: i32 = 0;
    let mut horiz: i32 = 0;
    let mut aim: i32 = 0;

    // get a vector of commands
    if let Ok(lines) = crate::io::read_lines(path) {
        for line in lines {
            let command: Vec<&str> = line.split_whitespace().collect();
            let instr = command[0];
            let value: i32 = command[1].parse().unwrap();
            println!("Nav Command: {:?}", command);

            match instr {
                "up" => {
                    depth1 -= value;
                    aim -= value;
                }
                "down" => {
                    depth1 += value;
                    aim += value;
                }
                "forward" => {
                    horiz += value;
                    depth2 += aim * value;
                }
                _ => panic!("this command shouldn't be here"),
            }
        }
    }

    println!("Depth position: {}", depth1);
    println!("Horizontal position: {}", horiz);
    println!("First result: {}", depth1 * horiz);
    println!("Second result: {}", depth2 * horiz);

    Ok(())
}
