use std::{fs, io};

// use serde::Deserialize;
use csv::ReaderBuilder;


// 1325	He's kicking me!	1767703	Hy skop my!
// #[derive(Debug, Deserialize)]
// struct Record {
//     source_num: i32,
//     source_content: String,
//     target_num:  i32,
//     target_content: String,
// }

fn main()-> Result<(), io::Error> {
    let paths = fs::read_dir("/Users/john/lingoSpringTatoeba/data").unwrap();

    let mut counter = 0;
    for path in paths {
        let file_path = path.unwrap().path();
        println!("Name: {}", file_path.display());
        let mut rdr = read_csv(file_path);
        for result in rdr.records() {
            println!("{:?}", result.unwrap());
            counter += 1;
        }
    }
    println!("counter: {}", counter);
    return Ok(());
}

fn read_csv(path: std::path::PathBuf) -> csv::Reader<std::fs::File> {
    let file = std::fs::File::open(path).unwrap();
    let rdr = ReaderBuilder::new()
        .delimiter(b'\t')
        .has_headers(false)
        .from_reader(file);
    return rdr
}
