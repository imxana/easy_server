require 'socket'
server = TCPServer.new("0.0.0.0", 8000)
loop do
  socket = server.accept
while socket.gets.chop.length > 0
end
socket.puts "HTTP/1.1 200 OK"
socket.puts "Content-type: text/html"
socket.puts ""
socket.puts "<html>"
socket.puts "<body>"
socket.puts "<center>"
socket.puts "<h1>#{Time.now}</h1>"
socket.puts "</center>"
socket.puts "</body>"
socket.puts "</html>"
socket.close
end
