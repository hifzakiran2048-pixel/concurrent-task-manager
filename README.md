# Concurrent Task Manager (Go)

A CLI-based concurrent task manager built using Go.

## 🚀 Features
- Add tasks
- View tasks
- Process tasks asynchronously
- Worker pool using goroutines
- Channel-based communication
- Mutex for safe concurrency

## 🧠 Concepts Used
- Goroutines
- Channels
- Worker Pool Pattern
- Mutex (sync)
- WaitGroup
- Clean Architecture (Layered Design)

## 📂 Project Structure
- model → data structures
- store → in-memory storage
- service → business logic
- worker → concurrency logic
- main → CLI interface

## ▶️ Run Project

```bash
go run main.go