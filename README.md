# Go Whiteboard Practice

A small, test-driven Go repository used to practice common whiteboard and technical interview problems.

This project is intentionally minimal and focused on:
- Clear problem framing
- Idiomatic Go solutions
- Verbal explainability during interviews
- Fast correctness checks via unit tests

---

## Project Structure

go-whiteboard/
├── go.mod
├── README.md
├── problems/
│ ├── twosum.go
│ ├── twosum_test.go
│ ├── parentheses.go
│ ├── parentheses_test.go
│ ├── list.go
│ ├── list_test.go
│ └── ...


Each problem lives in its own file with a corresponding `_test.go` file.

---

## Running Tests

From the project root:

```bash
go test ./... -v

Tests are used to quickly validate correctness and edge cases.

Problem-Solving Approach

For each problem, I follow a consistent interview-friendly flow:

Clarify inputs, outputs, and constraints

Walk through a small example

Choose an approach and state time/space complexity

Implement the solution

Walk the example against the code

Call out edge cases

Why Go

Go’s simplicity makes reasoning about data structures explicit and readable.
Its tooling enables fast feedback, which is ideal for interview-style problem solving.

Notes

This repository is designed for practice and discussion.
Solutions are written to be easily explained on a whiteboard.


Save the file (`Ctrl+S`).

---

## 3️⃣ Verify it exists
Back in terminal:

```bash
ls

You should see:

README.md  go.mod  problems/