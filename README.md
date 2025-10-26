```markdown
# Simple Chatroom RPC

This is a simple chatroom application using **Go (Golang)** and **RPC**  
Each client provides its name and sends messages; the server stores them and returns the full chat history.

---
```
## 📂 Structure

```

simple-chatroom-rpc/
├── server/
│   └── main.go
└── client/
└── main.go

````

- `server/main.go`: runs the chat server; stores messages and handles RPC calls  
- `client/main.go`: asks user for name, sends messages, fetches chat history  

---

## 🚀 How to Run

1. **Start server**  
   ```bash
   cd server
   go run main.go
````

2. **Start client(s)**
   In another terminal:

   ```bash
   cd client
   go run main.go
   ```

3. **Chat**

   * Enter your name
   * Type messages
   * To exit: type `exit`

---

## 📹 Demo Video
```
Watch this short demo showing multiple clients chatting:
[https://drive.google.com/file/d/1ku51nuXK2Fnsgn54KxjGWWOnN68RBtc5/view?usp=sharing]
ج

```

