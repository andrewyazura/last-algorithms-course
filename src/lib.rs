pub mod binary_search;
pub mod two_crystal_balls;
pub mod bubble_sort;
pub mod queue;
pub mod maze_solver;

#[cfg(test)]
mod tests {
    use super::*;
    use binary_search::binary_search;
    use two_crystal_balls::two_crystal_balls;
    use bubble_sort::bubble_sort;
    use queue::Queue;

    #[test]
    fn test_binary_search() {
        let haystack: [i32; 10] = core::array::from_fn(|i| i as i32);
        assert_eq!(binary_search(&haystack, 5), Some(5));
    }
    
    #[test]
    fn test_binary_search_no_element() {
        let haystack: [i32; 10] = core::array::from_fn(|i| i as i32);
        assert_eq!(binary_search(&haystack, 50), None);
    }

    #[test]
    fn test_two_crystal_balls() {
        let breaks: [bool; 10] = [false, false, false, false, false, false, false, true, true, true];
        assert_eq!(two_crystal_balls(&breaks), Some(7));
    }

    #[test]
    fn test_two_crystal_balls_no_break() {
        let breaks: [bool; 10] = [false, false, false, false, false, false, false, false, false, false];
        assert_eq!(two_crystal_balls(&breaks), None);
    }

    #[test]
    fn test_bubble_sort() {
        let mut unsorted_array: [i32; 10] = [9, 0, 7, 2, 5, 4, 3, 6, 1, 8];
        let sorted_array: [i32; 10] = core::array::from_fn(|i| i as i32);
        
        bubble_sort(&mut unsorted_array);
        assert_eq!(unsorted_array, sorted_array);
    }

    #[test]
    fn test_bubble_sort_ones() {
        let mut array: [i32; 10] = [1; 10];
        
        bubble_sort(&mut array);
        assert_eq!(array, array);
    }

    #[test]
    fn test_queue() {
        let mut q = Queue::new();

        q.push(1);
        q.push(2);
        q.push(3);

        assert_eq!(q.pop().unwrap(), 3);
        assert_eq!(q.peek().unwrap(), &2);
        assert_eq!(q.pop().unwrap(), 2);
        assert_eq!(q.pop().unwrap(), 1);
        assert_eq!(q.pop(), None);
    }
}
