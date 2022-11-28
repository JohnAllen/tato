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

fn main() -> Result<(), io::Error> {
    let paths = fs::read_dir("/Users/john/lingospring-tatoeba-html/data").unwrap();

    let mut counter = 0;
    for path in paths {
        let file_path = path.unwrap().path();
        // println!("Name: {}", file_path.display());
        let mut rdr = read_csv(file_path);
        //     let mut owned_string: String = "hello ".to_owned();
        let mut base_path: String =
            "/Users/john/lingospring-tatoeba-html/html/how-to-say-".to_owned();
        for result in rdr.records() {
            // let tgt_lang_name = &path;
            // println!("tgt_lang_name: {:?}", tgt_lang_name);
            let record = result.unwrap();
            let source_num = record.get(0).unwrap();
            let source_content = record.get(1).unwrap();
            let target_num = record.get(2).unwrap();
            let target_content = record.get(3);
            // + target_content + "-in-" + tgt_lang_name;
            let mut final_path = base_path.push_str(target_content.unwrap());
            // final_path.push_str(&"-in-");
            // (&tgt_lang_name);
            // println!("Final path: {}", final_path.);
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
    return rdr;
}

// fn create_html_file(file_name: String) -> std::fs::File {
//     let file = std::fs::File::create(file).unwrap();
//     return file;
// }

// fn get_file_name(path: std::path::PathBuf) -> String {
//     let file_name = path.file_name().unwrap().to_str().unwrap();
//     println!("file_name: {}", file_name);
//     return file_name.to_string();
// }
