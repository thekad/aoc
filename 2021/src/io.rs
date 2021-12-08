use std::fs::File;
use std::io::{self, BufRead, BufReader};
use std::path::Path;

// this returns a vector of strings
pub fn read_lines(filename: impl AsRef<Path>) -> io::Result<Vec<String>> {
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    Ok(buf
        .lines()
        .map(|l| l.expect("Could not parse line"))
        .collect())
}
