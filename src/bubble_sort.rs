use std::cmp::Ordering;

pub fn bubble_sort(array: &mut [i32]) {
    let n = array.len();
    for i in 0..n {
        for j in 0..(n-i-1) {
            match array[j].cmp(&array[j + 1]) {
                Ordering::Greater => {
                    array.swap(j, j+1);
                },
                _ => (),
            }
        }
    }
}
