Here, use `nc -v localhost 42069` to test the connection, and you should see a `connection succeeded` message.
Make sure you have run the `main.go` before you do the above.

Now, run the same command after terminating the go program.
You will see a `connection refused` message. This is because we are usign a TCP connection and it does a handshake while creating the connection. The same will not be done in protocols like `UDP`.  Now this is because TCP requires the connection to be established before data can be sent.