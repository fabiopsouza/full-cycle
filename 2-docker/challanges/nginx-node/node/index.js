const { response } = require('express');
const express = require('express')
const app = express();
const port = 3000
const config = {
  host: 'db',
  user: 'root',
  password: 'root',
  database: 'nodedb'
};

const mysql = require('mysql');
const connection = mysql.createConnection(config);
const ddl = `CREATE TABLE IF NOT EXISTS people (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255))`
const insert = `INSERT INTO people(name) VALUES('Fabio')`
connection.query(ddl);
connection.query(insert);

function buildResponse(result) {
  let response = `<h1>Full Cycle Rocks!</h1>`
  response += `<table><tr><th>ID</th><th>Name</th></tr>`
  result.forEach((p) => response += `<tr><td>${p.id}</td><td>${p.name}</td></tr>`);
  response += `</table>`

  return response
}

function selectCallback(err, result, res) {
  if (err) {
    console.error('mysql select error: ', err);
    res.send(`mysql select error: ${err}`);
  } else {
    const response = buildResponse(result)
    res.send(response);
  }
}

app.get('/', (re,res) => {
  const select = `SELECT id, name FROM people`
  connection.query(select, (err, result, _) => selectCallback(err, result, res));
})

app.listen(port, () => {
  console.log('Rodando na porta ' + port);
})