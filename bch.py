import bcrypt
import argparse

def out(x):
    print(x.decode('utf-8'))
    
parser = argparse.ArgumentParser(description='Generate bcrypt password hashes')
parser.add_argument('password', help='password to hash')
parser.add_argument('-f', '--factor', type=int, required=False, default=12, help='work factor', metavar='F')

args = parser.parse_args()

password = args.password.encode('utf-8')
factor = args.factor
salt = bcrypt.gensalt(factor)
hash = bcrypt.hashpw(password, salt)
out(hash)