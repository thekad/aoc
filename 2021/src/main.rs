mod five;
mod four;
mod io;
mod one;
mod three;
mod two;
use std::num::ParseIntError;
use std::path::PathBuf;
use structopt::StructOpt;

#[derive(StructOpt, Debug)]
#[structopt(about = "Advent of Code 2020")]
enum App {
    #[structopt(name = "one", about = "Runs the first day's exercise(s)")]
    One {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/one.txt")]
        #[structopt(help = "Path to the radar readings file")]
        path: PathBuf,
    },
    #[structopt(name = "two", about = "Runs the second day's exercise(s)")]
    Two {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/two.txt")]
        #[structopt(help = "Path to the nav map file")]
        path: PathBuf,
    },
    #[structopt(name = "three", about = "Runs the third day's exercise(s)")]
    Three {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/three.txt")]
        #[structopt(help = "Path to the nav map file")]
        path: PathBuf,
    },
    #[structopt(name = "four", about = "Runs the fourth day's exercise(s)")]
    Four {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/four.txt")]
        #[structopt(help = "Path to the bingo input file")]
        path: PathBuf,
    },
    #[structopt(name = "five", about = "Runs the fifth day's exercise(s)")]
    Five {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/five.txt")]
        #[structopt(help = "Path to the hydrothermal line map")]
        path: PathBuf,
    },
}

fn main() -> Result<(), ParseIntError> {
    let args = App::from_args();
    return match args {
        App::One { path } => one::cmd(path),
        App::Two { path } => two::cmd(path),
        App::Three { path } => three::cmd(path),
        App::Four { path } => four::cmd(path),
        App::Five { path } => five::cmd(path),
    };
}
