const express = require('express');
const config  = require('./config')
const app = express();

 
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
    fetch(`${config.API_URL}/post/all`, {
      method: 'GET'
    })
      .then(response => response.json())
      .then(data => {
        res.render('posts', { data });
      })
      .catch(error => {
        console.error(error);
        res.render('error');
      });
  });

app.get('/register', (req, res) => {
    res.render('signup');
});

app.get('/login', (req, res) => {
    res.render('login');
});

app.get('/create', (req, res) => {
    res.render('create_post');
});

const server = app.listen(4000, function () {
    console.log('listening to port 4000')
});
