use std::collections::HashSet;
use std::collections::VecDeque;
use std::num::ParseIntError;
use std::path::PathBuf;

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    if let Ok(lines) = crate::io::lines(path) {
        let mut numbers: VecDeque<u64> = VecDeque::new();
        for line in lines {
            let number: u64 = line.parse().unwrap();
            // preamble
            if numbers.len() <= 25 {
                numbers.push_back(number);
                continue;
            } else {
                numbers.pop_front();
            }
            let sums = calculate_sums(&numbers);
            if !sums.contains(&number) {
                println!("{} is not a sum of any of the previous numbers", number);
                break;
            }
            numbers.push_back(number);
        }
    }

    Ok(())
}

fn calculate_sums(numbers: &VecDeque<u64>) -> HashSet<u64> {
    let mut result: HashSet<u64> = HashSet::new();
    let mut idx = 0;
    let len = numbers.len();
    for first in numbers {
        idx += 1;
        for second in idx..len {
            result.insert(first + numbers[second]);
        }
    }

    result
}
