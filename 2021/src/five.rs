use std::collections::HashMap;
use std::num::ParseIntError;
use std::path::PathBuf;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Hash, Clone)]
struct Coord {
    x: i32,
    y: i32,
}

#[derive(Debug)]
struct Segment {
    from: Coord,
    to: Coord,
}

fn draw_map(input: &HashMap<Coord, i32>, max_x: i32, max_y: i32) {
    for y in 0..=max_y {
        for x in 0..=max_x {
            let key = Coord { x: x, y: y };
            if let Some(coord) = input.get(&key) {
                if *coord > 0 {
                    print!("{}", coord);
                } else {
                    print!(".");
                }
            } else {
                print!(".");
            }
        }
        print!("\n");
    }
}

fn distance(p: &Coord, q: &Coord) -> i32 {
    let d = (p.x - q.x) * (p.x - q.x) + (p.y - q.y) * (p.y - q.y);
    //println!("{}", d);
    return d;
}

fn is_square(p1: &Coord, p2: &Coord, q1: &Coord, q2: &Coord) -> bool {
    if distance(p1, q1) == distance(q1, p2)
        && distance(p2, q2) == distance(q2, p1)
        && distance(p1, q1) == distance(p2, q2)
        && distance(q1, p2) == distance(p1, q2)
    {
        return true;
    }

    return false;
}

fn trace(p1: &Coord, p2: &Coord, diagonals: bool) {
    // we get the two opposite corners, get the other 2
    let q1 = Coord { x: p1.x, y: p2.y };
    let q2 = Coord { x: p2.x, y: p1.y };
    println!("{:?} {:?} {:?} {:?}", p1, p2, q1, q2);
    if diagonals {
        if is_square(p1, p2, &q1, &q2) {
            println!("is square")
        }
    }
}

fn walk(segments: &Vec<Segment>, diagonals: bool) -> HashMap<Coord, i32> {
    // walk the map
    let mut map: HashMap<Coord, i32> = HashMap::new();
    for segment in segments {
        trace(&segment.from, &segment.to, false);
        let step_x: i32 = if segment.from.x <= segment.to.x {
            1
        } else {
            -1
        };
        let step_y: i32 = if segment.from.y <= segment.to.y {
            1
        } else {
            -1
        };
        let mut x: i32 = segment.from.x;
        let mut y: i32 = segment.from.y;
        'y: loop {
            'x: loop {
                let key = Coord { x: x, y: y };
                if !map.contains_key(&key) {
                    map.insert(key.clone(), 0);
                }

                // straight lines
                if segment.from.x == segment.to.x || segment.from.y == segment.to.y {
                    let coord = map.get_mut(&key).unwrap();
                    *coord += 1;
                }

                // diagonals
                if diagonals {
                    println!("{:?}", key);
                }

                if x == segment.to.x {
                    break 'x;
                } else {
                    x += step_x;
                }
            }
            if y == segment.to.y {
                break 'y;
            } else {
                y += step_y;
            }
        }
    }

    return map;
}

pub fn cmd(path: PathBuf) -> Result<(), ParseIntError> {
    let mut segments: Vec<Segment> = Vec::new();

    let mut xs: Vec<i32> = Vec::new();
    let mut ys: Vec<i32> = Vec::new();
    if let Ok(lines) = crate::io::read_lines(path) {
        for mut line in lines {
            line.retain(|c| !c.is_whitespace());
            let pairs = line.split("->").collect::<Vec<&str>>();
            let p1 = pairs[0].split(",").collect::<Vec<&str>>();
            let p2 = pairs[1].split(",").collect::<Vec<&str>>();
            let x1: i32 = p1[0].parse().unwrap();
            let y1: i32 = p1[1].parse().unwrap();
            let x2: i32 = p2[0].parse().unwrap();
            let y2: i32 = p2[1].parse().unwrap();
            xs.extend([x1, x2]);
            ys.extend([y1, y2]);
            segments.push(Segment {
                from: Coord { x: x1, y: y1 },
                to: Coord { x: x2, y: y2 },
            })
        }

        let straights = walk(&segments, false);
        // draw the map because why not
        let max_x: i32 = xs.iter().max().unwrap().clone();
        let max_y: i32 = ys.iter().max().unwrap().clone();

        // part 1
        let mut safe: i32 = 0;
        for (_, item) in &straights {
            if *item >= 2 {
                safe += 1;
            }
        }
        if max_y < 20 {
            draw_map(&straights, max_x, max_y);
        }
        println!("Safe spots in map: {}", &safe);

        /*
        let diagonals = walk(&segments, true);
        if max_y < 20 {
            draw_map(&diagonals, max_x, max_y);
        }
        */
    }

    Ok(())
}
