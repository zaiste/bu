from pyhocon import ConfigFactory

conf = ConfigFactory.parse_file('sample.conf')
host = conf.get_string('databases.mysql.host')
same_host = conf.get('databases.mysql.host')
same_host = conf['databases.mysql.host']
same_host = conf['databases']['mysql.host']
port = conf['databases.mysql.port']
username = conf['databases']['mysql']['username']
password = conf.get_config('databases')['mysql.password']

print host