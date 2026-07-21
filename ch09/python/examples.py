"""
Chapter 09: Examples
"""


class Stack:
    def __init__(self):
        self.data = []
    
    def push(self, element):
        self.data.append(element)

    def pop(self):
        if len(self.data) > 0:
            return self.data.pop()
        else:
            return None

    def read(self):
        if len(self.data) > 0:
            return self.data[-1]
        else:
            return None


class Linter:
    def __init__(self):
        self.stack = Stack()

    def lint(self, text):
        for char in text:
            if self._is_opening_brace(char):
                self.stack.push(char)
            elif self._is_closing_brace(char):
                popped_opening_brace = self.stack.pop()

                if not popped_opening_brace:
                    return f"{char} doesn't have an opening brace"
                
                if self._is_not_match(popped_opening_brace, char):
                    return f"{char} has mismatched opening brace"
                
        if self.stack.read():
            return f"{self.stack.read()} does not have a closing brace"

    def _is_opening_brace(self, char):
        return char in ["(", "[", "{"]
    
    def _is_closing_brace(self, char):
        return char in [")", "]", "}"]
    
    def _is_not_match(self, opening_brace, closing_brace):
        if (
            (opening_brace == "(" and closing_brace == ")")
            or (opening_brace == "[" and closing_brace == "]")
            or (opening_brace == "{" and closing_brace == "}")
        ):
            return False
        
        return True
    

class Queue:
    def __init__(self):
        self.data = []

    def enqueue(self, element):
        self.data.append(element)

    def dequeue(self):
        if len(self.data) > 0:
            return self.data.pop(0)
        else:
            return None
    
    def read(self):
        if len(self.data) > 0:
            return self.data[0]
        else:
            return None
        

class PrintManager:
    def __init__(self):
        self.queue = Queue()

    def queue_print_job(self, document):
        self.queue.enqueue(document)

    def run(self):
        while self.queue.read():
            self._print(self.queue.dequeue())

    def _print(self, document):
        print(f"Document: {document}")


print("******** Stack ********")
stack = Stack()
assert stack.pop() is None
assert stack.read() is None

stack.push(1)
stack.push(2)
stack.push(3)
assert stack.read() == 3
assert stack.pop() == 3
assert stack.read() == 2


print("******** Linter ********")
assert Linter().lint("(a)") is None
assert Linter().lint("(a[b]{c})") is None
assert Linter().lint("(a") == "( does not have a closing brace"
assert Linter().lint("(a]") == "] has mismatched opening brace"
assert Linter().lint("a)") == ") doesn't have an opening brace"


print("******** Queue ********")
queue = Queue()
assert queue.dequeue() is None
assert queue.read() is None

queue.enqueue("first.pdf")
queue.enqueue("second.pdf")
queue.enqueue("third.pdf")
assert queue.read() == "first.pdf"
assert queue.dequeue() == "first.pdf"
assert queue.read() == "second.pdf"


print("******** Print Manager ********")
print_manager = PrintManager()
print_manager.queue_print_job("first.pdf")
print_manager.queue_print_job("second.pdf")
print_manager.queue_print_job("third.pdf")
print_manager.run()