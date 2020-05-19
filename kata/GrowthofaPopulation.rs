
fn nb_year(p0: i32, percent: f64, aug: i32, p: i32)-> i32 {
    if p0 > p {
        return 0
    }

    return nb_year(p0 + (p0 as f64 * percent / 100.0) as i32 + aug, percent, aug, p) + 1
}