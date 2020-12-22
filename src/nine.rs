use std::collections::HashSet;
use std::collections::VecDeque;
use std::num::ParseIntError;
use std::path::PathBuf;

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    if let Ok(lines) = crate::io::lines(path) {
        let numbers: Vec<u64> = lines.iter().map(|x| x.parse().unwrap()).collect();
        let mut slice: VecDeque<u64> = VecDeque::new();
        let mut idx: usize = 0;
        let mut suspect = 0;
        for number in &numbers {
            // preamble
            if slice.len() <= 25 {
                slice.push_back(*number);
                continue;
            } else {
                slice.pop_front();
            }
            let sums = calculate_sums(&slice);
            if !sums.contains(number) {
                println!(
                    "{} (at {}) is not a sum of any of the previous 25 numbers",
                    number, idx
                );
                suspect = number.clone();
                break;
            }
            slice.push_back(*number);
            idx += 1;
        }

        let mut found = false;
        for length in 2..=idx {
            for offset in 0..idx {
                let slice = &numbers[offset..(offset + length)];
                let sum: u64 = slice.iter().sum();
                if sum == suspect {
                    println!("found contiguous range to sum {}: {:?}", sum, slice);
                    let min = slice.iter().min().unwrap();
                    let max = slice.iter().max().unwrap();
                    println!("min {} + max {} = {}", min, max, *min + *max);
                    found = true;
                    break;
                }
            }
            if found {
                break;
            }
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
