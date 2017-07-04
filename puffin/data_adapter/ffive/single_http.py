from BaseHTTPServer import BaseHTTPRequestHandler, HTTPServer
from SocketServer import ThreadingMixIn

import threading
import argparse
import re
import cgi

import sample_func

class LocalData(object):
  records = {}

class HTTPRequestHandler(BaseHTTPRequestHandler):
  def do_GET(self):
    if None != re.search('/getsample', self.path):
      self.send_response(200)
      self.send_header('Content-Type', 'application/json')
      self.end_headers()
      # self.wfile.write('{"name":"jimmy"}')

      ffive = sample_func.Ffive()
      s = ffive.sample()
      self.wfile.write(s)  #get date from ffive
    else:
      self.send_response(400, 'route wrong')
      self.send_header('Content-Type', 'application/json')
      self.end_headers()
    return

class ThreadHTTPServer(ThreadingMixIn, HTTPServer):
  allow_reuse_address = True

  def shutdown(self):
    self.socket.close()
    HTTPServer.shutdown()

class SimpleHttpServer():
  def __init__(self):
    self.server = ThreadHTTPServer(('127.0.0.1', 9000), HTTPRequestHandler)

  def start(self):
    self.server_thread = threading.Thread(target=self.server.serve_forever)
    self.server_thread.daemon = True
    self.server_thread.start()

  def waitForThread(self):
    self.server_thread.join()

  def stop(self):
    self.server.shutdown()
    self.waitForThread()


if __name__ == '__main__':
  server = SimpleHttpServer()
  print "http server running........"
  server.start()
  server.waitForThread()


















