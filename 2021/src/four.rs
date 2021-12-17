use std::fs;
use std::num::ParseIntError;
use std::path::PathBuf;

#[derive(Debug, Clone, Copy)]
struct Box {
    number: u8,
    checked: bool,
}

fn won(board: &Vec<Box>) -> bool {
    let winning: Vec<Vec<usize>> = vec![
        // horizontals
        vec![0, 1, 2, 3, 4],
        vec![5, 6, 7, 8, 9],
        vec![10, 11, 12, 13, 14],
        vec![15, 16, 17, 18, 19],
        vec![20, 21, 22, 23, 24],
        // verticals
        vec![0, 5, 10, 15, 20],
        vec![1, 6, 11, 16, 21],
        vec![2, 7, 12, 17, 22],
        vec![3, 8, 13, 18, 23],
        vec![4, 9, 14, 19, 24],
    ];

    for win in &winning {
        let mut counter = 0;
        for check in win {
            if board[*check].checked {
                counter += 1;
            }
        }
        if counter == 5 {
            println!("Winning line: {:?}", &win);
            return true;
        }
    }

    return false;
}

fn sum(board: &Vec<Box>) -> i32 {
    let mut total: i32 = 0;
    for _box in board {
        if !_box.checked {
            total += i32::from(_box.number);
        }
    }

    return total;
}

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    let mut numbers: Vec<u8> = Vec::new();
    let mut boards: Vec<Vec<Box>> = Vec::new();

    if let Ok(contents) = fs::read_to_string(path) {
        for part in contents.split("\n\n").collect::<Vec<&str>>() {
            // read the random numbers
            if numbers.len() == 0 {
                for piece in part.split(",").collect::<Vec<&str>>() {
                    numbers.push(piece.parse().unwrap())
                }
                continue;
            }

            // initialize boards
            let mut _board: Vec<Box> = Vec::new();
            for _row in part.split("\n").collect::<Vec<&str>>() {
                let mut _boxes: Vec<u8> = Vec::new();
                for _box in _row.split_whitespace().collect::<Vec<&str>>() {
                    _board.push(Box {
                        number: _box.parse().unwrap(),
                        checked: false,
                    });
                }
            }
            boards.push(_board);
        }
    }
    println!("Boards playing: {:?}", &boards.len());

    // part 1
    // play the game to win
    'numbers: for number in &numbers {
        println!("Ball #{} drawn", number);
        for idx in 0..boards.len() {
            let board = boards.get_mut(idx).unwrap();
            'boxes: for _box in board {
                if _box.number == *number {
                    _box.checked = true;
                    break 'boxes;
                }
            }
            if won(boards.get(idx).unwrap()) {
                println!("We have a winner! Board #{} wins!", idx + 1);
                let remaining = sum(boards.get(idx).unwrap());
                let score: i32 = remaining * i32::from(*number);
                println!("Sum of unchecked boxes: {}", remaining);
                println!("Winning board score: {}", score);
                break 'numbers;
            };
        }
    }

    // part 2
    // play the game to lose
    let mut winners: Vec<usize> = Vec::new();
    'numbers2: for number in &numbers {
        println!("Ball #{} drawn", number);
        'boards2: for idx in 0..boards.len() {
            if winners.contains(&idx) {
                println!("Board #{} already a winner, moving on...", &idx);
                continue 'boards2;
            }

            let board = boards.get_mut(idx).unwrap();
            for _box in board {
                if _box.number == *number {
                    _box.checked = true;
                    break;
                }
            }

            if won(boards.get(idx).unwrap()) {
                println!("We have a winner! Board #{} wins!", idx + 1);
                winners.push(idx);
                if winners.len() == boards.len() {
                    let remaining = sum(boards.get(idx).unwrap());
                    let score: i32 = remaining * i32::from(*number);
                    println!("Sum of unchecked boxes: {}", remaining);
                    println!("Winning board score: {}", score);
                    break 'numbers2;
                }
            };
        }
    }

    Ok(())
}
