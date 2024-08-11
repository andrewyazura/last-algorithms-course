use std::usize;

pub fn two_crystal_balls(breaks: &[bool]) -> Option<usize> {
    let n = breaks.len();
    let step = (n as f64).sqrt().floor() as usize;

    let i = (step..n).step_by(step).find(|&i| breaks[i]).unwrap_or(n);

    for j in (i - step)..i {
        if breaks[j] {
            return Some(j as usize)
        }
    }

    None
}
