#!/usr/bin/env python
# -*- coding:utf-8 -*-


__doc__ = "sample data source from python, make this in a function"

import requests

BIGIP_ADDRESS = '172.20.73.240'
BIGIP_USER = 'admin'
BIGIP_PASS = 'admin'

from requests.packages.urllib3.exceptions import InsecureRequestWarning
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)


class Ffive(object):
  def sample(self):
    try:
      bigip = requests.session()
      bigip.auth = (BIGIP_USER, BIGIP_PASS)
      bigip.verify = False
      bigip.headers.update({'Content-Type': 'application/json'})
      BIGIP_URL_BASE = 'https://%s/mgmt/tm' % BIGIP_ADDRESS
      result = bigip.get(BIGIP_URL_BASE)
      if result.status_code == 200:
        # print result.status_code
        return result.status_code
      else:
        # print "operation fail"
        # print result.text
        return 'error'
    except Exception, e:
      # print e
      return 'except'



if __name__ == '__main__':
  # Ffive.sample()
  f = Ffive()
  f.sample()


