var express = require('express')
var path = require('path')
var app = express()

//1. use static
app.use(express.static(path.join(__dirname, '../static')))

//2. use route
app.get('/a', function(req, res){
    //res.end('hello world')
  res.sendFile(path.join(__dirname, '../static', 'a.html'));
})


app.listen(8000, function(){
  console.log('Listening on *:8000')
})
