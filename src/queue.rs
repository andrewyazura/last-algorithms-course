pub struct Queue {
    queue: Vec<i32>,
}

impl Queue {
    pub fn new() -> Queue {
        Queue { queue: Vec::new() }
    }

    pub fn push(&mut self, value: i32) {
        self.queue.push(value);
    }

    pub fn pop(&mut self) -> Option<i32> {
        self.queue.pop()
    }

    pub fn peek(&self) -> Option<&i32> {
        self.queue.last()
    }
}
