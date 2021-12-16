use std::collections::HashMap;
use std::num::ParseIntError;
use std::path::PathBuf;

fn commonality(col: usize, map: &HashMap<String, bool>) -> (usize, usize) {
    let mut ones: usize = 0;
    let mut zero: usize = 0;
    for (diag, _) in map {
        if diag.chars().nth(col).unwrap() == '1' {
            ones += 1;
        } else {
            zero += 1;
        }
    }
    return (ones, zero);
}

fn reduce_readings(common: bool, length: usize, mut readings: HashMap<String, bool>) -> String {
    //let mut readings = map.clone();
    for col in 0..length {
        if readings.len() == 1 {
            break;
        }
        let (ones, zeroes) = commonality(col, &readings);
        let guide = if common {
            if ones >= zeroes {
                '1'
            } else {
                '0'
            }
        } else {
            if ones < zeroes {
                '1'
            } else {
                '0'
            }
        };
        for (reading, _) in readings.clone() {
            if reading.chars().nth(col).unwrap() != guide {
                readings.remove(&reading);
            }
        }
    }

    return readings.keys().next().unwrap().clone();
}

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    let mut diags: HashMap<String, bool> = HashMap::new();

    // read diagnostics report into a matrix
    let lines = crate::io::read_lines(path).unwrap();
    let length = lines[0].len();
    for line in lines {
        diags.insert(line, true);
    }

    // part 1
    let mut most_common: String = String::new();
    let mut least_common: String = String::new();
    for col in 0..length {
        let (ones, zeroes) = commonality(col, &diags);
        most_common.push(if ones >= zeroes { '1' } else { '0' });
        least_common.push(if ones < zeroes { '1' } else { '0' });
    }

    let gamma = isize::from_str_radix(&most_common, 2).unwrap();
    let epsilon = isize::from_str_radix(&least_common, 2).unwrap();

    println!("Most common digits: {} ({})", most_common, gamma);
    println!("Least common digits: {} ({})", least_common, epsilon);
    println!("Power consumption: {}", gamma * epsilon);

    // part 2
    let o2_readings = reduce_readings(true, length, diags.clone());
    let o2_gen_rating = isize::from_str_radix(o2_readings.as_str(), 2).unwrap();
    let co2_readings = reduce_readings(false, length, diags.clone());
    let co2_scrub_rating = isize::from_str_radix(co2_readings.as_str(), 2).unwrap();

    println!("O2 Generator Rating: {} ({})", o2_readings, o2_gen_rating);
    println!(
        "CO2 Scrubber Rating: {} ({})",
        co2_readings, co2_scrub_rating
    );
    println!("Life Support Rating: {}", o2_gen_rating * co2_scrub_rating);

    Ok(())
}
