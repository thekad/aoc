use std::num::ParseIntError;
use std::path::PathBuf;

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    let mut counter: i32 = 0;
    let mut counter3: i32 = 0;
    let mut prev: i32 = 0;
    let mut prev3: i32 = 0;
    let mut readings: Vec<i32> = Vec::new();

    // read it into a Vec first to slice it down below
    if let Ok(lines) = crate::io::read_lines(path) {
        for line in lines {
            if let Ok(num) = line.parse() {
                readings.push(num);
            }
        }
    }

    for (idx, reading) in readings.iter().enumerate() {
        if prev != 0 && reading > &prev {
            println!("Individual reading increased from {} to {}", prev, reading);
            counter += 1;
        }
        prev = *reading;

        // only read a triplet if there's enough of them
        if idx + 3 <= readings.len() {
            let triplet: &[i32] = &readings[idx..idx + 3];
            let sum3: i32 = triplet.iter().sum();
            if prev3 != 0 && sum3 > prev3 {
                println!("Triplet reading increased from {} to {}", prev3, sum3);
                counter3 += 1;
            }
            prev3 = sum3;
        }
    }

    println!("Individual depth readings increased {:?} times", counter);
    println!("Triplet depth readings increased {:?} times", counter3);
    Ok(())
}
