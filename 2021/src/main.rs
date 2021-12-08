mod day01;
mod io;
use std::num::ParseIntError;
use std::path::PathBuf;
use structopt::StructOpt;

#[derive(StructOpt, Debug)]
#[structopt(about = "Advent of Code 2020")]
enum App {
    #[structopt(name = "01", about = "Runs the first day's exercise(s)")]
    One {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/day01.txt")]
        #[structopt(help = "Path to the expense report")]
        path: PathBuf,
    },
}

fn main() -> Result<(), ParseIntError> {
    let args = App::from_args();
    return match args {
        App::One { path } => day01::cmd(path),
    };
}
