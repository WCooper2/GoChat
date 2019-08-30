/*
`* File: backend/main.go
 * Define the connection interface from this client to the server
 * Will Cooper
*/

//Open the socket for server communications
var socket = new WebSocket("ws://localhost:8080/ws");

//Attempts to connect to defined socket
let connect = () => {

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log("Received");
    console.log(msg);
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

//Send a message back to the server via websocket 
let sendMessage = msg => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export { connect, sendMessage };
