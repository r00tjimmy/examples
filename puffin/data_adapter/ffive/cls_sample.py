#!/usr/bin/env python
#-*- coding:utf-8 -*-

class TestClassMethod(object):
  METHOD = 'method hoho'

  def __init__(self):
    self.name = 'jimmy'

  def test1(self):
    print 'test1'
    print self
    print TestClassMethod.METHOD
    print "---------------------\n"

  @classmethod
  def test2(cls):
    print 'test2'
    print cls
    print TestClassMethod.METHOD
    print "---------------------\n"

  @staticmethod
  def test3():
    print 'test3'
    print TestClassMethod.METHOD
    print "---------------------\n"


if __name__ == '__main__':
  a = TestClassMethod()
  a.test1()
  a.test2()
  a.test3()

  # TestClassMethod.test1()
  TestClassMethod.test2()
  TestClassMethod.test3()


