use std::cmp::Ordering;

pub fn binary_search(haystack: &[i32], needle: i32) -> Option<usize> {
    let mut lo: usize = 0;
    let mut hi: usize = haystack.len();

    while lo < hi {
        let m: usize = lo + (hi - lo) / 2;

        match haystack[m].cmp(&needle) {
            Ordering::Equal => return Some(m),
            Ordering::Greater => lo = m + 1,
            Ordering::Less => hi = m,
        }
    }

    return None;
}
