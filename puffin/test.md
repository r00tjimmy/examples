

#### openladp介绍

 openldap是一个轻量级的目录权限协议，由OpenLDAP开发项目. 也是一个独立的协议，所有支持所有的Linux/Unix like 系统, Windows, AIX, Solaris 和Android.

OpenLDAP 包含有:

- slapd – 标准独立的 LDAP 进程(server)
- LDAP继承的库, 工具, 和简单的客户端

本文档记录怎么安装/配置/使用OpenLDAP， 服务端/客户端都是Debian8.x 64 位服务器(centos操作也一样)

这是我测试系统的详细信息:

- **服务端主机**:        172.20.73.229/24（ server1.example.com）
- **客户端主机**:        172.20.73.239/24 （client1.example.com)





#### openldap配置

- ##### 服务端		

1. 安装软件包

   ```shell
   sudo apt-get install slapd ldap-utils
   ```

​           安装的时候需要配置LDAP管理员密码

2.    编辑LDAP配置文件

   ```shell
   vi /etc/ldap/ldap.conf
   ```

     配置LDAP的域和服务器地址

   ```shell
   BASE  dc=example,dc=com
   URI ldap://172.20.73.229 ldap://172.20.73.229:666
   ```

3.    更新LDAP服务

   ```shell
   dpkg-reconfigure slapd
   ```

   会提示输入LDAP的域，配置成我们主机的域 example.com， 其他配置默认即可。

4.    验证LADP服务是否配置正确

   ```shell
   ldapsearch -x

   #输出包含下面信息

   # search result
   search: 2
   result: 0 Success
   ```


   即标识LDAP服务器配置成功。





- ##### 客户端


1. 安装软件包

   ```shell
   sudo apt-get install libnss-ldapd libpam-ldapd
   ```

   安装过程中根据安装向导输入ldap服务器IP和相应域信息，其中nss services要勾选group和passwd两项，即表示支持组和用户的LDAP认证登录

2. 编辑nss服务的配置文件，使系统支持ldap用户组用户

   ```shell
   /etc/nsswitch.conf
   #分别修改下面的三个值为
   passwd:         compat ldap
   group:          compat ldap
   shadow:         compat ldap
   netgroup:		nis ldap
   ```

​        



#### openldap使用

- ##### LDAP服务器管理

  phpLDAPadmin是一个基于LDAP的管理工具，是较为常用的开源openldap web ui，用来管理LDAP服务，使用phpLDAPadmin，你可以查看LDAP列表、查看LDAP计划、搜索、创建、删除、复制和编辑LDAP等，你还可以从不同服务器之间复制LDAP 的entries。

  ​

  安装在openldap服务端

  ```shell
  apt-get install apache2 php5 php5-mysql phpldapadmin
  ```

  配置

  ```shell
  vi /etc/phpldapadmin/config.php

  [...]
  // Set your LDAP server name //
  $servers->setValue('server','name','Unixmen LDAP Server');
  [...]
  // Set your LDAP server IP address // 
  $servers->setValue('server','host','172.20.73.229');
  [...]
  // Set Server domain name //
  $servers->setValue('server','base',array('dc=example,dc=com'));
  [...]
  // Set Server domain name again//
  $servers->setValue('login','bind_id','cn=admin,dc=example,dc=com');
  [...]
  ```


  访问

  ```shell
  http://172.20.73.229/phpldapadmin
  ```

​	用原来配置服务端的密码登录， 然后添加一个用户组sales-group，添加一个账户jimmy

​	![QQ截图20170630143332](C:\Users\r00xx\Desktop\QQ截图20170630143332.png)





- ##### LDAP客户端验证/登录

  原来客户端系统是没有jimmy这个用户的

  ```shell
  shell~# id jimmy
  id: jimmy: no such user
  ```

  在ldap服务器添加该用户后，在客户端验证刚才在服务端添加的jimmy用户是否生效

  ```shell
  shell~# id jimmy
  uid=1000(jimmy) gid=500(sales-group) groups=500(sales-group)
  ```

   可以在客户端检测出来存在 jimmy这个用户，证明系统已从ldap服务器去检测用户名了，接下来验证该用户的ssh登录

  ```shell
  λ ssh jimmy@172.20.73.239
  jimmy@172.20.73.239's password:

  The programs included with the Debian GNU/Linux system are free software;
  the exact distribution terms for each program are described in the
  individual files in /usr/share/doc/*/copyright.

  Debian GNU/Linux comes with ABSOLUTELY NO WARRANTY, to the extent
  permitted by applicable law.
  Last login: Tue Jun 13 10:43:29 2017 from 172.21.28.171
  jimmy@yypp-jimmy-vsphere:~$ whoami
  jimmy
  jimmy@yypp-jimmy-vsphere:~$

  ```


  以jimmy远程ssh登录成功，证明系统用户已可以支持从 openladp服务器统一认证登录。





#### 结论

​     通过一些测试使用，可见openldap满足多个操作系统共用一个用户，用户组认证中心的需求，可用作来统一管理多系统的用户信息，其中一  些配置细节还需要根据特定的场景可能有变动，例如加密算法，用户权限等。




