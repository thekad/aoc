mod eight;
mod five;
mod four;
mod io;
mod nine;
mod one;
mod seven;
mod six;
mod ten;
mod three;
mod two;
use std::num::ParseIntError;
use std::path::PathBuf;
use structopt::StructOpt;

#[derive(StructOpt, Debug)]
#[structopt(about = "Advent of Code 2020")]
enum App {
    #[structopt(name = "01", about = "Runs the first day's exercise(s)")]
    One {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/expense-report.txt")]
        #[structopt(help = "Path to the expense report")]
        path: PathBuf,
    },
    #[structopt(name = "02", about = "Runs the second day's exercise(s)")]
    Two {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/password-policies.txt")]
        #[structopt(help = "Path to the file containing the passwords and policies")]
        path: PathBuf,
    },
    #[structopt(name = "03", about = "Runs the third day's exercise(s)")]
    Three {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/tree-map.txt")]
        #[structopt(help = "Path to the file containing the tree map")]
        path: PathBuf,
    },
    #[structopt(name = "04", about = "Runs the fourth day's exercise(s)")]
    Four {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/passports.txt")]
        #[structopt(help = "Path to the file containing passport data")]
        path: PathBuf,
    },
    #[structopt(name = "05", about = "Runs the fifth day's exercise(s)")]
    Five {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/seat-assignment.txt")]
        #[structopt(help = "Path to the file containing flight seat assignments")]
        path: PathBuf,
    },
    #[structopt(name = "06", about = "Runs the sixth day's exercise(s)")]
    Six {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/customs-declaration.txt")]
        #[structopt(help = "Path to the file containing custom declaration form answers")]
        path: PathBuf,
    },
    #[structopt(name = "07", about = "Runs the seventh day's exercise(s)")]
    Seven {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/luggage-policies.txt")]
        #[structopt(help = "Path to the file containing airport luggage policies")]
        path: PathBuf,
    },
    #[structopt(name = "08", about = "Runs the eighth day's exercise(s)")]
    Eight {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/game-debug.txt")]
        #[structopt(help = "Path to the file containing game device debug info")]
        path: PathBuf,
    },
    #[structopt(name = "09", about = "Runs the ninth day's exercise(s)")]
    Nine {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/xmas-encoded.txt")]
        #[structopt(help = "Path to the file containing XMAS encrypted data")]
        path: PathBuf,
    },
    #[structopt(name = "10", about = "Runs the tenth day's exercise(s)")]
    Ten {
        #[structopt(parse(from_os_str))]
        #[structopt(default_value = "data/joltage-adapters.txt")]
        #[structopt(help = "Path to the file containing your joltage adapter list")]
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
        App::Six { path } => six::cmd(path),
        App::Seven { path } => seven::cmd(path),
        App::Eight { path } => eight::cmd(path),
        App::Nine { path } => nine::cmd(path),
        App::Ten { path } => ten::cmd(path),
    };
}
