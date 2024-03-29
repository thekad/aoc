use std::num::ParseIntError;
use std::path::PathBuf;
use itertools::Itertools;

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    if let Ok(lines) = crate::io::lines(path) {
        let adapters: Vec<u8> = lines.iter().map(|x| x.parse().unwrap()).collect();
        let mut differences: Vec<u8> = Vec::new();
        let mut next: u8 = 0;
        loop {
            if let Some(found) = find_next_best_adapter(next, &adapters) {
                differences.push(found - next);
                next = found;
            } else {
                break;
            }
        }
        let diff1: Vec<u8> = differences.iter().map(|x| *x).filter(|x| *x == 1).collect();
        let mut diff3: Vec<u8> = differences.iter().map(|x| *x).filter(|x| *x == 3).collect();
        // account for the last adapter which is always a diff of 3
        diff3.push(3);
        println!(
            "Differences of 1 ({}) * differences of 3 ({}) = {}",
            diff1.len(),
            diff3.len(),
            diff1.len() * diff3.len()
        );

        find_all_adapters(&adapters);
    }

    Ok(())
}

fn find_next_best_adapter(joltage: u8, adapters: &Vec<u8>) -> Option<u8> {
    let possible = find_next_adapters(joltage, adapters);

    if possible.len() > 0 {
        return Some(*possible.iter().min().unwrap());
    }

    None
}

fn find_all_adapters(adapters: &Vec<u8>) {
    for permutation in adapters.iter().permutations(adapters.len()).unique(){
        println!("{:?}", permutation);
    }
}

fn find_next_adapters(joltage: u8, adapters: &Vec<u8>) -> Vec<u8> {
    let possible: Vec<u8> = adapters
        .iter()
        .map(|x| *x)
        .filter(|x| *x > joltage)
        .filter(|x| x - joltage <= 3)
        .collect();

    possible
}
