#!/usr/bin/env python
#-*- coding:utf-8 -*-


__doc__ = "sample data source from python, make this in a function"

import requests, json

BIGIP_ADDRESS = '172.20.73.240'
BIGIP_USER = 'admin'
BIGIP_PASS = 'admin'


try:
  # '''
  bigip = requests.session()
  bigip.auth = (BIGIP_USER, BIGIP_PASS)
  bigip.verify = False
  bigip.headers.update({'Content-Type' : 'application/json'})
  BIGIP_URL_BASE = 'https://%s/mgmt/tm' % BIGIP_ADDRESS
  result = bigip.get(BIGIP_URL_BASE)
  if result.status_code == 200:
    print "operation success"
    print result.status_code
  else:
    print "operation fail, cause: "
    print result.text
  # '''
  '''
  BIGIP_URL_BASE = 'https://%s/mgmt/tm' % BIGIP_ADDRESS
  r = requests.get(BIGIP_URL_BASE, auth=('admin', 'admin'))
  print r.status_code
  '''
except Exception, e:
  print e


if __name__ == '__main__':
  print __doc__





