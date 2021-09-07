const netInterfaces = require("os").networkInterfaces();
const express = require("express");
const app = express();

app.get("/", (_, res) => {
  const serverIP = Object.values(netInterfaces).flat().find(i => i.family === "IPv4" && !i.internal).address;
	res.send({ status: "Ok", ip: serverIP });
});

app.listen("8080", () => {
  console.info("Server is up and running on port 8080");
})
